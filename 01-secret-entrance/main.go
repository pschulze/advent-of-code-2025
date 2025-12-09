package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

	pos := 50
	stopCount := 0
	passCount := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		instruction := scanner.Text()
		steps, err := decode(instruction)
		if err != nil {
			panic(err)
		}

		passCount += zeroPasses(pos, steps)
		pos = rotate(pos, steps)
		if pos == 0 {
			stopCount += 1
		}
	}

	println("Final Position:", pos)
	println("Number of times at position 0:", stopCount)
	println("Number of times position 0 passed: ", passCount)
}

func decode(instruction string) (int, error) {
	regex := regexp.MustCompile(`^(L|R)(\d+)$`)
	matches := regex.FindStringSubmatch(instruction)
	if len(matches) != 3 {
		return 0, errors.New("Invalid instruction format: " + instruction)
	}

	direction := matches[1]
	steps, error := strconv.Atoi(matches[2])
	if error != nil {
		return 0, error
	}

	if direction == "L" {
		steps = -steps
	}

	return steps, nil
}

func rotate(pos int, steps int) int {
	intermediate := pos + steps
	for intermediate < 0 {
		intermediate = 100 + intermediate
	}

	return intermediate % 100
}
func zeroPasses(pos int, steps int) int {
	count := 0

	absSteps := steps
	if absSteps < 0 {
		absSteps = -absSteps
	}

	// Check if partial revolution crosses zero
	remainder := absSteps % 100

	// Need extra check to avoid counting if already at 0
	if steps > 0 && pos+remainder >= 100 {
		count += 1
	} else if steps < 0 && pos-remainder <= 0 && pos != 0 {
		count += 1
	}

	// Add full revolutions
	count += absSteps / 100

	return count
}
