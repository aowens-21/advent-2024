package main

import (
	"os"
	"regexp"
	"strconv"
)

func Day3() (int, int) {
	return Day3Part1(), Day3Part2()
}

func Day3Part1() int {
	programText := getFileContent("./input/day3.txt")
	mulRegex := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	matchedBytes := mulRegex.FindAllSubmatch(programText, -1)

	result := 0

	for _, matches := range matchedBytes {
		num1, err1 := strconv.Atoi(string(matches[1]))
		num2, err2 := strconv.Atoi(string(matches[2]))

		if err1 != nil || err2 != nil {
			panic("Failed to parse number")
		}

		result += (num1 * num2)
	}

	return result
}

func Day3Part2() int {
	programText := getFileContent("./input/day3.txt")
	mulRegex := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)|do\(\)|don\'t\(\)`)
	matchedBytes := mulRegex.FindAllSubmatch(programText, -1)

	enabled := true
	result := 0

	for _, matches := range matchedBytes {
		if string(matches[0]) == "do()" {
			enabled = true
		} else if string(matches[0]) == "don't()" {
			enabled = false
		} else {
			if enabled {
				num1, err1 := strconv.Atoi(string(matches[1]))
				num2, err2 := strconv.Atoi(string(matches[2]))

				if err1 != nil || err2 != nil {
					panic("Failed to parse number")
				}

				result += (num1 * num2)
			}
		}
	}

	return result
}

func getFileContent(fileName string) []byte {
	bytes, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	return bytes
}
