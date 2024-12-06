package main

import (
	"slices"
	"strconv"
	"strings"
)

func Day5(lines []string) (int, int) {
	return Day5Part1(lines), Day5Part2(lines)
}

func Day5Part1(lines []string) int {
	ruleMap := make(map[int][]int)

	parsingRules := true
	var safeLines [][]int

	for _, l := range lines {
		if len(l) < 2 {
			parsingRules = false
		} else if parsingRules {
			orderingPair := strings.Split(l, "|")
			key, _ := strconv.Atoi(orderingPair[0])
			afterNum, _ := strconv.Atoi(orderingPair[1])

			vals, ok := ruleMap[key]

			if ok {
				ruleMap[key] = append(vals, afterNum)
			} else {
				ruleMap[key] = []int{afterNum}
			}
		} else {
			updateStrs := strings.Split(l, ",")
			var updateNums []int

			for _, str := range updateStrs {
				result, _ := strconv.Atoi(str)
				updateNums = append(updateNums, result)
			}

			if isPrintSafe(updateNums, ruleMap) {
				safeLines = append(safeLines, updateNums)
			}
		}
	}

	sum := 0

	for _, lines := range safeLines {
		sum += getMiddle(lines)
	}

	return sum
}

func isPrintSafe(nums []int, ruleMap map[int][]int) bool {
	for i, num := range nums {
		numsAfter := ruleMap[num]

		for j := i + 1; j < len(nums); j++ {
			if !slices.Contains(numsAfter, nums[j]) {
				return false
			}
		}
	}

	return true
}

func getMiddle(nums []int) int {
	middle := len(nums) / 2

	return nums[middle]
}

func Day5Part2(lines []string) int {
	return -1
}
