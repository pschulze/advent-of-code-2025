package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide a part number (1 or 2) and an input filename as arguments.")
		return
	}

	part := os.Args[1]
	filename := os.Args[2]

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

		switch part {
		case "1":
			totalJoltage += maxJoltage(joltages)
		case "2":
			totalJoltage += maxJoltageArbitrary(joltages, 12)
		default:
			fmt.Println("Invalid part number. Please provide 1 or 2.")
			return
		}
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

		if joltage > secondDigit && !firstChanged {
			secondDigit = joltage
		}
	}

	return firstDigit*10 + secondDigit
}

func maxJoltageArbitrary(joltages []int, n int) int {
	digits := make([]int, n)

	for i, joltage := range joltages {
		for j := range n {
			if joltage > digits[j] && len(joltages)-i >= n-j {
				digits[j] = joltage
				for k := j + 1; k < n; k++ {
					digits[k] = 0
				}
				break
			}
		}
	}

	sum := 0
	for _, digit := range digits {
		sum = sum*10 + digit
	}

	return sum
}
