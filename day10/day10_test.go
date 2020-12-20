package main

import (
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

func TestPart2(t *testing.T) {
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
			got := CountCombinations(items, 0, items[len(items)-1])
			assertIntEqual(t, got, test.want)
		})
	}
}

func TestCountCombinations(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 24},
		{[]int{1, 2, 3, 4, 5}, 13},
		{[]int{1, 2, 3, 4}, 7},
		{[]int{1, 2, 3}, 4},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := CountCombinations(test.input, 0, test.input[len(test.input)-1])
			assertIntEqual(t, got, test.want)
		})
	}
}

func BenchmarkCountCombinations(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < b.N; i++ {
		CountCombinations(input, 0, input[len(input)-1])
	}
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
