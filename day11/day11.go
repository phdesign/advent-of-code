package main

import (
	"fmt"
	"strings"
)

func MarkAdjacent(i int, width int, adjacent map[int]int) {
	adjacent[i+width]++
	adjacent[i-width]++
	// If not at the end of row
	if (i+1)%width != 0 {
		adjacent[i+1]++
		adjacent[i+width+1]++
		adjacent[i-width+1]++
	}
	// If not at the start of row
	if i%width != 0 {
		adjacent[i-1]++
		adjacent[i+width-1]++
		adjacent[i-width-1]++
	}
}

func CountAdjacent(i int, width int, adjacent map[int]int) (count int) {
	count += adjacent[i+width]
	count += adjacent[i-width]
	// If not at the end of row
	if (i+1)%width != 0 {
		count += adjacent[i+1]
		count += adjacent[i+width+1]
		count += adjacent[i-width+1]
	}
	// If not at the start of row
	if i%width != 0 {
		count += adjacent[i-1]
		count += adjacent[i+width-1]
		count += adjacent[i-width-1]
	}
	return
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func Evaluate(seats string, width int) string {
	adjacent := make(map[int]int, 0)
	for i, seat := range seats {
		if seat == '#' {
			MarkAdjacent(i, width, adjacent)
		}
	}
	for i, seat := range seats {
		adjacentCount := CountAdjacent(i, width, adjacent)
		switch seat {
		case 'L':
			if adjacentCount == 0 {
				seats = replaceAtIndex(seats, '#', i)
			}
		case '#':
			if adjacentCount >= 4 {
				seats = replaceAtIndex(seats, 'L', i)
			}
		}
	}
	return seats
}

func CountOccupiedSeats(input string) int {
	width := strings.Index(input, "\n")
	seats := strings.ReplaceAll(input, "\n", "")
	settled := Evaluate(seats, width)
	for {
		next := Evaluate(settled, width)
		if next == settled {
			break
		}
		settled = next
		fmt.Println(settled)
	}
	return strings.Count(settled, "#")
}
