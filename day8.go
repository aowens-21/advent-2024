package main

import (
	"fmt"
	"maps"
	"slices"
)

func Day8(lines []string) (int, int) {
	return Day8Part1(lines), Day8Part2(lines)
}

func Day8Part1(lines []string) int {
	var posMap = make(map[string]([][]int))

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] != '\n' && lines[i][j] != '.' {
				positions, ok := posMap[string(lines[i][j])]

				if ok {
					posMap[string(lines[i][j])] = append(positions, []int{i, j})
				} else {
					posMap[string(lines[i][j])] = [][]int{{i, j}}
				}
			}
		}
	}

	var distinctPositions []string

	for k := range maps.Keys(posMap) {
		positions := posMap[k]

		for i, p1 := range positions {
			for j, p2 := range positions {
				if i != j {
					diffRow := p2[0] - p1[0]
					diffCol := p2[1] - p1[1]

					nextPos := []int{p2[0] + diffRow, p2[1] + diffCol}

					if nextPos[0] >= 0 && nextPos[0] < len(lines) {
						if nextPos[1] >= 0 && nextPos[1] < len(lines[0]) {
							if !slices.Contains(distinctPositions, fmt.Sprintf("%d,%d", nextPos[0], nextPos[1])) {
								distinctPositions = append(distinctPositions, fmt.Sprintf("%d,%d", nextPos[0], nextPos[1]))
							}
						}
					}
				}
			}
		}
	}

	return len(distinctPositions)
}

func Day8Part2(lines []string) int {
	var posMap = make(map[string]([][]int))

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] != '\n' && lines[i][j] != '.' {
				positions, ok := posMap[string(lines[i][j])]

				if ok {
					posMap[string(lines[i][j])] = append(positions, []int{i, j})
				} else {
					posMap[string(lines[i][j])] = [][]int{{i, j}}
				}
			}
		}
	}

	var distinctPositions []string

	for k := range maps.Keys(posMap) {
		positions := posMap[k]

		for i, p1 := range positions {
			for j, p2 := range positions {
				if i != j {
					diffRow := p2[0] - p1[0]
					diffCol := p2[1] - p1[1]

					nextPos := []int{p2[0] + diffRow, p2[1] + diffCol}

					if !slices.Contains(distinctPositions, fmt.Sprintf("%d,%d", p1[0], p1[1])) {
						distinctPositions = append(distinctPositions, fmt.Sprintf("%d,%d", p1[0], p1[1]))
					}
					if !slices.Contains(distinctPositions, fmt.Sprintf("%d,%d", p2[0], p2[1])) {
						distinctPositions = append(distinctPositions, fmt.Sprintf("%d,%d", p2[0], p2[1]))
					}
					isLine := false

					for {
						if nextPos[0] >= 0 && nextPos[0] < len(lines) {
							if nextPos[1] >= 0 && nextPos[1] < len(lines[0]) {
								isLine = true
								if !slices.Contains(distinctPositions, fmt.Sprintf("%d,%d", nextPos[0], nextPos[1])) {
									distinctPositions = append(distinctPositions, fmt.Sprintf("%d,%d", nextPos[0], nextPos[1]))
								}

								nextPos[0] = nextPos[0] + diffRow
								nextPos[1] = nextPos[1] + diffCol
							} else {
								break
							}
						} else {
							break
						}
					}

					if isLine {
						prevPos := []int{p1[0] - diffRow, p1[1] - diffCol}
						for {
							if prevPos[0] >= 0 && prevPos[0] < len(lines) {
								if prevPos[1] >= 0 && prevPos[1] < len(lines[0]) {
									if !slices.Contains(distinctPositions, fmt.Sprintf("%d,%d", prevPos[0], prevPos[1])) {
										distinctPositions = append(distinctPositions, fmt.Sprintf("%d,%d", prevPos[0], prevPos[1]))
									}

									prevPos[0] = prevPos[0] - diffRow
									prevPos[1] = prevPos[1] - diffCol
								} else {
									break
								}
							} else {
								break
							}
						}
					}
				}
			}
		}
	}

	return len(distinctPositions)
}
