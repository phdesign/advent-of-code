package main

import (
	"fmt"
	"math"
)

type Range struct {
	lower int
	upper int
}

func Traverse(path string, size int) int {
	r := Range{0, size - 1}
	for _, direction := range path {
		mid := float64(r.lower) + (float64(r.upper - r.lower) / 2)
		switch direction {
		case '0':
			r.upper = int(math.Floor(mid))
		case '1':
			r.lower = int(math.Ceil(mid))
		}
	}
	if r.lower != r.upper {
		panic(fmt.Sprintf("Didn't settle on a result, range: %v", r))
	}
	return r.lower
}
