package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const sum = 2020

func stringsToInts(strings []string) ([]int, error) {
	ints := make([]int, 0, len(strings))
	for _, str := range strings {
		i, err := strconv.Atoi(str)
		if err != nil {
			return ints, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}

func multiplyAll(values []int) int {
	sum := values[0]
	for i := 1; i < len(values); i++ {
		sum *= values[i]
	}
	return sum
}

func sumAll(values []int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum
}

func findSumOf2(values []int, sum int) ([]int, error) {
	for i, val := range values {
		rest := values[i:]
		for _, other := range rest {
			if (val + other) == sum {
				return []int{val, other}, nil
			}
		}
	}
	return []int{}, errors.New(fmt.Sprintf("No values in %v that sum to %d", values, sum))
}

func findSumOf3(values []int, sum int) ([]int, error) {
	for i := 0; i < len(values); i++ {
		for j := 1; j < len(values); j++ {
			for k := 2; k < len(values); k++ {
				result := []int{values[i], values[j], values[k]}
				if sumAll(result) == sum {
					return result, nil
				}
			}
		}
	}
	return []int{}, errors.New(fmt.Sprintf("No values in %v that sum to %d", values, sum))
}

func Solve(expenseReport []int, numOfEntries int) (int, error) {
	var entries []int
	var err error
	if numOfEntries == 2 {
		entries, err = findSumOf2(expenseReport, sum)
	} else if numOfEntries == 3 {
		entries, err = findSumOf3(expenseReport, sum)
	} else {
		return 0, errors.New(fmt.Sprintf("Invalid numOfEntries argument %d", numOfEntries))
	}

	if err == nil {
		return multiplyAll(entries[:]), nil
	} else {
		return 0, err
	}
}

func main() {
	numOfEntries := flag.Int("n", 2, "Number of entries to find (3 or 4)")
	flag.Parse()
	fmt.Println(*numOfEntries)

	filename := flag.Arg(0)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Fields(string(content))
	expenseReport, err := stringsToInts(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := Solve(expenseReport, *numOfEntries)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
