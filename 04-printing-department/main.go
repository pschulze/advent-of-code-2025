package main

import (
	"bufio"
	"fmt"
	"os"
)

const ROW_RUNE = '@'
const EMPTY_RUNE = '.'

type Position struct {
	x int
	y int
}

type Grid [][]rune

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide a part number (1 or 2) and an input filename as arguments.")
		return
	}

	part := os.Args[1]
	filename := os.Args[2]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var inputLines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}

	grid := parseGrid(inputLines)

	switch part {
	case "1":
		_, removedCount := grid.removeAccessibleRolls()
		fmt.Printf("Number of accessible rolls: %d\n", removedCount)

	case "2":
		totalRemoved := 0

		for {
			var removedCount int

			grid, removedCount = grid.removeAccessibleRolls()
			fmt.Printf("removed: %d\n", removedCount)
			totalRemoved += removedCount

			if removedCount == 0 {
				break
			}
		}

		fmt.Printf("Number of rolls removed: %d\n", totalRemoved)

	default:
		fmt.Println("Invalid part number. Please provide 1 or 2.")
		return
	}

}

func parseGrid(input []string) Grid {
	grid := make(Grid, len(input))
	for i, line := range input {
		grid[i] = []rune(line)
	}

	return grid
}

func (g Grid) removeAccessibleRolls() (Grid, int) {
	count := 0
	newGrid := make([][]rune, len(g))
	for i, row := range g {
		newGrid[i] = make([]rune, len(row))
		copy(newGrid[i], row)
	}

	for i, row := range g {
		for j := range row {
			if g[i][j] != ROW_RUNE {
				continue
			}

			if g.posAccessible(Position{x: j, y: i}) {
				count++
				newGrid[i][j] = EMPTY_RUNE
			}
		}
	}

	return newGrid, count
}

func (g Grid) posAccessible(pos Position) bool {
	adjacentOffsets := []Position{
		{x: 0, y: 1},
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: -1, y: 0},
		{x: 1, y: 1},
		{x: 1, y: -1},
		{x: -1, y: 1},
		{x: -1, y: -1},
	}

	rows := len(g)
	cols := len(g[0])
	adjacentCount := 0

	for _, offset := range adjacentOffsets {
		newX := pos.x + offset.x
		newY := pos.y + offset.y
		if newX >= 0 && newX < cols && newY >= 0 && newY < rows && g[newY][newX] == ROW_RUNE {
			adjacentCount++
		}
	}

	return adjacentCount < 4
}
