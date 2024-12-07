package main

import (
	"slices"
	"strings"
)

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

	possibleGrids := getPossibleGrids(flatGrid, guardLoc, lines)

	sum := 0

	for _, g := range possibleGrids {
		if gridIsLoop([]rune(g), guardLoc, len(lines)) {
			sum++
		}
	}

	return sum
}

func gridIsLoop(grid []rune, startLoc int, lineLength int) bool {
	guardLoc := startLoc
	direction := "up"

	visitedDirMap := make(map[int][]string)

	steps := 0

	for {
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

		directionsVisitedThisLocFrom, ok := visitedDirMap[nextLoc]

		if ok {
			if slices.Contains(directionsVisitedThisLocFrom, direction) {
				return true
			}
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
			if ok {
				visitedDirMap[nextLoc] = append(directionsVisitedThisLocFrom, direction)
			} else {
				visitedDirMap[nextLoc] = []string{direction}
			}

			if grid[nextLoc] != 'X' {
				grid[nextLoc] = 'X'
				steps++
			}

			guardLoc = nextLoc
		}
	}
}

func getPossibleGrids(grid string, guardLoc int, lines []string) []string {
	var grids []string
	flatGrid := strings.Clone(grid)

	direction := "up"

	var indicesVisited []int

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

				if nextLoc != guardLoc {
					indicesVisited = append(indicesVisited, nextLoc)
				}
			}
			guardLoc = nextLoc
		}
	}

	for _, index := range indicesVisited {
		fgR := []rune(grid)
		fgR[index] = '#'
		grids = append(grids, string(fgR))
	}

	return grids
}
