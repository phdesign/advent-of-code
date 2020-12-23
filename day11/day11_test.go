package main

import (
	"strings"
	"testing"
)

func TestCountOccupiedAdjacentSeats(t *testing.T) {
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
	got := CountOccupiedSeats(input, 4, CountAdjacent)
	want := 37
	assertIntEqual(t, got, want)
}

func TestCountOccupiedLineOfSightSeats(t *testing.T) {
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
	got := CountOccupiedSeats(input, 5, CountLineOfSight)
	want := 26
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
		got := Evaluate(input, 10, 4, CountAdjacent)
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

func TestCountLineOfSight(t *testing.T) {
	t.Run("", func(t *testing.T) {
		input := strings.ReplaceAll(`##LL.LL.LL
L..#LLL.LL
..L.L..L..
#LLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`, "\n", "")
		got := CountLineOfSight(10, 10, input)
		want := 4
		assertIntEqual(t, got, want)
	})

	t.Run("", func(t *testing.T) {
		input := strings.ReplaceAll(`.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`, "\n", "")
		got := CountLineOfSight(39, 9, input)
		want := 8
		assertIntEqual(t, got, want)
	})

	t.Run("", func(t *testing.T) {
		input := strings.ReplaceAll(`.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.`, "\n", "")
		got := CountLineOfSight(24, 7, input)
		want := 0
		assertIntEqual(t, got, want)
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
