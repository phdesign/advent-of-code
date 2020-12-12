package main

import "testing"

func TestTraverse(t *testing.T) {
	tests := []struct {
		path string
		want int
	}{
		{"101", 5},
		{"100", 4},
		{"010", 2},
		{"011", 3},
	}

	for _, test := range tests {
		t.Run(test.path, func (t *testing.T)  {
			got := Traverse(test.path, 8)
			want := test.want

			assertIntEqual(t, want, got)
		})
	}
}

func assertIntEqual(t *testing.T, want, got int) {
	t.Helper()
	if want != got {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}
