package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"../aocutil"
)

func main() {
	file, err := os.Open("input02.txt")
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

	fmt.Printf("A: %v\n", SolveA(lines))
	fmt.Printf("B: %v\n", SolveB(lines))
}

// SolveA solves the first half
func SolveA(input []string) int {
	counter := 0
	for _, line := range input {
		policy := ParseLine(line)
		if c := len(regexp.MustCompile(policy.letter).
			FindAllString(policy.password, -1)); c >= policy.min && c <= policy.max {
			counter++
		}

	}
	return counter
}

// SolveB solves the second half
func SolveB(input []string) int {
	counter := 0
	for _, line := range input {
		policy := ParseLine(line)
		if (string(policy.password[policy.min-1]) == policy.letter) != (string(policy.password[policy.max-1]) == policy.letter) {
			counter++
		}
	}
	return counter
}

// ParseLine parses lines
func ParseLine(input string) PasswordPolicy {
	re := regexp.MustCompile(`[-:\s]+`)
	match := re.Split(input, -1)
	return PasswordPolicy{
		min:      aocutil.MustAtoi(match[0]),
		max:      aocutil.MustAtoi(match[1]),
		letter:   match[2],
		password: match[3],
	}
}

// PasswordPolicy manages puzzle data
type PasswordPolicy struct {
	min      int
	max      int
	letter   string
	password string
}
