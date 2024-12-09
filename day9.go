package main

import (
	"slices"
	"strconv"
)

func Day9(lines []string) (int, int) {
	return Day9Part1(lines), Day9Part2(lines)
}

func Day9Part1(lines []string) int {
	var fullDisk []string

	mode := "file"
	fileId := 0

	for _, n := range lines[0] {
		num, _ := strconv.Atoi(string(n))
		if mode == "file" {
			for i := 0; i < num; i++ {
				fullDisk = append(fullDisk, strconv.Itoa(fileId))
			}
			mode = "empty"
			fileId++
		} else {
			for i := 0; i < num; i++ {
				fullDisk = append(fullDisk, ".")
			}
			mode = "file"
		}
	}

	chunkStart := len(fullDisk) - 1
	freeStart := 0

	for {
		if freeStart >= chunkStart {
			break
		}

		nextNextStart, nextChunk := getNextChunk(chunkStart, 1, fullDisk)
		nextnextFreeStart, nextFree := getNextNFreeSpaces(freeStart, 1, fullDisk)

		if nextFree[0] >= nextChunk[0] {
			break
		}

		moveChunk(nextChunk, nextFree, fullDisk)

		chunkStart = nextNextStart
		freeStart = nextnextFreeStart
	}

	return getChecksum(fullDisk)
}

func getNextChunk(startPos int, targetSize int, disk []string) (int, []int) {
	currentIndex := startPos
	var nextChunkIs []int
	nextNextChunkStart := -1

	for {
		if disk[currentIndex] == "." {
			currentIndex--
		} else {

			if currentIndex < 0 {
				break
			}

			nextChunkIs = append(nextChunkIs, currentIndex)
			nextNextChunkStart = currentIndex - 1
			break
		}
	}

	return nextNextChunkStart, nextChunkIs
}

func getNextNFreeSpaces(startPos int, targetSize int, disk []string) (int, []int) {
	curI := startPos
	var nextFreeIs []int

	for {
		if curI > len(disk)-1 {
			break
		}

		if disk[curI] == "." {
			nextFreeIs = append(nextFreeIs, curI)
		}

		if len(nextFreeIs) == targetSize {
			break
		}

		curI++
	}

	return curI, nextFreeIs
}

func moveChunk(chunkIs []int, freeIs []int, disk []string) {
	for i := 0; i < len(chunkIs); i++ {
		chunkI := chunkIs[i]

		disk[freeIs[i]] = disk[chunkI]
		disk[chunkI] = "."
	}
}

func getChecksum(disk []string) int {
	sum := 0
	for i, char := range disk {
		if char == "." {
		} else {
			n, _ := strconv.Atoi(char)
			sum += n * i
		}
	}

	return sum
}

func Day9Part2(lines []string) int {
	var fullDisk []string

	mode := "file"
	fileId := 0

	for _, n := range lines[0] {
		num, _ := strconv.Atoi(string(n))
		if mode == "file" {
			for i := 0; i < num; i++ {
				fullDisk = append(fullDisk, strconv.Itoa(fileId))
			}
			mode = "empty"
			fileId++
		} else {
			for i := 0; i < num; i++ {
				fullDisk = append(fullDisk, ".")
			}
			mode = "file"
		}
	}

	chunkStart := len(fullDisk) - 1

	for {
		nextLocToLookForChunk, nextChunk := getContiguousChunk(chunkStart, fullDisk)
		freeSpaces := getFirstContiguousFreeSpace(len(nextChunk), fullDisk)

		if len(freeSpaces) != 0 {
			if slices.Contains(nextChunk, 0) {
				break
			}

			if freeSpaces[0] > nextChunk[0] {
				//noop
			} else {
				moveChunk(nextChunk, freeSpaces, fullDisk)
			}
		}

		chunkStart = nextLocToLookForChunk
	}

	return getChecksum(fullDisk)
}

func getContiguousChunk(start int, disk []string) (int, []int) {
	var indices []int
	fileId := ""
	i := start
	for {
		if i < 0 || (len(indices) > 0 && disk[i] == ".") {
			break
		}

		if disk[i] == "." {
			i--
		} else {
			if fileId == "" {
				indices = append(indices, i)
				fileId = disk[i]
				i--
			} else {
				if disk[i] != fileId {
					break
				}

				indices = append(indices, i)
				i--
			}
		}
	}

	return i, indices
}

func getFirstContiguousFreeSpace(size int, disk []string) []int {
	var indices []int

	for i := 0; i < len(disk); i++ {
		if disk[i] == "." {
			indices = append(indices, i)
		} else {
			indices = []int{}
		}

		if len(indices) == size {
			break
		}
	}

	if len(indices) != size {
		return []int{}
	}

	return indices
}
