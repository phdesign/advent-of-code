package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func toInts(items []string) (result []int) {
	for _, str := range items {
		i, err := strconv.Atoi(str)
		if err != nil {
			panic(fmt.Sprintf("Unable to parse number %q", str))
		}
		result = append(result, i)
	}
	return
}

func Parse(input string) []int {
	fields := strings.Split(input, "\n")
	return toInts(fields)
}

func FindJoltDifferences(items []int) []int {
	sort.Ints(items)
	diffCount := make([]int, 3)
	for i, item := range items {
		var prev int
		if i == 0 {
			prev = 0
		} else {
			prev = items[i-1]
		}
		difference := item - prev
		fmt.Printf("%d - %d = %d\n", item, prev, difference)
		diffCount[difference-1]++
	}
	// Last difference is always 3
	diffCount[2]++
	return diffCount
}
