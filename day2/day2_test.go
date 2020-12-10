package main

import (
    "testing"
    "reflect"
)

func TestValidatePasswords(t *testing.T) {
    check := func (t *testing.T, policy Policy, expected int, passwords []PasswordPolicy) {
        t.Helper()
        actual := ValidatePasswords(passwords, policy)

        if actual != expected {
            t.Errorf("Expected %d but got %d", expected, actual)
        }
    }

    t.Run("should count valid passwords given occurrence policy", func(t *testing.T) {
        passwords := []PasswordPolicy{
            {1, 3, 'a', "abcde"},
            {1, 3, 'b', "cdefg"},
            {2, 9, 'c', "ccccccccc"},
        }
        check(t, OccurrencePolicy, 2, passwords)
    })

    t.Run("should count valid passwords given position policy", func(t *testing.T) {
        passwords := []PasswordPolicy{{1, 3, 'a', "abcde"}}
        check(t, PositionPolicy, 1, passwords)
    })

    t.Run("should fail if no match given position policy", func(t *testing.T) {
        passwords := []PasswordPolicy{{1, 3, 'b', "cdefg"}}
        check(t, PositionPolicy, 0, passwords)
    })

    t.Run("should fail when both positions match given position policy", func(t *testing.T) {
        passwords := []PasswordPolicy{{2, 9, 'c', "ccccccccc"}}
        check(t, PositionPolicy, 0, passwords)
    })
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

