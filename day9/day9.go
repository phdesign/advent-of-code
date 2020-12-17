package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func stringsToInts(items []string) (result []int) {
	for _, str := range items {
		i, err := strconv.Atoi(str)
		if err != nil {
			panic(fmt.Sprintf("Unable to parse number %q", str))
		}
		result = append(result, i)
	}
	return
}

func sumAll(values []int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum
}

func find2ThatSum(values []int, sum int) ([]int, bool) {
	for i, val := range values {
		rest := values[i:]
		for _, other := range rest {
			if (val + other) == sum {
				return []int{val, other}, true
			}
		}
	}
	return []int{}, false
}

func findWeakness(items []int, preambleLength int) int {
	for i := preambleLength; i < len(items); i++ {
		item := items[i]
		preamble := items[i-preambleLength : i]
		if _, ok := find2ThatSum(preamble, item); !ok {
			return item
		}
	}
	panic("No weakness found")
}

func Hack(input string, preambleLength int) int {
	items := strings.Split(input, "\n")
	numbers := stringsToInts(items)
	return findWeakness(numbers, preambleLength)
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
	result := Hack(input, 25)
	fmt.Println(result)
}
