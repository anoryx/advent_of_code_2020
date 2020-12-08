package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input06.txt")
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
	sum := 0
	characterCount := make(map[string]int)
	for _, line := range input {
		if line != "" {
			for _, char := range line {
				characterCount[string(char)]++
			}
		} else {
			sum += len(characterCount)
			characterCount = make(map[string]int)
		}
	}
	sum += len(characterCount)
	return sum
}

// solveB solves the second half
func solveB(input []string) int {
	sum := 0
	characterCount := make(map[string]int)
	lineCount := 0
	for _, line := range input {
		if line != "" {
			for _, char := range line {
				characterCount[string(char)]++
			}
			lineCount++
		} else {
			sum += mapCounter(characterCount, lineCount)
			lineCount = 0
			characterCount = make(map[string]int)
		}
	}
	sum += mapCounter(characterCount, lineCount)
	return sum
}

// mapCounter returns the number of keys with value == n
func mapCounter(m map[string]int, n int) int {
	count := 0
	for _, v := range m {
		if v == n {
			count++
		}
	}
	return count
}
