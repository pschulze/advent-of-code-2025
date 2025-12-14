package main

import (
	"slices"
	"strconv"
	"testing"
)

func TestProcessInput(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	want := []Range{
		{Min: 11, Max: 22},
		{Min: 95, Max: 115},
		{Min: 998, Max: 1012},
		{Min: 1188511880, Max: 1188511890},
		{Min: 222220, Max: 222224},
		{Min: 1698522, Max: 1698528},
		{Min: 446443, Max: 446449},
		{Min: 38593856, Max: 38593862},
		{Min: 565653, Max: 565659},
		{Min: 824824821, Max: 824824827},
		{Min: 2121212118, Max: 2121212124},
	}

	got := processInput(input)

	// TODO: Implement equality check for slices of Range
	if len(got) != len(want) {
		t.Errorf("processInput(%q) returned %d ranges; want %d ranges", input, len(got), len(want))
		return
	}

	for i := range got {
		if got[i].Min != want[i].Min || got[i].Max != want[i].Max {
			t.Errorf("processInput(%q)[%d] = %v; want %v", input, i, got[i], want[i])
		}
	}
}

func TestValidId(t *testing.T) {
	type TestCase struct {
		id   int
		want bool
	}

	tests := map[string]TestCase{
		"odd digits":                           {id: 123, want: true},
		"even digits, sequence doesn't repeat": {id: 12, want: true},
		"even digits, sequence repeats":        {id: 1212, want: false},
		"single digit":                         {id: 1, want: true},
		"long valid id":                        {id: 123456789, want: true},
		"long invalid id":                      {id: 1188511885, want: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := validId(tc.id)
			if got != tc.want {
				t.Errorf("validId(%d) = %v; want %v", tc.id, got, tc.want)
			}
		})
	}
}

func TestValidId2(t *testing.T) {
	type TestCase struct {
		id   int
		want bool
	}

	tests := []TestCase{
		{id: 1, want: true},
		{id: 11, want: false},
		{id: 12341234, want: false},
		{id: 123123123, want: false},
		{id: 1212121212, want: false},
		{id: 1111111, want: false},
		{id: 123456, want: true},
	}

	for _, tc := range tests {
		t.Run(strconv.Itoa(tc.id), func(t *testing.T) {
			got := validId2(tc.id)
			if got != tc.want {
				t.Errorf("validId2(%d) = %v; want %v", tc.id, got, tc.want)
			}
		})
	}
}

func TestInvalidIds(t *testing.T) {
	type TestCase struct {
		rng  Range
		want []int
	}

	tests := map[string]TestCase{
		"11-22":                 {rng: Range{Min: 11, Max: 22}, want: []int{11, 22}},
		"95-115":                {rng: Range{Min: 95, Max: 115}, want: []int{99}},
		"998-1012":              {rng: Range{Min: 998, Max: 1012}, want: []int{1010}},
		"1188511880-1188511890": {rng: Range{Min: 1188511880, Max: 1188511890}, want: []int{1188511885}},
		"222220-222224":         {rng: Range{Min: 222220, Max: 222224}, want: []int{222222}},
		"1698522-1698528":       {rng: Range{Min: 1698522, Max: 1698528}, want: []int{}},
		"446443-446449":         {rng: Range{Min: 446443, Max: 446449}, want: []int{446446}},
		"38593856-38593862":     {rng: Range{Min: 38593856, Max: 38593862}, want: []int{38593859}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := invalidIds(tc.rng)
			if !slices.Equal(got, tc.want) {
				t.Errorf("invalidIds(%d, %d) = %v; want %v", tc.rng.Min, tc.rng.Max, got, tc.want)
			}
		})
	}
}
