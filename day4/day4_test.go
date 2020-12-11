package main

import (
	"reflect"
	"testing"
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
