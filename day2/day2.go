package main

import (
    "fmt"
    "io/ioutil"
    "flag"
    "os"
    "strings"
    "regexp"
    "strconv"
)

type PasswordPolicy struct {
    min int
    max int
    char rune
    password string
}

func (p PasswordPolicy) isValid() bool {
    occurs := CountOccurrences(p.password, p.char)
    return p.min <= occurs && p.max >= occurs
}

func CountOccurrences(str string, char rune) (count int) {
    for _, c := range str {
        if c == char {
            count += 1
        }
    }
    return
}

func ValidatePasswords(passwords []PasswordPolicy) (count int) {
    for _, password := range passwords {
        if password.isValid() {
            count += 1
        }
    }
    return
}

func ParsePasswords(lines []string) (passwords []PasswordPolicy) {
    re := regexp.MustCompile(`(\d+)-(\d+)\s(\w):\s(\w+)`)
    for _, line := range lines {
        if line == "" {
            break
        }
        m := re.FindStringSubmatch(line)
        min, _ := strconv.Atoi(m[1])
        max, _ := strconv.Atoi(m[2])
        passwords = append(passwords, PasswordPolicy{
            min: min,
            max: max,
            char: []rune(m[3])[0],
            password: m[4],
        })
    }
    return passwords
}

func die(err error) {
    fmt.Println(err)
    os.Exit(1)
}

func main() {
    flag.Parse()
    filename := flag.Arg(0)
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        die(err)
    }

    lines := strings.Split(string(content), "\n")
    passwords := ParsePasswords(lines)

    result := ValidatePasswords(passwords)
    fmt.Println(result)
}
