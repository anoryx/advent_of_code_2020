package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	file, err := os.Open("input03.txt")
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
	slopes := []Slope{
		Slope{3, 1},
	}
	return solver(slopes, input)
}

// solveB solves the second half
func solveB(input []string) int {
	slopes := []Slope{
		Slope{1, 1},
		Slope{3, 1},
		Slope{5, 1},
		Slope{7, 1},
		Slope{1, 2},
	}
	return solver(slopes, input)
}

// solver concurrently analyzes a run and combines results
func solver(slopes []Slope, input []string) int {
	var wg sync.WaitGroup
	c := make(chan int, len(slopes))

	for _, slope := range slopes {
		wg.Add(1)
		go countTrees(slope, input, c, &wg)
	}
	wg.Wait()
	close(c)
	result := 1
	for t := range c {
		result *= t
	}
	return result
}

// countTrees counts the trees in a run
func countTrees(slope Slope, input []string, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	trees := 0
	columnIndex := 0
	for rowIndex, rowData := range input {
		if rowIndex%slope.y == 0 {
			k := (slope.x * columnIndex) % len(rowData)
			columnIndex++
			if string(rowData[k]) == "#" {
				trees++
			}
		}
	}
	results <- trees
}

// Slope is the slope
type Slope struct {
	x int
	y int
}
