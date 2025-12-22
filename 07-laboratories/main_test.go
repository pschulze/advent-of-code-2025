package main

import "testing"

func TestPartOne(t *testing.T) {
	input := "test_input.txt"
	want := 21
	got := partOne(input)
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	input := "test_input.txt"
	want := 40
	got := partTwo(input)
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}
