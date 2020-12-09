package main

import "testing"

func TestSolve(t *testing.T) {
    t.Run("should return product of numbers that sum 2020", func(t *testing.T) {
        tests := []struct {
            expenseReport []int
            expected int
        }{
            {expenseReport: []int{1721, 979, 366, 299, 675, 1456}, expected: 514579},
            {expenseReport: []int{2019, 1}, expected: 2019},
        }

        for _, test := range tests {
            actual, err := Solve(test.expenseReport)

            if err != nil {
                t.Errorf("Unexpected error: %v", err)
            }
            if actual != test.expected {
                t.Errorf("Expected %d but got %d", test.expected, actual)
            }
        }
    })

    t.Run("should raise error if no numbers sum 2020", func(t *testing.T) {
        expenseReport := []int{1, 2, 3}
        _, err := Solve(expenseReport)

        if err == nil {
            t.Error("Expected error to be raised, but didn't see one")
        }
    })
}
