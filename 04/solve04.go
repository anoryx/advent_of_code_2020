package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"../aocutil"
)

func main() {
	file, err := os.Open("input04.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("A: %v\n", solveA(lines))
	fmt.Printf("B: %v\n", solveB(lines))
}

// solveA solves the first half
func solveA(input []string) int {
	f := func(s string) bool {
		return checkPassport(s)
	}
	return solver(input, f)
}

// solveB solves the second half
func solveB(input []string) int {
	f := func(s string) bool {
		return checkPassport(s) && validatePassport(s)
	}
	return solver(input, f)
}

// solver iterates through input and performs c
func solver(input []string, f func(string) bool) int {
	validPassports := 0
	passportLines := []string{}
	for _, line := range input {
		if line != "" {
			passportLines = append(passportLines, line)
		} else {
			passport := strings.Join(passportLines, " ")
			if f(passport) {
				validPassports++
			}
			passportLines = []string{}
		}
	}
	passport := strings.Join(passportLines, " ")
	if f(passport) {
		validPassports++
	}
	return validPassports
}

// checkPassport verifies whether the passport contains all required fields
func checkPassport(passport string) bool {
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, v := range fields {
		if !strings.Contains(passport, v) {
			return false
		}
	}
	return true
}

func validatePassport(passport string) bool {
	fields := strings.Split(passport, " ")
	for _, v := range fields {
		p := strings.Split(v, ":")
		item := PassportField{p[0], p[1]}
		switch item.key {
		case "byr":
			// 4 digits 1920 <= x <= 2002
			re := regexp.MustCompile(`^\d{4}$`)
			if !(re.MatchString(item.value) &&
				aocutil.MustAtoi(item.value) >= 1920 &&
				aocutil.MustAtoi(item.value) <= 2002) {
				return false
			}
		case "iyr":
			// 4 digits 2010 <= x <= 2020
			re := regexp.MustCompile(`^\d{4}$`)
			if !(re.MatchString(item.value) &&
				aocutil.MustAtoi(item.value) >= 2010 &&
				aocutil.MustAtoi(item.value) <= 2020) {
				return false
			}
		case "eyr":
			// 4 digits; 2020 <= x <= 2030
			re := regexp.MustCompile(`^\d{4}$`)
			if !(re.MatchString(item.value) &&
				aocutil.MustAtoi(item.value) >= 2020 &&
				aocutil.MustAtoi(item.value) <= 2030) {
				return false
			}
		case "hgt":
			// num followed by "cm" or "in"
			// "cm": 150 <= x <= 193
			// "in": 59 <= x <= 76
			re := regexp.MustCompile(`^(?P<height>\d+)(?P<unit>in|cm)$`)
			matchMap := aocutil.GetRegexpMap(re, item.value)
			if len(matchMap) == 0 {
				return false
			}
			height := aocutil.MustAtoi(matchMap["height"])
			unit := matchMap["unit"]
			validHeight := (unit == "in" && height >= 59 && height <= 76) ||
				(unit == "cm" && height >= 150 && height <= 193)
			if !validHeight {
				return false
			}
		case "hcl":
			// "#" followed by 6 [0-9a-f]
			re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
			if !re.MatchString(item.value) {
				return false
			}
		case "ecl":
			// is member of validColors
			validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			if !aocutil.ContainsString(validColors, item.value) {
				return false
			}
		case "pid":
			// 9 digit number
			re := regexp.MustCompile(`^\d{9}$`)
			if !re.MatchString(item.value) {
				return false
			}
		}
	}
	return true
}

// PassportField is a very simple helper struct
type PassportField struct {
	key   string
	value string
}
