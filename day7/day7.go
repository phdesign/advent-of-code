package main

import (
	"fmt"
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

func ParseBags(input string) (bags []Bag) {
	lines := strings.Split(input, "\n")
	cache := make(Cache)
	for _, line := range lines {
		ParseBag(line, &cache)
	}
	for _, bag := range cache {
		bags = append(bags, *bag)
	}
	return
}

func CountContainingBags(input string, bagColor string) int {
	return 0
}
