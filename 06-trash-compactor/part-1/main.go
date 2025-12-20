package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	AddOp  = "+"
	MultOp = "*"
)

type problem struct {
	values   []int
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

	problems := make([]problem, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if fields[0] == AddOp || fields[0] == MultOp {
			operators := parseOperators(fields)
			for i := range problems {
				problems[i].operator = operators[i]
			}

			continue
		}

		if len(problems) == 0 {
			for range len(fields) {
				problems = append(problems, problem{values: make([]int, 0)})
			}
		}

		values := parseValues(fields)
		for i := range problems {
			problems[i].values = append(problems[i].values, values[i])
		}

	}

	total := 0
	for _, p := range problems {
		result := p.solve()
		fmt.Printf("Problem: %v %s = %d\n", p.values, p.operator, result)
		total += result
	}

	fmt.Printf("Total of all problem results: %d\n", total)
}

func parseFields(line string) []string {
	return strings.Fields(line)
}

func parseValues(fields []string) []int {
	values := make([]int, 0, len(fields))
	for _, field := range fields {
		v, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}

		values = append(values, v)
	}

	return values
}

func parseOperators(fields []string) []string {
	operators := make([]string, 0, len(fields))

	for _, field := range fields {
		if field != AddOp && field != MultOp {
			panic("unknown operator: " + field)
		}

		operators = append(operators, field)
	}

	return operators
}

func (p *problem) solve() int {
	switch p.operator {
	case "+":
		return p.sum()
	case "*":
		return p.product()
	default:
		panic("unknown operator: " + p.operator)
	}
}

func (p *problem) sum() int {
	total := 0
	for _, v := range p.values {
		total += v
	}

	return total
}

func (p *problem) product() int {
	total := 1
	for _, v := range p.values {
		total *= v
	}

	return total
}
