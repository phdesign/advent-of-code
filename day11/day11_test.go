package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCountOccupiedSeats(t *testing.T) {
	input := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
	got := CountOccupiedSeats(input)
	want := 37
	assertIntEqual(t, got, want)
}

func TestEvaluate(t *testing.T) {
	t.Run("should occupy every seat given no seats occupied", func(t *testing.T) {
		input := strings.ReplaceAll(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`, "\n", "")
		got := Evaluate(input, 10)
		want := strings.ReplaceAll(`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`, "\n", "")
		assertStringEqual(t, got, want)
	})
}

func TestMarkAdjacent(t *testing.T) {
	t.Run("should not mark seats before start of row", func(t *testing.T) {
		got := make(map[int]int)
		MarkAdjacent(10, 10, got)
		want := map[int]int{0: 1, 1: 1, 11: 1, 20: 1, 21: 1}
		assertMapEqual(t, got, want)
	})
	t.Run("should not mark seats after end of row", func(t *testing.T) {
		got := make(map[int]int)
		MarkAdjacent(19, 10, got)
		want := map[int]int{8: 1, 9: 1, 18: 1, 28: 1, 29: 1}
		assertMapEqual(t, got, want)
	})
	t.Run("should update existing values", func(t *testing.T) {
		got := map[int]int{1: 2}
		MarkAdjacent(0, 10, got)
		want := 3
		assertIntEqual(t, got[1], want)
	})
}

func assertIntEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}

func assertStringEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Wanted %q, got %q", want, got)
	}
}

func assertMapEqual(t *testing.T, got, want map[int]int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}
