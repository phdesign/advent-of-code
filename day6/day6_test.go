package main

import (
	"reflect"
	"testing"
)

func TestSumUniqueInGroup(t *testing.T) {
	input := `abc

a
b
c

ab
ac

a
a
a
a

b`
	got := SumUniqueInGroup(input)
	want := 11
	assertIntEqual(t, got, want)
}

func TestSumIntersectInGroup(t *testing.T) {
	input := `abc

a
b
c

ab
ac

a
a
a
a

b`
	got := SumIntersectInGroup(input)
	want := 6
	assertIntEqual(t, got, want)
}

func assertIntEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}

func assertSliceEqual(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v wanted %v", got, want)
	}
}
