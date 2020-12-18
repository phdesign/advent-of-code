package main

import (
	"reflect"
	"testing"
)

func TestFindJoltDifferences(t *testing.T) {
	input := "16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4"
	items := Parse(input)
	got := FindJoltDifferences(items)
	want := []int{7, 0, 5}
	assertIntSliceEqual(t, got, want)
}

func assertIntSliceEqual(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}
