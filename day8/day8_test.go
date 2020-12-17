package main

import (
	//"reflect"
	"testing"
)

func TestParseAndEvaluate(t *testing.T) {
	input := `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`
	instructions := Parse(input)
	got, err := Evaluate(instructions)
	want := 5

	//assertError(t, err, reflect.TypeOf(InfiniteLoopError{}))
	if _, ok := err.(*InfiniteLoopError); !ok {
		t.Error("Wanted InfiniteLoopError, got none")
	}
	assertIntEqual(t, got, want)
}

func TestEvaluate(t *testing.T) {
	instructions := []Instruction{
		Instruction{"nop", 0},
	}
	got, err := Evaluate(instructions)
	want := 0

	assertNoError(t, err)
	assertIntEqual(t, got, want)
}

//func assertError(t *testing.T, err error, errorType reflect.Type) {
//t.Helper()
//if _, ok := err.(*errorType); !ok {
//t.Errorf("Wanted %v error type, got %v", errorType, err)
//}
//}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Wanted no error, got %v", err)
	}
}

func assertIntEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}
