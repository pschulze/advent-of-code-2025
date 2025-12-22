package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide a part number (1 or 2) and an input filename as arguments.")
		return
	}

	part := os.Args[1]
	filename := os.Args[2]

	switch part {
	case "1":
		splitCount := partOne(filename)
		fmt.Println("Total splits:", splitCount)
	case "2":
		pathsCount := partTwo(filename)
		fmt.Println("Total unique paths:", pathsCount)
	default:
		fmt.Println("Invalid part number. Please provide 1.")
		return
	}
}

func partOne(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	beamIdxs := make(map[int]struct{})
	splitCount := -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Handle first line
		if splitCount == -1 {
			for i, ch := range line {
				if ch == 'S' {
					beamIdxs[i] = struct{}{}
				}
			}
			splitCount = 0
			continue
		}

		for i, ch := range line {
			if _, exists := beamIdxs[i]; !exists {
				continue
			}

			if ch == '^' {
				delete(beamIdxs, i)
				beamIdxs[i-1] = struct{}{}
				beamIdxs[i+1] = struct{}{}
				splitCount++
			}
		}
	}

	return splitCount
}

func partTwo(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := make([]string, 0)
	var beamStart int

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		for i, ch := range scanner.Text() {
			if ch == 'S' {
				beamStart = i
				break
			}
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return countPaths(lines, beamStart)
}

func countPaths(lines []string, beamIdx int) int {
	lineLen := len(lines[0])
	beamCounts := make([]int, lineLen)
	beamCounts[beamIdx] = 1
	for _, line := range lines {
		for i, ch := range line {
			switch ch {
			// Noting, beam continues straight down
			case '.':
			// Splitter, add to beam counts
			case '^':
				if i > 0 {
					beamCounts[i-1] += beamCounts[i]
				}
				if i < lineLen-1 {
					beamCounts[i+1] += beamCounts[i]
				}
				beamCounts[i] = 0
			}
		}
	}

	sum := 0
	for _, count := range beamCounts {
		sum += count
	}

	return sum
}

// Recursive function to count unique paths.
// Works on test input, but far too slow for full input.
func down(lines []string, beamIdx int, lineIdx int) int {
	// We've hit the bottom and have completed one unique path
	if lineIdx >= len(lines) {
		return 1
	}

	line := lines[lineIdx]
	ch := line[beamIdx]
	pathCount := 0

	switch ch {
	// Nothing, beam continues straight down
	case '.':
		pathCount += down(lines, beamIdx, lineIdx+1)
	// Splitter, beam splits left and right
	case '^':
		pathCount += down(lines, beamIdx-1, lineIdx+1)
		pathCount += down(lines, beamIdx+1, lineIdx+1)
	}

	return pathCount
}
