package main

import (
    "errors"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
    "strconv"
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

func findSum(values []int, sum int) ([2]int, error) {
    for i, val := range values {
        rest := values[i:]
        for _, other := range rest {
            if (val + other) == sum {
                return [2]int{val, other}, nil
            }
        }
    }
    return [2]int{}, errors.New(fmt.Sprintf("No values in %v that sum to %d", values, sum))
}

func Solve(expenseReport []int) (int, error) {
    if entries, err := findSum(expenseReport, sum); err == nil {
        return entries[0] * entries[1], nil
    } else {
        return 0, err
    }
}

func main() {
    filename := os.Args[1]
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

    result, err := Solve(expenseReport)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
}
