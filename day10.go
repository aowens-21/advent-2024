package main

import (
	"fmt"
	"slices"
	"strconv"
)

func Day10(lines []string) (int, int) {
	return Day10Part1(lines), Day10Part2(lines)
}

func Day10Part1(lines []string) int {
	grid := readGrid(lines)
	trailheads := getTrailheads(grid)
	score := 0

	for _, trailhead := range trailheads {
		score += getScore(grid, trailhead)
		found = []string{}
	}

	return score
}

var found []string

func getScore(grid [][]string, start []int) int {
	row := start[0]
	col := start[1]
	valAtPos, _ := strconv.Atoi(grid[row][col])

	if grid[row][col] == "9" && !slices.Contains(found, fmt.Sprintf("%d,%d", row, col)) {
		found = append(found, fmt.Sprintf("%d,%d", row, col))
		return 1
	}

	adjacents := getAdjacentPositions(grid, []int{row, col})
	sum := 0

	for _, a := range adjacents {
		numAtA, _ := strconv.Atoi(grid[a[0]][a[1]])

		if numAtA == valAtPos+1 {
			sum += getScore(grid, a)
		}
	}

	return sum
}

func getAdjacentPositions(grid [][]string, pos []int) [][]int {
	var adj [][]int

	row := pos[0]
	col := pos[1]

	// left
	if col-1 >= 0 {
		adj = append(adj, []int{row, col - 1})
	}

	// right
	if col+1 < len(grid[0]) {
		adj = append(adj, []int{row, col + 1})
	}

	// top
	if row-1 >= 0 {
		adj = append(adj, []int{row - 1, col})
	}

	// bottom
	if row+1 < len(grid) {
		adj = append(adj, []int{row + 1, col})
	}

	return adj
}

func Day10Part2(lines []string) int {
	grid := readGrid(lines)
	trailheads := getTrailheads(grid)
	rank := 0

	for _, trailhead := range trailheads {
		rank += getRank(grid, trailhead)
	}

	return rank
}

func getRank(grid [][]string, start []int) int {
	row := start[0]
	col := start[1]
	valAtPos, _ := strconv.Atoi(grid[row][col])

	if grid[row][col] == "9" {
		return 1
	}

	adjacents := getAdjacentPositions(grid, []int{row, col})
	sum := 0

	for _, a := range adjacents {
		numAtA, _ := strconv.Atoi(grid[a[0]][a[1]])

		if numAtA == valAtPos+1 {
			sum += getRank(grid, a)
		}
	}

	return sum
}

func readGrid(lines []string) [][]string {
	var grid [][]string

	for _, l := range lines {
		var row []string
		for _, c := range l {
			row = append(row, string(c))
		}

		grid = append(grid, row)
	}

	return grid
}

func getTrailheads(grid [][]string) [][]int {
	var trailheadPositions [][]int

	for i, r := range grid {
		for j, c := range r {
			if string(c) == "0" {
				trailheadPositions = append(trailheadPositions, []int{i, j})
			}
		}
	}

	return trailheadPositions
}
