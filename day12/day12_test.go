package main

import "testing"

func TestNavigatePath(t *testing.T) {
	input := `F10
N3
F7
R90
F11`
	got := NavigatePath(Parse(input))
	want := Position{17, 8}
	assertPositionEqual(t, got, want)
}

func TestManhattanDistance(t *testing.T) {
	got := int(ManhattanDistance(Position{-17, 8}))
	want := 25
	assertIntEqual(t, got, want)
}

func assertPositionEqual(t *testing.T, got, want Position) {
	t.Helper()
	if got.x != want.x || got.y != want.y {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func assertIntEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}
