package main

import "testing"

func TestDecode(t *testing.T) {
	type testCase struct {
		instruction string
		want        int
		wantErr     bool
	}

	tests := map[string]testCase{
		"simple forward":  {instruction: "R1", want: 1, wantErr: false},
		"simple backward": {instruction: "L343", want: -343, wantErr: false},
		"zero forward":    {instruction: "R0", want: 0, wantErr: false},
		"zero backward":   {instruction: "L0", want: 0, wantErr: false},
		"invalid format":  {instruction: "X10", want: 0, wantErr: true},
		"missing steps":   {instruction: "R", want: 0, wantErr: true},
		"non-numeric":     {instruction: "Labc", want: 0, wantErr: true},
		"negative steps":  {instruction: "R-5", want: 0, wantErr: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := decode(tc.instruction)

			if (err != nil) && !tc.wantErr {
				t.Fatalf("decode(%q) unexpected error: %v", tc.instruction, err)
			}

			if got != tc.want {
				t.Fatalf("decode(%q) = %d; want %d", tc.instruction, got, tc.want)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	type testCase struct {
		pos   int
		steps int
		want  int
	}

	tests := map[string]testCase{
		"rotate forward within bounds":  {pos: 10, steps: 15, want: 25},
		"rotate forward with wrap":      {pos: 90, steps: 15, want: 5},
		"rotate backward within bounds": {pos: 50, steps: -20, want: 30},
		"rotate backward with wrap":     {pos: 10, steps: -15, want: 95},
		"rotate zero steps":             {pos: 30, steps: 0, want: 30},
		"rotate full circle":            {pos: 25, steps: 100, want: 25},
		"rotate negative full circle":   {pos: 75, steps: -100, want: 75},
		"rotate large positive steps":   {pos: 20, steps: 250, want: 70},
		"rotate large negative steps":   {pos: 80, steps: -250, want: 30},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := rotate(tc.pos, tc.steps)
			if got != tc.want {
				t.Fatalf("rotate(%d, %d) = %d; want %d", tc.pos, tc.steps, got, tc.want)
			}
		})
	}
}

func TestZeroPasses(t *testing.T) {
	type testCase struct {
		pos   int
		steps int
		want  int
	}

	tests := map[string]testCase{
		"forward pass":                   {pos: 90, steps: 20, want: 1},
		"backward pass":                  {pos: 10, steps: -20, want: 1},
		"no pass forward":                {pos: 30, steps: 10, want: 0},
		"no pass backward":               {pos: 70, steps: -10, want: 0},
		"exactly at zero forward":        {pos: 80, steps: 20, want: 1},
		"exactly at zero backward":       {pos: 20, steps: -20, want: 1},
		"multiple passes forward":        {pos: 95, steps: 210, want: 3},
		"multiple passes backward":       {pos: 5, steps: -210, want: 3},
		"zero steps":                     {pos: 50, steps: 0, want: 0},
		"full circle forward":            {pos: 25, steps: 1000, want: 10},
		"full circle backward":           {pos: 75, steps: -1000, want: 10},
		"multiple passes exact forward":  {pos: 90, steps: 310, want: 4},
		"multiple passes exact backward": {pos: 10, steps: -310, want: 4},
		"from zero forward":              {pos: 0, steps: 10, want: 0},
		"from zero backward":             {pos: 0, steps: -10, want: 0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := zeroPasses(tc.pos, tc.steps)
			if got != tc.want {
				t.Fatalf("zeroPasses(%d, %d) = %d; want %d", tc.pos, tc.steps, got, tc.want)
			}
		})
	}
}
