package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Min int
	Max int
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide a part number (1 or 2) and an input filename as arguments.")
		return
	}

	part := os.Args[1]
	filename := os.Args[2]

	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	idRanges := processInput(strings.TrimSpace(string(b)))

	invalidIdSum := 0
	for _, r := range idRanges {
		var invalid []int

		switch part {
		case "1":
			invalid = invalidIds(r)
		case "2":
			invalid = invalidIds2(r)
		default:
			fmt.Println("Invalid part number. Please provide 1 or 2.")
			return
		}

		for _, id := range invalid {
			invalidIdSum += id
		}
		fmt.Printf("Invalid IDs in range %d-%d: %v\n", r.Min, r.Max, invalid)
	}

	fmt.Printf("Sum of all invalid IDs: %d\n", invalidIdSum)
}

func processInput(input string) []Range {
	ranges := []Range{}
	rawRanges := strings.SplitSeq(input, ",")
	for r := range rawRanges {
		bounds := strings.Split(r, "-")
		if len(bounds) != 2 {
			fmt.Printf("Invalid range: %s\n", r)
			continue
		}

		min, err := strconv.Atoi(bounds[0])
		if err != nil {
			fmt.Printf("Invalid min value: %s\n", bounds[0])
			continue
		}

		max, err := strconv.Atoi(bounds[1])
		if err != nil {
			fmt.Printf("Invalid max value: %s\n", bounds[1])
			continue
		}

		ranges = append(ranges, Range{Min: min, Max: max})
	}

	return ranges
}

func invalidIds(r Range) []int {
	invalidIds := []int{}
	for i := r.Min; i <= r.Max; i++ {
		if !validId(i) {
			invalidIds = append(invalidIds, i)
		}
	}
	return invalidIds
}

func validId(id int) bool {
	idStr := strconv.Itoa(id)
	idLen := len(idStr)

	// Can't be comprised only of some sequence of repeated digits if
	// there are an odd number of digits.
	if idLen%2 != 0 {
		return true
	}

	halfLen := idLen / 2

	return idStr[:halfLen] != idStr[halfLen:]
}

func invalidIds2(r Range) []int {
	invalidIds := []int{}
	for i := r.Min; i <= r.Max; i++ {
		if !validId2(i) {
			invalidIds = append(invalidIds, i)
		}
	}
	return invalidIds
}

func validId2(id int) bool {
	idStr := strconv.Itoa(id)
	substrs := []string{}

	for length := 1; length <= len(idStr)/2; length++ {
		substrs = append(substrs, idStr[:length])
	}

	for _, substr := range substrs {
		// Can't form a repeated sequence if number of digits of id
		// is not a multiple of length of substring.
		if len(idStr)%len(substr) != 0 {
			continue
		}

		repeated := strings.Repeat(substr, len(idStr)/len(substr))
		if repeated == idStr {
			return false
		}
	}

	return true
}
