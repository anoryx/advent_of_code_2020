package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"

	"../aocutil"
)

func main() {
	file, err := os.Open("input01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nums := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums = append(nums, aocutil.MustAtoi(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(nums)

	fmt.Printf("A: %v\n", SolveA(nums))
	fmt.Printf("B: %v\n", SolveB(nums))
}

// SolveA returns sum of two numbers
func SolveA(input []int) int {
	target := 2020
	left, right := 0, len(input)-1
Loop:
	for {
		sum := input[left] + input[right]
		switch {
		case left >= right:
			break Loop
		case sum == target:
			return input[left] * input[right]
		case target > sum:
			left++
		case sum > target:
			right--
		}
	}
	log.Fatal("exhausted search, nothing found")
	return 0
}

// SolveB returns sum of three numbers
func SolveB(input []int) int {
	target := 2020
	for i := 0; i < len(input)-3; i++ {
		j, k := i+1, len(input)-1
	Loop:
		for {

			sum := input[i] + input[j] + input[k]
			switch {
			case j >= k:
				break Loop
			case sum == target:
				return input[i] * input[j] * input[k]
			case target > sum:
				j++
			case sum > target:
				k--
			}
		}
	}
	log.Fatal("exhausted search, nothing found")
	return 0
}
