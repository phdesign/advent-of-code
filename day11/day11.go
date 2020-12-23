package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Counter func(int, int, string) int

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

func look(from int, width int, seats string, dx int, dy int) int {
	i := 1
	for {
		x := i * dx
		y := i * dy

		if x < 0 || x >= width {
			return 0
		}
		loc := i + (x + (y * width))
		if loc < 0 || loc >= len(seats) {
			return 0
		}
		switch seats[loc] {
		case '#':
			return 1
		case 'L':
			return 0
		}
		i++
	}
}

func CountLineOfSight(i int, width int, seats string) (count int) {
	// N
	for loc := i - width; loc >= 0; loc -= width {
		if seats[loc] == '#' {
			//fmt.Printf("North at %d\n", loc)
			count++
			break
		} else if seats[loc] == 'L' {
			break
		}
	}
	// S
	for loc := i + width; loc < len(seats); loc += width {
		if seats[loc] == '#' {
			//fmt.Printf("South at %d\n", loc)
			count++
			break
		} else if seats[loc] == 'L' {
			break
		}
	}
	// E
	for loc := i + 1; loc/width == i/width && loc < len(seats); loc += 1 {
		if seats[loc] == '#' {
			//fmt.Printf("East at %d\n", loc)
			count++
			break
		} else if seats[loc] == 'L' {
			break
		}
	}
	// W
	for loc := i - 1; loc/width == i/width && loc >= 0; loc -= 1 {
		if seats[loc] == '#' {
			//fmt.Printf("West at %d\n", loc)
			count++
			break
		} else if seats[loc] == 'L' {
			break
		}
	}
	// NE
	for loc := i - (width + 1); loc >= 0; loc -= (width + 1) {
		if seats[loc] == '#' {
			//fmt.Printf("North east at %d\n", loc)
			count++
			break
		} else if seats[loc] == 'L' {
			break
		}
	}
	// NW
	for loc := i - (width + 1); loc >= 0; loc -= (width + 1) {
		if seats[loc] == '#' {
			//fmt.Printf("North west at %d\n", loc)
			count++
			break
		} else if seats[loc] == 'L' {
			break
		}
	}
	// SE
	for loc := i + (width + 1); loc < len(seats); loc += (width + 1) {
		if seats[loc] == '#' {
			//fmt.Printf("South east at %d\n", loc)
			count++
			break
		} else if seats[loc] == 'L' {
			break
		}
	}
	// SW
	for loc := i + (width - 1); loc < len(seats); loc += (width - 1) {
		if seats[loc] == '#' {
			//fmt.Printf("South west at %d\n", loc)
			count++
			break
		} else if seats[loc] == 'L' {
			break
		}
	}
	return
}

func Evaluate(seats string, width int, occupiedLimit int, countOccupied Counter) (result string) {
	for i, seat := range seats {
		occupiedCount := countOccupied(i, width, seats)
		switch seat {
		case 'L':
			if occupiedCount == 0 {
				result += "#"
			} else {
				result += string(seat)
			}
		case '#':
			if occupiedCount >= occupiedLimit {
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

func CountOccupiedSeats(input string, occupiedLimit int, countOccupied Counter) int {
	width := strings.Index(input, "\n")
	seats := strings.ReplaceAll(input, "\n", "")
	settled := Evaluate(seats, width, occupiedLimit, countOccupied)
	for {
		next := Evaluate(settled, width, occupiedLimit, countOccupied)
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
	result := CountOccupiedSeats(input, 4, CountAdjacent)
	fmt.Println(result)
}
