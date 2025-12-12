package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an input file.")
		return
	}

	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	totalJoltage := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := scanner.Text()
		joltages, err := parseJoltage(input)
		if err != nil {
			panic(err)
		}

		totalJoltage += maxJoltage(joltages)
	}

	println("Total output joltage:", totalJoltage)
}

func parseJoltage(input string) ([]int, error) {
	joltages := make([]int, len(input))

	for i, val := range strings.Split(input, "") {
		joltage, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}

		joltages[i] = joltage
	}

	return joltages, nil
}

func maxJoltage(joltages []int) int {
	// Find biggest digit that isn't in the last place
	// If multiple digits qualify, choose the leftmost one
	// Find the biggest digit that comes after that one
	firstDigit := 0
	secondDigit := 0

	for i, joltage := range joltages {
		firstChanged := false

		if i != len(joltages)-1 {
			if joltage > firstDigit {
				firstChanged = true
				firstDigit = joltage
				secondDigit = 0
			}
		}

		if i != 0 {
			if joltage > secondDigit && !firstChanged {
				secondDigit = joltage
			}
		}
	}

	return firstDigit*10 + secondDigit
}
