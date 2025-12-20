package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	AddOp  = "+"
	MultOp = "*"
)

type problem struct {
	values   []string
	operator string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an input filename as an argument.")
		return
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	idxs := make([]int, 0)
	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		if strings.ContainsAny(line, "+*") {
			break
		}

		if len(idxs) == 0 {
			idxs = spaceIndexes(line)
		} else {
			idxs = intersection(idxs, spaceIndexes(line))
		}
	}

	problems := make([]problem, 0)
	for _, line := range lines {
		if strings.ContainsAny(line, "+*") {
			operators := strings.Fields(line)
			for i := range problems {
				problems[i].operator = operators[i]
			}

			break
		}

		parts := splitAtIndexes(line, idxs)
		if len(problems) == 0 {
			problems = make([]problem, len(parts))
		}

		for i, part := range parts {
			problems[i].values = append(problems[i].values, part)
		}
	}

	total := 0
	for _, p := range problems {
		values := transformValue(p.values)

		switch p.operator {
		case AddOp:
			result := 0
			for _, v := range values {
				result += v
			}
			total += result

		case MultOp:
			result := 1
			for _, v := range values {
				result *= v
			}
			total += result
		}
	}

	fmt.Println("Total of all problem results:", total)
}

// Returns the indexes of all spaces in a string.
// Sorted in ascending order.
func spaceIndexes(line string) []int {
	idxs := make([]int, 0)
	for i := 0; i < len(line); i++ {
		if line[i] == ' ' {
			idxs = append(idxs, i)
		}
	}

	return idxs
}

func intersection(a, b []int) []int {
	result := make([]int, 0)

	for i, j := 0, 0; i < len(a) && j < len(b); {
		if a[i] == b[j] {
			result = append(result, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}

	return result
}

func splitAtIndexes(line string, idxs []int) []string {
	parts := make([]string, 0)
	startIdx := 0
	for i := range idxs {
		endIdx := idxs[i]
		parts = append(parts, line[startIdx:endIdx])
		startIdx = endIdx + 1
	}

	parts = append(parts, line[startIdx:])
	return parts
}

func transformValue(values []string) []int {
	result := make([]int, len(values[0]))
	for _, vStr := range values {
		for c := 0; c < len(vStr); c++ {
			if vStr[c] == ' ' {
				continue
			}

			v := int(vStr[c] - '0')
			result[c] = result[c]*10 + v
		}
	}

	return result
}
