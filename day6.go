package main

import "strings"

func Day6(lines []string) (int, int) {
	return Day6Part1(lines), Day6Part2(lines)
}

func Day6Part1(lines []string) int {
	var flatGrid string

	for _, line := range lines {
		line2 := line + "\n"
		flatGrid += line2
	}

	guardLoc := strings.Index(flatGrid, "^")
	direction := "up"

	steps := 0

	for {
		var nextLoc int

		if direction == "up" {
			nextLoc = guardLoc - len(lines) - 1
		} else if direction == "down" {
			nextLoc = guardLoc + len(lines) + 1
		} else if direction == "left" {
			nextLoc = guardLoc - 1
		} else if direction == "right" {
			nextLoc = guardLoc + 1
		}

		if nextLoc < 0 || nextLoc > len(flatGrid)-1 || flatGrid[nextLoc] == '\n' {
			break
		}

		if flatGrid[nextLoc] == '#' {
			if direction == "up" {
				direction = "right"
			} else if direction == "down" {
				direction = "left"
			} else if direction == "left" {
				direction = "up"
			} else if direction == "right" {
				direction = "down"
			}
		} else {
			if flatGrid[nextLoc] != 'X' {
				fgR := []rune(flatGrid)
				fgR[nextLoc] = 'X'
				flatGrid = string(fgR)
				steps++
			}

			guardLoc = nextLoc
		}
	}

	return steps
}

func Day6Part2(lines []string) int {
	var flatGrid string

	for _, line := range lines {
		line2 := line + "\n"
		flatGrid += line2
	}

	guardLoc := strings.Index(flatGrid, "^")

	possibleGrids := getPossibleGrids(flatGrid, guardLoc)

	sum := 0

	for i, g := range possibleGrids {
		println("Checking grid %d of %d", i, len(possibleGrids))
		if gridIsLoop(g, guardLoc, len(lines)) {
			sum++
		}
	}

	return sum
}

func gridIsLoop(grid string, startLoc int, lineLength int) bool {
	guardLoc := startLoc
	direction := "up"

	steps := 0
	iterations := 0

	for {
		iterations++
		if iterations > lineLength*lineLength {
			return true
		}

		var nextLoc int

		if direction == "up" {
			nextLoc = guardLoc - lineLength - 1
		} else if direction == "down" {
			nextLoc = guardLoc + lineLength + 1
		} else if direction == "left" {
			nextLoc = guardLoc - 1
		} else if direction == "right" {
			nextLoc = guardLoc + 1
		}

		if nextLoc < 0 || nextLoc > len(grid)-1 || grid[nextLoc] == '\n' {
			return false
		}

		if grid[nextLoc] == '#' {
			if direction == "up" {
				direction = "right"
			} else if direction == "down" {
				direction = "left"
			} else if direction == "left" {
				direction = "up"
			} else if direction == "right" {
				direction = "down"
			}
		} else {
			if grid[nextLoc] != 'X' {
				fgR := []rune(grid)
				fgR[nextLoc] = 'X'
				grid = string(fgR)
				steps++
			}

			guardLoc = nextLoc
		}
	}
}

func getPossibleGrids(grid string, exclude int) []string {
	var grids []string

	for i, c := range grid {
		if i == exclude || c == '#' {
			// noop
		} else {
			fgR := []rune(grid)
			fgR[i] = '#'
			grids = append(grids, string(fgR))
		}
	}

	return grids
}
