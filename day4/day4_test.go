package main

import (
	"reflect"
	"testing"
	"fmt"
)

func TestParseFields(t *testing.T) {
	got := ParseFields("ecl:gry pid:860033327")
	want := map[string]string{
		"ecl": "gry",
		"pid": "860033327",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, received %v", want, got)
	}
}

func TestParsePassports(t *testing.T) {
	input := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

	got := ParsePassports(input)
	want := []Passport{
		{birthYear: "1937", issueYear: "2017", expiryYear: "2020", height: "183cm", hairColor: "#fffffd", eyeColor: "gry", passportId: "860033327", countryId: "147"},
		{birthYear: "1929", issueYear: "2013", expiryYear: "2023", height: "", hairColor: "#cfa07d", eyeColor: "amb", passportId: "028048884", countryId: "350"},
		{birthYear: "1931", issueYear: "2013", expiryYear: "2024", height: "179cm", hairColor: "#ae17e1", eyeColor: "brn", passportId: "760753108", countryId: ""},
		{birthYear: "", issueYear: "2011", expiryYear: "2025", height: "59in", hairColor: "#cfa07d", eyeColor: "brn", passportId: "166559648", countryId: ""},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, received %v", want, got)
	}
}

func TestPassports(t *testing.T) {
	t.Run("should return true when all required fields are set given required fields validator", func(t *testing.T) {
		passports := []Passport {
			Passport{birthYear: "1937", issueYear: "2017", expiryYear: "2020", height: "183cm", hairColor: "#fffffd", eyeColor: "gry", passportId: "860033327", countryId: "147"},
		}
		for _, passport := range passports {
			t.Run("", func(t *testing.T) {
				got := passport.isValid(RequiredFieldValidator)
				assertTrue(t, got)
			})
		}
	})

	t.Run("should return false when required fields are missing given required fields validator", func(t *testing.T) {
		passports := []Passport {
			Passport{birthYear: "1929", issueYear: "2013", expiryYear: "2023", height: "", hairColor: "#cfa07d", eyeColor: "amb", passportId: "028048884", countryId: "350"},
		}
		for _, passport := range passports {
			t.Run("", func(t *testing.T) {
				got := passport.isValid(RequiredFieldValidator)
				assertFalse(t, got)
			})
		}
	})

	t.Run("should return true when all fields are valid given field values validator", func(t *testing.T) {
		passports := []Passport {
			Passport{birthYear: "2002", issueYear: "2017", expiryYear: "2020", height: "183cm", hairColor: "#fffffd", eyeColor: "gry", passportId: "860033327", countryId: "147"},
			Passport{birthYear: "2002", issueYear: "2017", expiryYear: "2020", height: "183cm", hairColor: "#fffffd", eyeColor: "gry", passportId: "060033327", countryId: "147"},
		}
		for _, passport := range passports {
			t.Run("", func(t *testing.T) {
				got := passport.isValid(FieldValueValidator)
				assertTrue(t, got)
			})
		}
	})

	t.Run("should return false when any fields are invalid given field values validator", func(t *testing.T) {
		passports := []Passport {
			Passport{birthYear: "2003", issueYear: "2017", expiryYear: "2020", height: "183cm", hairColor: "#fffffd", eyeColor: "gry", passportId: "860033327", countryId: "147"},
			Passport{birthYear: "2002", issueYear: "2017", expiryYear: "2020", height: "183cm", hairColor: "#fffffd", eyeColor: "gry", passportId: "06003332", countryId: "147"},
			Passport{birthYear: "2002", issueYear: "2017", expiryYear: "2020", height: "183cm", hairColor: "#fffffd", eyeColor: "gry", passportId: "86003332a", countryId: "147"},
			Passport{birthYear: "2002", issueYear: "2017", expiryYear: "2020", height: "183cm", hairColor: "#fffffd", eyeColor: "purple", passportId: "860033327", countryId: "147"},
			Passport{birthYear: "2002", issueYear: "2017", expiryYear: "2020", height: "183cm", hairColor: "#xffffd", eyeColor: "gry", passportId: "860033327", countryId: "147"},
			Passport{birthYear: "2002", issueYear: "2017", expiryYear: "2020", height: "183in", hairColor: "#fffffd", eyeColor: "gry", passportId: "860033327", countryId: "147"},
			Passport{birthYear: "2002", issueYear: "2017", expiryYear: "2020", height: "59cm", hairColor: "#fffffd", eyeColor: "gry", passportId: "860033327", countryId: "147"},
			Passport{birthYear: "2002", issueYear: "2017", expiryYear: "2020", height: "59", hairColor: "#fffffd", eyeColor: "gry", passportId: "860033327", countryId: "147"},
			Passport{birthYear: "2002", issueYear: "2017", expiryYear: "2020", height: "59ri", hairColor: "#fffffd", eyeColor: "gry", passportId: "860033327", countryId: "147"},
		}
		for _, passport := range passports {
			t.Run("", func(t *testing.T) {
				got := passport.isValid(FieldValueValidator)
				assertFalse(t, got)
			})
		}
	})
}

func TestValidatePassports(t *testing.T) {
	passports := []Passport{
		{birthYear: "1937", issueYear: "2017", expiryYear: "2020", height: "183cm", hairColor: "#fffffd", eyeColor: "gry", passportId: "860033327", countryId: "147"},
		{birthYear: "1929", issueYear: "2013", expiryYear: "2023", height: "", hairColor: "#cfa07d", eyeColor: "amb", passportId: "028048884", countryId: "350"},
		{birthYear: "1931", issueYear: "2013", expiryYear: "2024", height: "179cm", hairColor: "#ae17e1", eyeColor: "brn", passportId: "760753108", countryId: ""},
		{birthYear: "", issueYear: "2011", expiryYear: "2025", height: "59in", hairColor: "#cfa07d", eyeColor: "brn", passportId: "166559648", countryId: ""},
	}
	want := 2
	got := ValidatePassports(passports, RequiredFieldValidator)

	if got != want {
		t.Errorf("Expected %d, received %d", want, got)
	}
}

func TestValidPassports(t *testing.T) {
	input := `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`
	passports := ParsePassports(input)

	for _, passport := range passports {
		name := fmt.Sprintf("%v", passport)
		t.Run(name, func (t *testing.T)  {
			got := ValidatePassports([]Passport{passport}, FieldValueValidator)
			assertIntEqual(t, got, 1)
		})
	}
}

func TestInvalidPassports(t *testing.T) {
	input := `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`
	passports := ParsePassports(input)

	for _, passport := range passports {
		name := fmt.Sprintf("%v", passport)
		t.Run(name, func (t *testing.T)  {
			got := ValidatePassports([]Passport{passport}, FieldValueValidator)
			assertIntEqual(t, got, 0)
		})
	}
}

func assertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Error("Expected result to be true, was false")
	}
}

func assertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Error("Expected result to be false, was true")
	}
}

func assertBool(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("Expected result to be %v, was %v", got, want)
	}
}

func assertIntEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Expected result to be %d, was %d", got, want)
	}
}
