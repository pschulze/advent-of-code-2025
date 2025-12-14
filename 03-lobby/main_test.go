package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestParseJoltage(t *testing.T) {
	type testCase struct {
		input    string
		want     []int
		hasError bool
	}

	tests := []testCase{
		{input: "987654321111111", want: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}, hasError: false},
		{input: "811111111111119", want: []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}, hasError: false},
		{input: "234234234234278", want: []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}, hasError: false},
		{input: "818181911112111", want: []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}, hasError: false},
		{input: "foobar", want: nil, hasError: true},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got, err := parseJoltage(tc.input)
			if (err != nil) != tc.hasError {
				t.Errorf("parseJoltage(%q) error = %v; want error: %v", tc.input, err != nil, tc.hasError)
			}

			if !slices.Equal(got, tc.want) {
				t.Errorf("parseJoltage(%q) = %v; want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestMaxJoltage(t *testing.T) {
	type testCase struct {
		joltages []int
		want     int
	}

	tests := []testCase{
		{joltages: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}, want: 98},
		{joltages: []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}, want: 89},
		{joltages: []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}, want: 78},
		{joltages: []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}, want: 92},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.joltages), func(t *testing.T) {
			got := maxJoltage(tc.joltages)
			if got != tc.want {
				t.Errorf("maxJoltage(%v) = %d; want %d", tc.joltages, got, tc.want)
			}
		})
	}
}

func TestMaxJoltageArbitrary(t *testing.T) {
	type testCase struct {
		joltages []int
		n        int
		want     int
	}

	tests := []testCase{
		{joltages: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}, n: 12, want: 987654321111},
		{joltages: []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}, n: 12, want: 811111111119},
		{joltages: []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}, n: 12, want: 434234234278},
		{joltages: []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}, n: 12, want: 888911112111},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.joltages), func(t *testing.T) {
			got := maxJoltageArbitrary(tc.joltages, tc.n)
			if got != tc.want {
				t.Errorf("maxJoltageArbitrary(%v, %d) = %d; want %d", tc.joltages, tc.n, got, tc.want)
			}
		})
	}
}
