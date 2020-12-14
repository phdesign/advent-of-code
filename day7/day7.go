package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	color    string
	contains []Bag
}

type BagParseError struct {
	text string
}

func (e *BagParseError) Error() string {
	return fmt.Sprintf("Unable to parse bag %q")
}

var bagPattern = regexp.MustCompile(`^((\w+ ){1,2})bags contain(.*)$`)
var innerPattern = regexp.MustCompile(`^ ?(\d )+((\w+ ){1,2})bags.?$`)

func ParseInner(text string) (int, Bag, error) {
	matches := innerPattern.FindStringSubmatch(text)
	if len(matches) == 0 {
		return 0, nil, &BagParseError{text}
	}
	qty, err := strconv.Atoi(strings.TrimSpace(matches[1]))
	if err != nil {
		return 0, nil, err
	}
	color := strings.TrimSpace(matches[2])
	bag := Bag{color}
	return qty, bag, nil
}

func ParseBag(text string) (Bag, error) {
	matches := bagPattern.FindStringSubmatch(text)
	if len(matches) == 0 {
		return nil, &BagParseError{text}
	}
	color := strings.TrimSpace(matches[1])
	bag := Bag{color}
	contents := strings.Split(matches[len(matches)-1])
	for _, inner := range contents {
		qty, innerBag, err := ParseBag(inner)
		if err != nil {
			continue
		}
		for i := 0; i < qty; i++ {
			append(bag.contains, innerBag)
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
		append(bags, bag)
	}
	return
}
