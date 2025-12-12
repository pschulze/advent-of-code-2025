package main

import (
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
		got, err := parseJoltage(tc.input)
		if (err != nil) != tc.hasError {
			t.Errorf("parseJoltage(%q) error = %v; want error: %v", tc.input, err != nil, tc.hasError)
			continue
		}

		if !slices.Equal(got, tc.want) {
			t.Errorf("parseJoltage(%q) = %v; want %v", tc.input, got, tc.want)
		}
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
		got := maxJoltage(tc.joltages)
		if got != tc.want {
			t.Errorf("maxJoltage(%v) = %d; want %d", tc.joltages, got, tc.want)
		}
	}
}
