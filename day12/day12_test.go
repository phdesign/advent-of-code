package main

import "testing"

func TestNavigatePath(t *testing.T) {
	t.Run("", func(t *testing.T) {
		input := `F10
N3
F7
R90
F11`
		got := NavigatePath(Parse(input))
		want := Position{17, 8}
		assertPositionEqual(t, got, want)
	})
	t.Run("", func(t *testing.T) {
		input := `R90
F11`
		got := NavigatePath(Parse(input))
		want := Position{0, 11}
		assertPositionEqual(t, got, want)
	})
	t.Run("", func(t *testing.T) {
		input := `R180
F11`
		got := NavigatePath(Parse(input))
		want := Position{-11, 0}
		assertPositionEqual(t, got, want)
	})
	t.Run("", func(t *testing.T) {
		input := `R270
F11`
		got := NavigatePath(Parse(input))
		want := Position{0, -11}
		assertPositionEqual(t, got, want)
	})
	t.Run("", func(t *testing.T) {
		input := `L90
F11`
		got := NavigatePath(Parse(input))
		want := Position{0, -11}
		assertPositionEqual(t, got, want)
	})
}

func TestManhattanDistance(t *testing.T) {
	got := ManhattanDistance(Position{-17, 8})
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
