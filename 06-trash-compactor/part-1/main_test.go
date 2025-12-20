package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestParseFields(t *testing.T) {
	type testCase struct {
		input string
		want  []string
	}

	tests := []testCase{
		{input: "123 328  51 64 ", want: []string{"123", "328", "51", "64"}},
		{input: " 45 64  387 23 ", want: []string{"45", "64", "387", "23"}},
		{input: "  6 98  215 314", want: []string{"6", "98", "215", "314"}},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := parseFields(tc.input)
			if !slices.Equal(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestParseValues(t *testing.T) {
	type testCase struct {
		input []string
		want  []int
	}

	tests := []testCase{
		{input: []string{"123", "328", "51", "64"}, want: []int{123, 328, 51, 64}},
		{input: []string{"45", "64", "387", "23"}, want: []int{45, 64, 387, 23}},
		{input: []string{"6", "98", "215", "314"}, want: []int{6, 98, 215, 314}},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {
			got := parseValues(tc.input)
			if !slices.Equal(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestProblemSolve(t *testing.T) {
	type testCase struct {
		p    problem
		want int
	}

	tests := []testCase{
		{p: problem{values: []int{123, 45, 6}, operator: "*"}, want: 33210},
		{p: problem{values: []int{328, 64, 98}, operator: "+"}, want: 490},
		{p: problem{values: []int{51, 387, 215}, operator: "*"}, want: 4243455},
		{p: problem{values: []int{64, 23, 314}, operator: "+"}, want: 401},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v%v", tc.p.values, tc.p.operator), func(t *testing.T) {
			got := tc.p.solve()
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
