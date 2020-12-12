package main

import "testing"

func TestBinaryStringToInt(t *testing.T) {
	tests := []struct {
		value string
		want int
	}{
		{"101", 5},
		{"100", 4},
		{"010", 2},
		{"011", 3},
	}

	for _, test := range tests {
		t.Run(test.value, func (t *testing.T)  {
			got := BinaryStringToInt(test.value)
			assertIntEqual(t, test.want, got)
		})
	}
}

func TestSeatId(t *testing.T) {
	tests := []struct {
		route string
		want int
	}{
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, test := range tests {
		t.Run(test.route, func (t *testing.T)  {
			got := SeatId(test.route)
			assertIntEqual(t, test.want, got)
		})
	}

}

func TestHighestSeatId(t *testing.T) {
	input := `BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`
	got := HighestSeatId(input)
	want := 820

	assertIntEqual(t, want, got)
}

func TestEmptySeatId(t *testing.T) {
	input := `FFFFFFBLLL
FFFFFFBLLR
FFFFFFBLRL
FFFFFFBRLL`
	got := EmptySeatId(input)
	want := 11

	assertIntEqual(t, want, got)
}

func assertIntEqual(t *testing.T, want, got int) {
	t.Helper()
	if want != got {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}
