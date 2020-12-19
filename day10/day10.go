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

// 1, 2, 3, 4
// 1, 2, 3
// 1, 2, 4
// 1, 3, 4
// 2, 3, 4
// 1, 2
// 1, 3
// 1, 4
// 2, 3
// 2, 4
// 3, 4
// 1
// 2
// 3
// 41
func MakeCombinations(items []int) {
	combinations := make([][]int, 0)
	for _, item := range items {
		combinations = append(combinations, []int{item})
	}
	fmt.Println(combinations)
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
func CountCombinations(items []int) (count int) {
	adapter := items[0]
	for i, item := range items[1:] {
		fmt.Printf("%d, %d, %d (%d)\n", i, adapter, item, count)
		if item > adapter+3 {
			fmt.Println("no good")
			return
		}
		if i == len(items) {
			fmt.Println("end")
			count++
			return
		}
		count += CountCombinations(items[i+1:])
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
	diffs := FindJoltDifferences(items)
	result := MultiplyEnds(diffs)
	fmt.Println(result)
}
