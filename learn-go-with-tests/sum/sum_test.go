package sum

import (
    "testing"
    "reflect"
)

func TestSum(t *testing.T) {
    numbers := []int{1, 2, 3, 4, 5}

    got := Sum(numbers)
    want := 15

    if got != want {
        t.Errorf("got %d want %d given, %v", got, want, numbers)
    }
}

func TestSumAll(t *testing.T) {
    got := SumAll([]int{1, 2}, []int{0,9})
    want := []int{3, 9}

    if !reflect.DeepEqual(got, want) {
        t.Errorf("want %v, got %v", want, got)
    }
}

func TestSumAllTails(t *testing.T) {
    assertEqual := func(t *testing.T, got, want []int) {
        t.Helper()

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    }

    t.Run("should sum tail of slices", func(t *testing.T) {
        got := SumAllTails([]int{1, 2}, []int{0, 9})
        want := []int{2, 9}
        assertEqual(t, got, want)
    })

    t.Run("should handle suming tail of empty slice", func(t *testing.T) {
        got := SumAllTails([]int{}, []int{3, 4, 5})
        want := []int{0, 9}
        assertEqual(t, got, want)
    })
}
