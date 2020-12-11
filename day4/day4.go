package main

import (
    "strings"
    //"fmt"
)

type Passport struct {
    birthYear string
    issueYear string
    expiryYear string
    height string
    hairColor string
    eyeColor string
    passportId string
    countryId string
}

func (p Passport) isValid() bool {
    if p.birthYear == "" {
        return false
    }
    if p.issueYear == "" {
        return false
    }
    if p.expiryYear == "" {
        return false
    }
    if p.height == "" {
        return false
    }
    if p.hairColor == "" {
        return false
    }
    if p.eyeColor == "" {
        return false
    }
    if p.passportId == "" {
        return false
    }
    return true
}

func ParseFields(input string) map[string]string {
    fields := strings.Fields(input)
    m := make(map[string]string)
    for _, field := range fields {
        parts := strings.Split(field, ":")
        m[parts[0]] = parts[1]
    }
    return m
}

func ParsePassports(input string) []Passport {
    paragraphs := strings.Split(input, "\n\n")
    passports := make([]Passport, len(paragraphs))
    for i, paragraph := range paragraphs {
        fields := ParseFields(paragraph)
        passports[i] = Passport{
            birthYear: fields["byr"],
            issueYear: fields["iyr"],
            expiryYear: fields["eyr"],
            height: fields["hgt"],
            hairColor: fields["hcl"],
            eyeColor: fields["ecl"],
            passportId: fields["pid"],
            countryId: fields["cid"],
        }
    }
    return passports
}

func ValidatePasswords(passports []Passport) (count int) {
    for _, passport := range passports {
        if passport.isValid() {
            count++
        }
    }
    return
}
