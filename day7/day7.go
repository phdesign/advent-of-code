package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	color    string
	children []Bag
	parents  []Bag
}

type BagParseError struct {
	text string
}

func (e *BagParseError) Error() string {
	return fmt.Sprintf("Unable to parse bag %q", e.text)
}

var bagPattern = regexp.MustCompile(`^((\w+ ){1,2})bags contain (.*)$`)
var childPattern = regexp.MustCompile(`^(\d )+((\w+ ){1,2})bags.?$`)

func ParseChild(text string) (int, *Bag, error) {
	matches := childPattern.FindStringSubmatch(text)
	if len(matches) == 0 {
		return 0, nil, &BagParseError{text}
	}
	qty, err := strconv.Atoi(strings.TrimSpace(matches[1]))
	if err != nil {
		return 0, nil, err
	}
	color := strings.TrimSpace(matches[2])
	bag := new(Bag)
	bag.color = color
	return qty, bag, nil
}

func ParseBag(text string) (*Bag, error) {
	matches := bagPattern.FindStringSubmatch(text)
	if len(matches) == 0 {
		return nil, &BagParseError{text}
	}
	color := strings.TrimSpace(matches[1])
	bag := new(Bag)
	bag.color = color
	contents := strings.Split(matches[len(matches)-1], ", ")
	for _, childText := range contents {
		qty, child, err := ParseChild(childText)
		if err != nil {
			continue
		}
		for i := 0; i < qty; i++ {
			bag.children = append(bag.children, *child)
			child.parents = append(child.parents, *bag)
		}
	}
	return bag, nil
}

func ParseBags(input string) (bags []Bag) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		bag, err := ParseBag(line)
		if err != nil {
			continue
		}
		bags = append(bags, *bag)
	}
	fmt.Println(bags[0])
	return
}
