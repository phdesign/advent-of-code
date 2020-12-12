package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var validEyeColours = [...]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

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

func indexOf(items []string, value string) int {
	for i, item := range items {
		if item == value {
			return i
		}
	}
	return -1
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
	heightPattern := regexp.MustCompile(`^(\d+)(cm|in)$`)
	matches := heightPattern.FindStringSubmatch(p.height)
	if len(matches) == 0 {
		return false
	}
	height, _ := strconv.Atoi(matches[1])
	switch matches[2] {
	case "cm":
		if height < 150 || height > 193 {
			return false
		}
	case "in":
		if height < 59 || height > 76 {
			return false
		}
	default:
		return false
	}
	hairColorPattern := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	if !hairColorPattern.MatchString(p.hairColor) {
		return false
	}
	if indexOf(validEyeColours[:], p.eyeColor) == -1 {
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
	result := ValidatePassports(passports, FieldValueValidator)
	fmt.Println(result)
}
