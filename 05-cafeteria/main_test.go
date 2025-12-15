package main

import (
	"reflect"
	"testing"
)

func TestParseIngredientIdRange(t *testing.T) {
	type testCase struct {
		input   string
		want    Range
		wantErr bool
	}

	tests := []testCase{
		{input: "1-3", want: Range{Min: 1, Max: 3}, wantErr: false},
		{input: "foobar", want: Range{}, wantErr: true},
		{input: "5-5", want: Range{Min: 5, Max: 5}, wantErr: false},
		{input: "a-3", want: Range{}, wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got, err := parseIngredientIdRange(tc.input)
			if (err != nil) && !tc.wantErr {
				t.Errorf("got error %v, want no error", err)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestReduceRanges(t *testing.T) {
	type testCase struct {
		input []Range
		want  []Range
	}

	tests := map[string]testCase{
		"no overlap": {
			input: []Range{{Min: 1, Max: 3}, {Min: 5, Max: 7}},
			want:  []Range{{Min: 1, Max: 3}, {Min: 5, Max: 7}},
		},
		"simple overlap": {
			input: []Range{{Min: 1, Max: 5}, {Min: 4, Max: 7}},
			want:  []Range{{Min: 1, Max: 7}},
		},
		"total containment": {
			input: []Range{{Min: 1, Max: 10}, {Min: 3, Max: 7}},
			want:  []Range{{Min: 1, Max: 10}},
		},
		"total containment reversed": {
			input: []Range{{Min: 3, Max: 7}, {Min: 1, Max: 10}},
			want:  []Range{{Min: 1, Max: 10}},
		},
		"multiple overlaps": {
			input: []Range{{Min: 1, Max: 3}, {Min: 2, Max: 5}, {Min: 4, Max: 6}},
			want:  []Range{{Min: 1, Max: 6}},
		},
		"touching ranges": {
			input: []Range{{Min: 1, Max: 3}, {Min: 3, Max: 5}},
			want:  []Range{{Min: 1, Max: 5}},
		},
		"multiple distinct overlaps": {
			input: []Range{{Min: 1, Max: 3}, {Min: 2, Max: 4}, {Min: 6, Max: 10}, {Min: 7, Max: 12}, {Min: 11, Max: 15}},
			want:  []Range{{Min: 1, Max: 4}, {Min: 6, Max: 15}},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := reduceRanges(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCombine(t *testing.T) {
	type testCase struct {
		r1          Range
		r2          Range
		wantRange   Range
		wantSuccess bool
	}

	tests := map[string]testCase{
		"overlapping": {
			r1:          Range{Min: 1, Max: 5},
			r2:          Range{Min: 4, Max: 10},
			wantRange:   Range{Min: 1, Max: 10},
			wantSuccess: true,
		},
		"non-overlapping": {
			r1:          Range{Min: 1, Max: 3},
			r2:          Range{Min: 5, Max: 7},
			wantRange:   Range{Min: 1, Max: 3},
			wantSuccess: false,
		},
		"contained": {
			r1:          Range{Min: 1, Max: 10},
			r2:          Range{Min: 3, Max: 7},
			wantRange:   Range{Min: 1, Max: 10},
			wantSuccess: true,
		},
		"touching min": {
			r1:          Range{Min: 6, Max: 10},
			r2:          Range{Min: 1, Max: 5},
			wantRange:   Range{Min: 1, Max: 10},
			wantSuccess: true,
		},
		"touching max": {
			r1:          Range{Min: 1, Max: 5},
			r2:          Range{Min: 6, Max: 10},
			wantRange:   Range{Min: 1, Max: 10},
			wantSuccess: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotRange, gotSuccess := tc.r1.Combine(tc.r2)

			if !reflect.DeepEqual(gotRange, tc.wantRange) {
				t.Errorf("got %v, want %v", gotRange, tc.wantRange)
			}

			if gotSuccess != tc.wantSuccess {
				t.Errorf("got success %v, want %v", gotSuccess, tc.wantSuccess)
			}
		})
	}
}

func TestDeepEquality(t *testing.T) {
	ranges1 := []Range{{Min: 1, Max: 5}, {Min: 6, Max: 10}}
	ranges2 := []Range{{Min: 1, Max: 5}, {Min: 6, Max: 10}}
	ranges3 := []Range{{Min: 1, Max: 4}, {Min: 6, Max: 10}}

	if !reflect.DeepEqual(ranges1, ranges2) {
		t.Errorf("expected ranges1 and ranges2 to be equal")
	}

	if reflect.DeepEqual(ranges1, ranges3) {
		t.Errorf("expected ranges1 and ranges3 to be not equal")
	}
}
