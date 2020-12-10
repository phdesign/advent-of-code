package main

import (
    "testing"
    "reflect"
)

func TestValidatePasswords(t *testing.T) {
    passwords := []PasswordPolicy{
        {1, 3, 'a', "abcde"},
        {1, 3, 'b', "cdefg"},
        {2, 9, 'c', "ccccccccc"},
    }
    actual := ValidatePasswords(passwords)
    expected := 2

    if actual != expected {
        t.Errorf("Expected %d but got %d", expected, actual)
    }
}

func TestCountOccurrences(t *testing.T) {
    actual := CountOccurrences("abac", 'a')
    expected := 2

    if actual != expected {
        t.Errorf("Expected %d but got %d", expected, actual)
    }
}

func TestParsePasswords(t *testing.T) {
    lines := []string {
        "1-3 a: abcde",
        "1-3 b: cdefg",
        "8-9 n: nnnnnnnnn",
    }
    actual := ParsePasswords(lines)
    expected := []PasswordPolicy{
        {1, 3, 'a', "abcde"},
        {1, 3, 'b', "cdefg"},
        {8, 9, 'n', "nnnnnnnnn"},
    }

    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("Expected %v but got %v", expected, actual)
    }
}

