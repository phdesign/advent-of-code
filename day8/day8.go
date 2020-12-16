package main

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

type Instruction struct {
	operation string
	argument  int
}

func RunProgram(input string) int {
	instructions := Parse(input)
	pc := 0
	acc := 0
	call_stack := make([]int, 0)

	for {
		instruction := instructions[pc]
		pc++
		switch instruction.operation {
		case "nop":
			break
		case "acc":
			acc += instruction.argument
		case "jmp":
			pc += instruction.argument
		}

	}

	return 0
}
