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
			items := Parse(test.input)
			got := FindJoltDifferences(items)
			assertIntSliceEqual(t, got, test.want)
		})
	}
}

func assertIntSliceEqual(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}
