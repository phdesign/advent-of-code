package main

import "testing"

func TestValidatePasswords(t *testing.T) {
    actual := ValidatePasswords()
    expected := 2

    if actual != expected {
        t.Errorf("Expected %d but got %d", expected, actual)
    }
}
