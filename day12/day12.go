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
	x float64
	y float64
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
			pos.y -= valueFloat
		case 'S':
			pos.y += valueFloat
		case 'E':
			pos.x += valueFloat
		case 'W':
			pos.x -= valueFloat
		case 'L':
			direction = addAngle(direction, valueFloat)
		case 'R':
			direction = addAngle(direction, +valueFloat)
		case 'F':
			radians := direction * math.Pi / 180
			pos.x += valueFloat * math.Cos(radians)
			pos.y += valueFloat * math.Sin(radians)
		}
		fmt.Println(pos)
	}
	return pos
}

func ManhattanDistance(pos Position) float64 {
	return math.Abs(pos.x) + math.Abs(pos.y)
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
	fmt.Println(math.Round(result))
}
