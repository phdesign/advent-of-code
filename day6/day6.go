package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func CountUnique(group string) int {
	items := strings.Join(strings.Fields(group), "")
	hash := make(map[rune]bool)
	for _, item := range items {
		hash[item] = true
	}
	return len(hash)
}

func CountIntersect(group string) (count int) {
	sets := strings.Split(group, "\n")
	hash := make(map[rune]int)
	for _, set := range sets {
		for _, item := range set {
			hash[item] = hash[item] + 1
		}
	}
	setLength := len(sets)
	for _, v := range hash {
		if v == setLength {
			count++
		}
	}
	return
}

func MapGroup(input string, mapper func(string) int) []int {
	groups := strings.Split(input, "\n\n")
	counts := make([]int, len(groups))
	for i, group := range groups {
		counts[i] = mapper(group)
	}
	return counts
}

func Sum(items []int) (sum int) {
	for _, item := range items {
		sum += item
	}
	return
}

func SumUniqueInGroup(input string) int {
	counts := MapGroup(input, CountUnique)
	return Sum(counts)
}

func SumIntersectInGroup(input string) int {
	counts := MapGroup(input, CountIntersect)
	return Sum(counts)
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
	result := SumIntersectInGroup(input)
	fmt.Println(result)
}
