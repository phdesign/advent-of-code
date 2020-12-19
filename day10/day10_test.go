package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"testing"
)

func TestFindJoltDifferences(t *testing.T) {
	tests := []struct {
		input string
		want  []int
	}{
		{
			"16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4",
			[]int{7, 0, 5},
		},
		{
			"28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3",
			[]int{22, 0, 10},
		},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			items := Sort(Parse(test.input))
			got := FindJoltDifferences(items)
			assertIntSliceEqual(t, got, test.want)
		})
	}
}

func TestCountCombinations(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			"1\n4\n5\n7",
			2,
		},
		{
			"16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4",
			8,
		},
		{
			"28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3",
			19208,
		},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			items := Sort(Parse(test.input))
			got := MakeCombinationsRecursive(items, 0, items[len(items)-1])
			assertIntEqual(t, len(got), test.want)
		})
	}
}

func TestShortestPath(t *testing.T) {
	input := `1
4
5
7`
	items := Sort(Parse(input))
	got := ShortestPath(items)
	want := []int{1, 4, 7}
	assertIntSliceEqual(t, got, want)
}

func TestShortestPath2(t *testing.T) {
	input := "28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3"
	items := Sort(Parse(input))
	got := len(ShortestPath(items))
	long := math.Pow(2, float64(len(items)))
	short := math.Pow(2, float64(got))
	combinations := long - short
	fmt.Printf("combinations: %f", combinations)
	want := 19208
	assertIntEqual(t, got, want)
}

func TestSort(t *testing.T) {
	t.Run("should sort slice", func(t *testing.T) {
		input := []int{3, 2, 1}
		got := Sort(input)
		want := []int{1, 2, 3}
		assertIntSliceEqual(t, got, want)
	})

	t.Run("should not modify passed slice", func(t *testing.T) {
		input := []int{3, 2, 1}
		Sort(input)
		want := []int{3, 2, 1}
		assertIntSliceEqual(t, input, want)
	})
}

func TestFactorial(t *testing.T) {
	got := Factorial(4)
	want := 24
	assertIntEqual(t, got, want)
}

func TestMask(t *testing.T) {
	input := []int{1, 2, 3}
	tests := []struct {
		mask int
		want []int
	}{
		{1, []int{1}},
		{2, []int{2}},
		{3, []int{1, 2}},
		{4, []int{3}},
		{5, []int{1, 3}},
		{6, []int{2, 3}},
		{7, []int{1, 2, 3}},
	}
	for _, test := range tests {
		t.Run(strconv.Itoa(test.mask), func(t *testing.T) {
			got := Mask(input, test.mask)
			assertIntSliceEqual(t, got, test.want)
		})
	}
}

func TestCombinations(t *testing.T) {
	tests := []struct {
		n    int
		r    int
		want int
	}{
		{4, 4, 1},
		{4, 3, 4},
		{4, 2, 6},
		{4, 1, 4},
	}
	for _, test := range tests {
		name := fmt.Sprintf("%v", test)
		t.Run(name, func(t *testing.T) {
			got := Combinations(test.n, test.r)
			assertIntEqual(t, got, test.want)
		})
	}
}

func TestMakeCombination(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 63},
		{[]int{1, 2, 3, 4, 5}, 31},
		{[]int{1, 2, 3, 4}, 15},
		{[]int{1, 2, 3}, 7},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := MakeCombinations(test.input)
			assertIntEqual(t, len(got), test.want)
		})
	}
}

func TestMakeCombinationRecursive(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 63},
		{[]int{1, 2, 3, 4, 5}, 31},
		{[]int{1, 2, 3, 4}, 15},
		{[]int{1, 2, 3}, 7},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := MakeCombinationsRecursive(test.input, 0, test.input[len(test.input)-1])
			fmt.Println(got)
			assertIntEqual(t, len(got), test.want)
		})
	}
}

func assertIntEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Wanted %d. got %d", want, got)
	}
}

func assertIntSliceEqual(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}
