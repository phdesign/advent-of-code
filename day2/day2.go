package main

type PasswordPolicy struct {
    min int
    max int
    char rune
    password string
}

func ValidatePasswords(passwords []PasswordPolicy) int {
    return 0
}
