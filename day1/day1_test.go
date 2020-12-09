package main

import "testing"

func TestSolveFor2(t *testing.T) {
    t.Run("should return product of numbers that sum 2020", func(t *testing.T) {
        tests := []struct {
            expenseReport []int
            expected int
        }{
            {expenseReport: []int{1721, 979, 366, 299, 675, 1456}, expected: 514579},
            {expenseReport: []int{2019, 1}, expected: 2019},
        }

        for _, test := range tests {
            actual, err := Solve(test.expenseReport, 2)

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
        _, err := Solve(expenseReport, 2)

        if err == nil {
            t.Error("Expected error to be raised, but didn't see one")
        }
    })
}

func TestSolveFor3(t *testing.T) {
    expenseReport := []int{1721, 979, 366, 299, 675, 1456}
    expected := 241861950
    actual, err := Solve(expenseReport, 3)

    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if actual != expected {
        t.Errorf("Expected %d but got %d", expected, actual)
    }
}
