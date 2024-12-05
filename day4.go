package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Day4(lines []string) (int, int) {
	return Day4Part1(), Day4Part2()
}

func Day4Part1() int {
	wordSearch := readSearch("./input/day4.txt")
	sum := 0

	for i := 0; i < len(wordSearch); i++ {
		for j := 0; j < len(wordSearch[0]); j++ {
			if wordSearch[i][j] == 'X' {
				sum += checkLeft(wordSearch, j, i)
				sum += checkTop(wordSearch, j, i)
				sum += checkRight(wordSearch, j, i, len(wordSearch[0]))
				sum += checkBottom(wordSearch, j, i, len(wordSearch))
				sum += checkTopLeft(wordSearch, j, i)
				sum += checkTopRight(wordSearch, j, i, len(wordSearch[0]))
				sum += checkBottomRight(wordSearch, j, i, len(wordSearch[0]), len(wordSearch))
				sum += checkBottomLeft(wordSearch, j, i, len(wordSearch))
			}
		}
	}

	return sum
}

func Day4Part2() int {
	wordSearch := readSearch("./input/day4.txt")
	sum := 0

	for i := 0; i < len(wordSearch); i++ {
		for j := 0; j < len(wordSearch[0]); j++ {
			if wordSearch[i][j] == 'A' {
				sum += checkForXmas(wordSearch, j, i, len(wordSearch), len(wordSearch[0]))
			}
		}
	}

	return sum
}

func checkForXmas(wordSearch [][]byte, x int, y int, width int, height int) int {
	topLeft, e1 := safeGetCoord(wordSearch, x-1, y-1, width, height)
	topRight, e2 := safeGetCoord(wordSearch, x+1, y-1, width, height)
	bottomLeft, e3 := safeGetCoord(wordSearch, x-1, y+1, width, height)
	bottomRight, e4 := safeGetCoord(wordSearch, x+1, y+1, width, height)

	if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
		return 0
	}

	masCount := 0

	if (topLeft == 'M' && bottomRight == 'S') || (topLeft == 'S' && bottomRight == 'M') {
		masCount++
	}

	if (topRight == 'M' && bottomLeft == 'S') || (topRight == 'S' && bottomLeft == 'M') {
		masCount++
	}

	if masCount == 2 {
		return 1
	}

	return 0
}

func safeGetCoord(wordSearch [][]byte, x int, y int, width int, height int) (byte, error) {
	if x >= 0 && x < width && y >= 0 && y < height {
		return wordSearch[y][x], nil
	}

	return byte('0'), errors.New("not safe")
}

func checkLeft(wordSearch [][]byte, startX int, startY int) int {
	if startX-1 >= 0 && wordSearch[startY][startX-1] == 'M' {
		if startX-2 >= 0 && wordSearch[startY][startX-2] == 'A' {
			if startX-3 >= 0 && wordSearch[startY][startX-3] == 'S' {
				return 1
			}
		}
	}

	return 0
}

func checkTop(wordSearch [][]byte, startX int, startY int) int {
	if startY-1 >= 0 && wordSearch[startY-1][startX] == 'M' {
		if startY-2 >= 0 && wordSearch[startY-2][startX] == 'A' {
			if startY-3 >= 0 && wordSearch[startY-3][startX] == 'S' {
				return 1
			}
		}
	}

	return 0
}

func checkBottom(wordSearch [][]byte, startX int, startY int, height int) int {
	if startY+1 < height && wordSearch[startY+1][startX] == 'M' {
		if startY+2 < height && wordSearch[startY+2][startX] == 'A' {
			if startY+3 < height && wordSearch[startY+3][startX] == 'S' {
				return 1
			}
		}
	}

	return 0
}

func checkRight(wordSearch [][]byte, startX int, startY int, width int) int {
	if startX+1 < width && wordSearch[startY][startX+1] == 'M' {
		if startX+2 < width && wordSearch[startY][startX+2] == 'A' {
			if startX+3 < width && wordSearch[startY][startX+3] == 'S' {
				return 1
			}
		}
	}

	return 0
}

func checkTopRight(wordSearch [][]byte, startX int, startY int, width int) int {
	if startX+1 < width && startY-1 >= 0 && wordSearch[startY-1][startX+1] == 'M' {
		if startX+2 < width && startY-2 >= 0 && wordSearch[startY-2][startX+2] == 'A' {
			if startX+3 < width && startY-3 >= 0 && wordSearch[startY-3][startX+3] == 'S' {
				return 1
			}
		}
	}

	return 0
}

func checkBottomRight(wordSearch [][]byte, startX int, startY int, width int, height int) int {
	if startX+1 < width && startY+1 < height && wordSearch[startY+1][startX+1] == 'M' {
		if startX+2 < width && startY+2 < height && wordSearch[startY+2][startX+2] == 'A' {
			if startX+3 < width && startY+3 < height && wordSearch[startY+3][startX+3] == 'S' {
				return 1
			}
		}
	}

	return 0
}

func checkTopLeft(wordSearch [][]byte, startX int, startY int) int {
	if startX-1 >= 0 && startY-1 >= 0 && wordSearch[startY-1][startX-1] == 'M' {
		if startX-2 >= 0 && startY-2 >= 0 && wordSearch[startY-2][startX-2] == 'A' {
			if startX-3 >= 0 && startY-3 >= 0 && wordSearch[startY-3][startX-3] == 'S' {
				return 1
			}
		}
	}

	return 0
}

func checkBottomLeft(wordSearch [][]byte, startX int, startY int, height int) int {
	if startX-1 >= 0 && startY+1 < height && wordSearch[startY+1][startX-1] == 'M' {
		if startX-2 >= 0 && startY+2 < height && wordSearch[startY+2][startX-2] == 'A' {
			if startX-3 >= 0 && startY+3 < height && wordSearch[startY+3][startX-3] == 'S' {
				return 1
			}
		}
	}

	return 0
}

func readSearch(fileName string) [][]byte {
	f, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	var lines [][]byte

	scanner := bufio.NewScanner(bufio.NewReader(f))

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []byte(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading input failed: ", err)
	}

	return lines
}
