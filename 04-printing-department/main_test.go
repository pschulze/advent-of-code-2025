package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testGrid = Grid{
	{'.', '.', '@', '@', '.', '@', '@', '@', '@', '.'},
	{'@', '@', '@', '.', '@', '.', '@', '.', '@', '@'},
	{'@', '@', '@', '@', '@', '.', '@', '.', '@', '@'},
	{'@', '.', '@', '@', '@', '@', '.', '.', '@', '.'},
	{'@', '@', '.', '@', '@', '@', '@', '.', '@', '@'},
	{'.', '@', '@', '@', '@', '@', '@', '@', '.', '@'},
	{'.', '@', '.', '@', '.', '@', '.', '@', '@', '@'},
	{'@', '.', '@', '@', '@', '.', '@', '@', '@', '@'},
	{'.', '@', '@', '@', '@', '@', '@', '@', '@', '.'},
	{'@', '.', '@', '.', '@', '@', '@', '.', '@', '.'},
}

func TestPosAccessible(t *testing.T) {
	type testCase struct {
		grid Grid
		pos  Position
		want bool
	}

	tests := []testCase{
		{grid: testGrid, pos: Position{x: 2, y: 0}, want: true},
		{grid: testGrid, pos: Position{x: 0, y: 9}, want: true},
		{grid: testGrid, pos: Position{x: 9, y: 4}, want: true},
		{grid: testGrid, pos: Position{x: 6, y: 2}, want: true},
		{grid: testGrid, pos: Position{x: 4, y: 9}, want: false},
		{grid: testGrid, pos: Position{x: 5, y: 5}, want: false},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d, %d", tc.pos.x, tc.pos.y), func(t *testing.T) {
			got := tc.grid.posAccessible(tc.pos)
			if got != tc.want {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestRemoveAccessibleRolls(t *testing.T) {
	wantCount := 13
	wantGrid := Grid{
		{'.', '.', '.', '.', '.', '.', '.', '@', '.', '.'},
		{'.', '@', '@', '.', '@', '.', '@', '.', '@', '@'},
		{'@', '@', '@', '@', '@', '.', '.', '.', '@', '@'},
		{'@', '.', '@', '@', '@', '@', '.', '.', '@', '.'},
		{'.', '@', '.', '@', '@', '@', '@', '.', '@', '.'},
		{'.', '@', '@', '@', '@', '@', '@', '@', '.', '@'},
		{'.', '@', '.', '@', '.', '@', '.', '@', '@', '@'},
		{'.', '.', '@', '@', '@', '.', '@', '@', '@', '@'},
		{'.', '@', '@', '@', '@', '@', '@', '@', '@', '.'},
		{'.', '.', '.', '.', '@', '@', '@', '.', '.', '.'},
	}

	gotGrid, gotCount := testGrid.removeAccessibleRolls()
	if gotCount != wantCount {
		t.Errorf("got removed count %d; want %d", gotCount, wantCount)
	}
	if !reflect.DeepEqual(gotGrid, wantGrid) {
		t.Errorf("got grid:\n%v\nwant grid:\n%v", gotGrid, wantGrid)
	}
}
