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

func Factorial(num int) int {
	f := 1
	for i := 1; i <= num; i++ {
		f = f * i
	}
	return f
}

func Combinations(n int, r int) int {
	return Factorial(n) / (Factorial(r) * Factorial(n-r))
}

func Mask(items []int, mask int) (result []int) {
	for i, item := range items {
		if (1<<i)&mask != 0 {
			result = append(result, item)
		}
	}
	return
}

func MakeCombinations(items []int) [][]int {
	length := 1 << len(items)
	combinations := make([][]int, 0)
	for i := 1; i < length; i++ {
		combination := Mask(items, i)
		combinations = append(combinations, combination)
	}
	return combinations
}

func MakeCombinationsRecursive(items []int, depth int, last int) (combinations [][]int) {
	if depth >= len(items) {
		return
	}
	if depth == 0 && items[depth] > 3 {
		return
	}
	if depth > 0 && items[depth]-items[depth-1] > 3 {
		return
	}

	other := MakeCombinationsRecursive(items, depth+1, last)
	combinations = append(combinations, other...)

	skip := make([]int, depth)
	copy(skip, items[:depth])
	skip = append(skip, items[depth+1:]...)
	other = MakeCombinationsRecursive(skip, depth, last)
	combinations = append(combinations, other...)

	if items[depth] == last {
		combinations = append(combinations, items[:depth+1])
	}
	return
}

func ShortestPath(sorted []int) []int {
	i := 1
	for i < len(sorted)-1 {
		difference := sorted[i+1] - sorted[i-1]
		if difference <= 3 {
			sorted = append(sorted[:i], sorted[i+1:]...)
		} else {
			i++
		}
	}
	return sorted
}

// (0), 1, 4, 5, 7, (10)
// (0), 1, 4, 7, (10)
func CountCombinations(sorted []int) (count int) {
	shortestPath := ShortestPath(sorted)
	from := len(shortestPath)
	to := len(sorted)
	for i := from; i < to; i++ {
		count += Combinations(to, i)
	}
	return
}

func MultiplyEnds(items []int) int {
	return items[0] * items[len(items)-1]
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
	combinations := MakeCombinationsRecursive(items, 0, items[len(items)-1])
	result := len(combinations)
	fmt.Println(result)
}
