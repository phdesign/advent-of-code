package main

import (
    "reflect"
    "testing"
)

func TestParseFields(t *testing.T) {
    actual := ParseFields("ecl:gry pid:860033327")
    expected := map[string]string{
        "ecl": "gry",
        "pid": "860033327",
    }

    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("Expected %v, received %v", expected, actual)
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

    actual := ParsePassports(input)
    expected := []Passport{
        Passport{birthYear: "1937", issueYear: "2017", expiryYear: "2020", height: "183cm", hairColor: "#fffffd", eyeColor: "gry", passportId: "860033327", countryId: "147"},
        Passport{birthYear: "1929", issueYear: "2013", expiryYear: "2023", height: "", hairColor: "#cfa07d", eyeColor: "amb", passportId: "028048884", countryId: "350"},
        Passport{birthYear: "1931", issueYear: "2013", expiryYear: "2024", height: "179cm", hairColor: "#ae17e1", eyeColor: "brn", passportId: "760753108", countryId: ""},
        Passport{birthYear: "", issueYear: "2011", expiryYear: "2025", height: "59in", hairColor: "#cfa07d", eyeColor: "brn", passportId: "166559648", countryId: ""},
    }

    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("Expected %v, received %v", expected, actual)
    }
}

func TestPassports(t *testing.T) {
    passport := Passport{birthYear: "1937", issueYear: "2017", expiryYear: "2020", height: "183cm", hairColor: "#fffffd", eyeColor: "gry", passportId: "860033327", countryId: "147"}
    actual := passport.isValid()
    assertTrue(t, actual)
}

func TestValidatePassports(t *testing.T) {

}

func assertTrue(t *testing.T, actual bool) {
    t.Helper()
    if !actual {
        t.Error("Expected result to be true, was false")
    }
}
