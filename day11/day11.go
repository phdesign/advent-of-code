package main

import (
	"fmt"
	"strings"
)

func CountOccupiedSeats(input string) int {
	width := strings.Index(input, "\n")
	seats := strings.ReplaceAll(input, "\n", "")
	adjacent := make(map[int]int, 0)
	for i, seat := range seats {
		if seat == '#' {
			adjacent[i+width]++
			adjacent[i-width]++
			if (i+1)%width == 0 {
				adjacent[i+1]++
				adjacent[i+width+1]++
				adjacent[i-width+1]++
			}
			if (i+1)%width == 1 {
				adjacent[i-1]++
				adjacent[i+width-1]++
				adjacent[i-width-1]++
			}
		}
	}
	fmt.Println(adjacent)
	return 0
}
