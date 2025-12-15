package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Range struct {
	Min int
	Max int
}

func (r Range) Equal(o Range) bool {
	return r.Min == o.Min && r.Max == o.Max
}

func (r Range) Combine(o Range) (Range, bool) {
	success := false

	// New values on the left side. E.g. r = 4-10, o = 1-5 -> r = 1-10
	if o.Min < r.Min && o.Max >= r.Min && o.Max <= r.Max {
		r.Min = o.Min
		success = true
	}

	// New values on the right side. E.g. r = 1-10, o = 6-15 -> r = 1-15
	if o.Max > r.Max && o.Min >= r.Min && o.Min <= r.Max {
		r.Max = o.Max
		success = true
	}

	// Handle o.Max adjacent to r.Min. E.g. r = 6-10, o = 1-5 -> r = 1-10
	if o.Max+1 == r.Min {
		r.Min = o.Min
		success = true
	}

	// Handle o.Min adjacent to r.Max. E.g. r = 1-5, o = 6-10 -> r = 1-10
	if o.Min-1 == r.Max {
		r.Max = o.Max
		success = true
	}

	// o contains r completely. E.g. r = 4-10, o = 1-15 -> r = 1-15
	if o.Min <= r.Min && o.Max >= r.Max {
		r.Min = o.Min
		r.Max = o.Max
		success = true
	}

	// r contains o completely. E.g. r = 1-15, o = 4-10 -> r = 1-15
	if r.Min <= o.Min && r.Max >= o.Max {
		success = true
	}

	return r, success
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

	idRanges := make([]Range, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		idRange, err := parseIngredientIdRange(line)
		if err != nil {
			panic(err)
		}

		idRanges = append(idRanges, idRange)
	}

	freshCount := 0
	for scanner.Scan() {
		id, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		for _, idRange := range idRanges {
			if id >= idRange.Min && id <= idRange.Max {
				freshCount++
				break
			}
		}
	}

	combinedRanges := reduceRanges(idRanges)
	for !reflect.DeepEqual(combinedRanges, idRanges) {
		idRanges = combinedRanges
		combinedRanges = reduceRanges(idRanges)
	}

	numIngredients := 0
	for _, r := range combinedRanges {
		numIngredients += (r.Max - r.Min + 1)
	}

	fmt.Printf("Number of fresh ingredients in stock: %d\n", freshCount)
	fmt.Printf("Number of fresh ingredients in total: %d\n", numIngredients)
}

// Assumes an input of a string in the format "minId-maxId"
// Where min and max are integers.
func parseIngredientIdRange(input string) (Range, error) {
	ids := strings.Split(input, "-")
	if len(ids) != 2 {
		return Range{}, errors.New("invalid input format, expecting \"minId-maxId\"")
	}

	min, err := strconv.Atoi(ids[0])
	if err != nil {
		return Range{}, err
	}

	max, err := strconv.Atoi(ids[1])
	if err != nil {
		return Range{}, err
	}

	return Range{Min: min, Max: max}, nil
}

func reduceRanges(ranges []Range) []Range {
	anyCombined := false
	combined := make([]Range, 0)

	for _, r := range ranges {
		wasCombined := false
		for i, c := range combined {
			var combinedRange Range

			combinedRange, wasCombined = r.Combine(c)
			if wasCombined {
				anyCombined = true
				combined[i] = combinedRange
				break
			}
		}

		if !wasCombined {
			combined = append(combined, r)
		}
	}

	if anyCombined {
		return reduceRanges(combined)
	}

	return combined
}
