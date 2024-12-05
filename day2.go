package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day2(lines []string) (int, int) {
	return Day2Part1(), Day2Part2()
}

func Day2Part1() int {
	reports := readReports("./input/day2.txt")

	safeReports := 0

	for _, report := range reports {
		safeReports += isSafeReport(report)
	}

	return safeReports
}

func Day2Part2() int {
	reports := readReports("./input/day2.txt")

	safeReports := 0

	for _, report := range reports {
		safeReports += isSafeReportV2(report)
	}

	return safeReports
}

func isSafeReportV2(report []string) int {
	adjacentReports := buildFullReportList(report)

	for _, r := range adjacentReports {
		if isSafeReport(r) == 1 {
			return 1
		}
	}

	return 0
}

func buildFullReportList(report []string) [][]string {
	var fullReports [][]string

	fullReports = append(fullReports, report)

	for i := 0; i < len(report); i++ {
		var modifiedReport []string

		for j := 0; j < len(report); j++ {
			if i != j {
				modifiedReport = append(modifiedReport, report[j])
			}
		}

		fullReports = append(fullReports, modifiedReport)
	}

	return fullReports
}

func isSafeReport(report []string) int {
	var isIncrementing bool

	for i := 0; i < len(report)-1; i++ {
		num, err := strconv.Atoi(report[i])
		next, err2 := strconv.Atoi(report[i+1])

		if err != nil || err2 != nil {
			panic("Failed to convert string to int")
		}

		if num == next || !isSafeDiff(num, next) {
			return 0
		}

		if i == 0 {
			// Set inc or dec flag
			if num < next {
				isIncrementing = true
			} else {
				isIncrementing = false
			}
		} else {
			if isIncrementing && next < num {
				return 0
			} else if !isIncrementing && next > num {
				return 0
			}
		}
	}

	return 1
}

func isSafeDiff(num1 int, num2 int) bool {
	diff := math.Abs(float64(num1 - num2))

	return diff > 0 && diff < 4
}

func readReports(fileName string) [][]string {
	f, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	var reports [][]string

	scanner := bufio.NewScanner(bufio.NewReader(f))

	for scanner.Scan() {
		line := scanner.Text()
		reports = append(reports, strings.Split(line, " "))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading input failed: ", err)
	}

	return reports
}
