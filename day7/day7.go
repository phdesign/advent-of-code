package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	color    string
	children []*Bag
	parents  []*Bag
}

func (b Bag) String() string {
	return fmt.Sprintf("Bag{color: %q, children: [%d], parents: [%d]}", b.color, len(b.children), len(b.parents))
}

type Cache map[string]*Bag

func (c *Cache) GetOrAdd(color string) *Bag {
	bag, ok := (*c)[color]
	if !ok {
		bag = &Bag{color: color}
		(*c)[color] = bag
	}
	return bag
}

var bagPattern = regexp.MustCompile(`^((\w+ ){1,2})bags contain (.*)$`)
var childPattern = regexp.MustCompile(`^(\d )+((\w+ ){1,2})bags?.?$`)

func ParseChild(text string, parent *Bag, cache *Cache) {
	if text == "no other bags." {
		return
	}
	matches := childPattern.FindStringSubmatch(text)
	if len(matches) == 0 {
		panic(fmt.Sprintf("Unable to parse child %q", text))
	}
	qty, err := strconv.Atoi(strings.TrimSpace(matches[1]))
	if err != nil {
		panic(fmt.Sprintf("Unable to parse child qty %q", strings.TrimSpace(matches[1])))
	}
	color := strings.TrimSpace(matches[2])
	bag := cache.GetOrAdd(color)
	for i := 0; i < qty; i++ {
		parent.children = append(parent.children, bag)
		bag.parents = append(bag.parents, parent)
	}
}

func ParseBag(text string, cache *Cache) {
	matches := bagPattern.FindStringSubmatch(text)
	if len(matches) == 0 {
		panic(fmt.Sprintf("Unable to parse bag %q", text))
	}
	color := strings.TrimSpace(matches[1])
	bag := cache.GetOrAdd(color)
	contents := strings.Split(matches[len(matches)-1], ", ")
	for _, childText := range contents {
		ParseChild(childText, bag, cache)
	}
}

func ParseBags(input string) Cache {
	lines := strings.Split(input, "\n")
	cache := make(Cache)
	for _, line := range lines {
		ParseBag(line, &cache)
	}
	return cache
}

func FindAncestors(bag *Bag) (result []string) {
	for _, parent := range bag.parents {
		result = append(result, parent.color)
		grandparents := FindAncestors(parent)
		for _, grandparent := range grandparents {
			result = append(result, grandparent)
		}
	}
	return
}

func CountUniqueAncestors(bag *Bag) int {
	ancestors := FindAncestors(bag)
	hash := make(map[string]bool)
	for _, ancestor := range ancestors {
		hash[ancestor] = true
	}
	return len(hash)
}

func CountContainingBags(input string, bagColor string) int {
	cache := ParseBags(input)
	bag, ok := cache[bagColor]
	if !ok {
		panic(fmt.Sprintf("No bag found with color %q", bagColor))
	}
	return CountUniqueAncestors(bag)
}

func CountChildrenRecursive(bag *Bag) (count int) {
	for _, child := range bag.children {
		count++
		count += CountChildrenRecursive(child)
	}
	return
}

func CountChildBags(input string, bagColor string) int {
	cache := ParseBags(input)
	bag, ok := cache[bagColor]
	if !ok {
		panic(fmt.Sprintf("No bag found with color %q", bagColor))
	}
	return CountChildrenRecursive(bag)
}

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	input := strings.Trim(string(content), "\n")
	result := CountChildBags(input, "shiny gold")
	fmt.Println(result)
}
