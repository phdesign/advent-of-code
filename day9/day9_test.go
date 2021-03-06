package main

import "testing"

func TestHack1(t *testing.T) {
	input := `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`
	preambleLength := 5
	got := Hack1(input, preambleLength)
	want := 127

	assertIntEqual(t, got, want)
}

func TestHack2(t *testing.T) {
	input := `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`
	preambleLength := 5
	got := Hack2(input, preambleLength)
	want := 62

	assertIntEqual(t, got, want)
}

func assertIntEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}
