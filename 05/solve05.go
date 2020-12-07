package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("input05.txt")
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
	result := 0
	for _, line := range input {
		if seatID := getSeatID(line); seatID > result {
			result = seatID
		}
	}
	return result
}

// solveB solves the second half
func solveB(input []string) int {
	seats := sort.IntSlice{}
	for _, line := range input {
		seats = append(seats, getSeatID(line))
	}
	seats.Sort()
	for i, v := range seats {
		if i > 0 && i < len(seats)-1 && (v+1) != seats[i+1] {
			return v + 1
		}
	}
	panic("all seats accounted for")
}

func getSeatID(line string) int {
	row := binaryCount(line[:7], "F", "B")
	col := binaryCount(line[7:], "L", "R")
	return row*8 + col
}

func binaryCount(line string, lowerGlyph string, upperGlyph string) int {
	sum := 0
	stride := 1
	for i := len(line) - 1; i >= 0; i-- {
		c := string(line[i])
		if !(c == lowerGlyph || c == upperGlyph) {
			fmt.Printf("bad string")
			return 0
		}
		if c == upperGlyph {
			sum += stride
		}
		stride *= 2
	}
	return sum
}
