package main

import "testing"

func TestCountTrees(t *testing.T) {
    grid := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

    expected := 7
    actual := CountTrees(grid, 1, 3)

    if actual != expected  {
        t.Errorf("Expected %d but got %d", expected, actual)
    }
}
