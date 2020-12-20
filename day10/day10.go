package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
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

func Sort(items []int) []int {
	sorted := make([]int, len(items))
	copy(sorted, items)
	sort.Ints(sorted)
	return sorted
}

func MultiplyEnds(items []int) int {
	return items[0] * items[len(items)-1]
}

func FindJoltDifferences(sorted []int) []int {
	diffCount := make([]int, 3)
	for i, item := range sorted {
		var prev int
		if i == 0 {
			prev = 0
		} else {
			prev = sorted[i-1]
		}
		difference := item - prev
		diffCount[difference-1]++
	}
	// Last difference is always 3
	diffCount[2]++
	return diffCount
}

func CountCombinations(sorted []int) int {
	counts := map[int]int{0: 1}
	last := sorted[len(sorted)-1] + 3
	items := append(sorted, last)
	for _, item := range items {
		counts[item] = counts[item-3] + counts[item-2] + counts[item-1]
	}
	return counts[last]
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
	items := Sort(Parse(input))
	result := CountCombinations(items)
	fmt.Println(result)
}
