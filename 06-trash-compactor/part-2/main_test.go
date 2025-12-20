package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestSpaceIndexes(t *testing.T) {
	input := "  6 98  215 314"
	want := []int{0, 1, 3, 6, 7, 11}
	got := spaceIndexes(input)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestIntersection(t *testing.T) {
	type testCase struct {
		a, b []int
		want []int
	}

	tests := map[string]testCase{
		"some common":      {a: []int{1, 3, 5, 7}, b: []int{3, 4, 5, 6}, want: []int{3, 5}},
		"no common":        {a: []int{1, 2, 3}, b: []int{4, 5, 6}, want: []int{}},
		"all common":       {a: []int{1, 2, 3}, b: []int{1, 2, 3}, want: []int{1, 2, 3}},
		"empty a":          {a: []int{}, b: []int{1, 2, 3}, want: []int{}},
		"empty b":          {a: []int{1, 2, 3}, b: []int{}, want: []int{}},
		"both empty":       {a: []int{}, b: []int{}, want: []int{}},
		"a contains all b": {a: []int{1, 2, 3, 4, 5}, b: []int{2, 3}, want: []int{2, 3}},
		"b contains all a": {a: []int{2, 3}, b: []int{1, 2, 3, 4, 5}, want: []int{2, 3}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := intersection(tc.a, tc.b)
			if !slices.Equal(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSplitAtIndexes(t *testing.T) {
	type testCase struct {
		line string
		idxs []int
		want []string
	}

	idxs := []int{3, 7, 11}
	tests := []testCase{
		{line: "123 328  51 64 ", idxs: idxs, want: []string{"123", "328", " 51", "64 "}},
		{line: " 45 64  387 23 ", idxs: idxs, want: []string{" 45", "64 ", "387", "23 "}},
		{line: "  6 98  215 314", idxs: idxs, want: []string{"  6", "98 ", "215", "314"}},
	}

	for _, tc := range tests {
		t.Run(tc.line, func(t *testing.T) {
			got := splitAtIndexes(tc.line, tc.idxs)
			if !slices.Equal(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestTransformValues(t *testing.T) {
	type testCase struct {
		values []string
		want   []int
	}

	tests := []testCase{
		{values: []string{"64 ", "23 ", "314"}, want: []int{623, 431, 4}},
		{values: []string{" 51", "387", "215"}, want: []int{32, 581, 175}},
		{values: []string{"328", "64 ", "98 "}, want: []int{369, 248, 8}},
		{values: []string{"123", " 45", "  6"}, want: []int{1, 24, 356}},
		{values: []string{"  5", "535", "717", "995"}, want: []int{579, 319, 5575}},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.values), func(t *testing.T) {
			got := transformValue(tc.values)
			if !slices.Equal(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
