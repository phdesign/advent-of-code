package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type Action struct {
	action rune
	value  int
}

func Parse(input string) (path []Action) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		action := rune(line[0])
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		path = append(path, Action{action, value})
	}
	return
}

func addAngle(a, b float64) float64 {
	result := a + b
	if result < 0 {
		result = 360 + result
	} else if result >= 360 {
		result = result - 360
	}
	return result
}

func NavigatePath(path []Action) Position {
	direction := 0.0
	pos := Position{}
	for _, step := range path {
		valueFloat := float64(step.value)
		switch step.action {
		case 'N':
			pos.y -= step.value
		case 'S':
			pos.y += step.value
		case 'E':
			pos.x += step.value
		case 'W':
			pos.x -= step.value
		case 'L':
			direction = addAngle(direction, -valueFloat)
		case 'R':
			direction = addAngle(direction, valueFloat)
		case 'F':
			radians := direction * math.Pi / 180
			pos.x += int(math.Round(valueFloat * math.Cos(radians)))
			pos.y += int(math.Round(valueFloat * math.Sin(radians)))
		}
	}
	return pos
}

func ManhattanDistance(pos Position) int {
	return int(math.Round(math.Abs(float64(pos.x)) + math.Abs(float64(pos.y))))
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
	path := Parse(input)
	finalPosition := NavigatePath(path)
	fmt.Println(finalPosition)
	result := ManhattanDistance(finalPosition)
	fmt.Println(result)
}
