package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func CountAdjacent(i int, width int, seats string) (count int) {
	adjacents := []int{i + width, i - width}
	// If not at the end of row
	if (i+1)%width != 0 {
		adjacents = append(adjacents, i+1)
		adjacents = append(adjacents, i+width+1)
		adjacents = append(adjacents, i-width+1)
	}
	// If not at the start of row
	if i%width != 0 {
		adjacents = append(adjacents, i-1)
		adjacents = append(adjacents, i+width-1)
		adjacents = append(adjacents, i-width-1)
	}
	for _, loc := range adjacents {
		if loc < 0 || loc >= len(seats) {
			continue
		}
		if seats[loc] == '#' {
			count++
		}
	}
	return
}

func Evaluate(seats string, width int) (result string) {
	for i, seat := range seats {
		adjacentCount := CountAdjacent(i, width, seats)
		switch seat {
		case 'L':
			if adjacentCount == 0 {
				result += "#"
			} else {
				result += string(seat)
			}
		case '#':
			if adjacentCount >= 4 {
				result += "L"
			} else {
				result += string(seat)
			}
		default:
			result += string(seat)
		}
	}
	return
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
	}
	return strings.Count(settled, "#")
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
	result := CountOccupiedSeats(input)
	fmt.Println(result)
}
