package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"strconv"
)

type Passport struct {
	birthYear  string
	issueYear  string
	expiryYear string
	height     string
	hairColor  string
	eyeColor   string
	passportId string
	countryId  string
}

func (p Passport) isValid(validator Validator) bool {
	return validator(p)
}

type Validator func(Passport) bool

func RequiredFieldValidator(p Passport) bool {
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

func FieldValueValidator(p Passport) bool {
	if !RequiredFieldValidator(p) {
		return false
	}
	if len(p.birthYear) != 4 {
		return false
	}
	if v, _ := strconv.Atoi(p.birthYear); v < 1920 || v > 2002 {
		return false
	}
	if len(p.issueYear) != 4 {
		return false
	}
	if v, _ := strconv.Atoi(p.issueYear); v < 2010 || v > 2020 {
		return false
	}
	if len(p.expiryYear) != 4 {
		return false
	}
	if v, _ := strconv.Atoi(p.expiryYear); v < 2020 || v > 2030 {
		return false
	}

	if len(p.passportId) != 9 {
		return false
	}
	if _, err := strconv.Atoi(p.passportId); err != nil {
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
			birthYear:  fields["byr"],
			issueYear:  fields["iyr"],
			expiryYear: fields["eyr"],
			height:     fields["hgt"],
			hairColor:  fields["hcl"],
			eyeColor:   fields["ecl"],
			passportId: fields["pid"],
			countryId:  fields["cid"],
		}
	}
	return passports
}

func ValidatePassports(passports []Passport, validator Validator) (count int) {
	for _, passport := range passports {
		if passport.isValid(validator) {
			count++
		}
	}
	return
}

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	input := strings.Trim(string(content), "\n")
	passports := ParsePassports(input)
	result := ValidatePassports(passports, RequiredFieldValidator)
	fmt.Println(result)
}
