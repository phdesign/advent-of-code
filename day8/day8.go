package main

import (
	"fmt"
	"strconv"
	"strings"
)

//type Context struct {
//accumulator int
//pc          int
//}

//type Noop struct{}

//func (n *Noop) Interpret(context *Context) {
//}

//type Accumulate struct {
//argument int
//}

//func (a *Accumulate) Interpret(context *Context) {
//context.accumulator += a.argument
//}

//func NewAccumulate(argument int) *Accumulate {
//return &Accumulate{argument: argument}
//}

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
