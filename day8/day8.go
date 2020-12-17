package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type InfiniteLoopError struct {
	instruction int
}

func (e *InfiniteLoopError) Error() string {
	return fmt.Sprintf("Infinite loop detected at instruction %d!", e.instruction)
}

type Instruction struct {
	operation string
	argument  int
}

func Parse(input string) (instructions []Instruction) {
	expressions := strings.Split(input, "\n")
	for _, expression := range expressions {
		tokens := strings.Fields(expression)
		if len(tokens) != 2 {
			panic(fmt.Sprintf("Unable to parse expression %q", expression))
		}
		operation := tokens[0]
		argument, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, Instruction{operation, argument})
	}
	return
}

func Evaluate(instructions []Instruction) (int, error) {
	pc := 0
	acc := 0
	call_stack := make([]int, 0)

	for pc < len(instructions) {
		// Check for infinite loops
		for _, prev := range call_stack {
			if prev == pc {
				return acc, &InfiniteLoopError{pc}
			}
		}
		call_stack = append(call_stack, pc)
		instruction := instructions[pc]
		switch instruction.operation {
		case "nop":
			pc++
		case "acc":
			acc += instruction.argument
			pc++
		case "jmp":
			pc += instruction.argument
		}

	}

	return acc, nil
}

func Debug(instructions []Instruction) int {
	for i, instruction := range instructions {
		var fix_op string
		switch instruction.operation {
		case "nop":
			fix_op = "jmp"
		case "jmp":
			fix_op = "nop"
		default:
			continue
		}

		fix := append([]Instruction{}, instructions...)
		fix[i] = Instruction{fix_op, instruction.argument}
		if result, err := Evaluate(fix); err == nil {
			return result
		}
	}
	panic("No way to fix the program!")
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
	instructions := Parse(input)
	result := Debug(instructions)
	fmt.Println(result)
}
