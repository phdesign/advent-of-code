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

func BinaryStringToInt(value string) int {
	i, err := strconv.ParseInt(value, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func translate(val string, dict map[rune]rune) string {
	return strings.Map(func(r rune) rune {
		if i := dict[r]; i == 0 {
			return r
		} else {
			return i
		}
	}, val)
}

func SeatId(route string) int {
	bin := translate(route, map[rune]rune{'B': '1', 'F': '0', 'R': '1', 'L': '0'})
	val := BinaryStringToInt(bin)
	return val
}

func HighestSeatId(routesStr string) (maxSeatId int) {
	routes := strings.Split(routesStr, "\n")
	for i, route := range routes {
		seatId := SeatId(route)
		if i == 0 || seatId > maxSeatId {
			maxSeatId = seatId
		}
	}
	return
}

func EmptySeatId(routesStr string) int {
	routes := strings.Split(routesStr, "\n")
	seatIds := make([]int, len(routes))
	for i, route := range routes {
		seatIds[i] = SeatId(route)
	}
	sort.Ints(seatIds)
	for i, seatId := range seatIds {
		if seatIds[0]+i != seatId {
			return seatId - 1
		}
	}
	panic("No empty seat found!")
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
	result := EmptySeatId(input)
	fmt.Println(result)
}
