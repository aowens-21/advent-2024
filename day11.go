package main

import (
	"maps"
	"strconv"
	"strings"
)

func Day11(lines []string) (int, int) {
	return Day11Part1(lines), Day11Part2(lines)
}

func Day11Part1(lines []string) int {
	return getBlinkCount(lines, 25)
}

func getBlinkCount(lines []string, blinks int) int {
	var arr []uint

	for _, s := range strings.Split(lines[0], " ") {
		i, _ := strconv.Atoi(string(s))
		arr = append(arr, uint(i))
	}

	countMap := buildCountMap(arr)

	for i := 0; i < blinks; i++ {
		countMap = blink(countMap)
	}

	count := 0

	for k := range maps.Keys(countMap) {
		count += countMap[k]
	}

	return count
}

func buildCountMap(arr []uint) map[uint]int {
	m := make(map[uint]int)

	for _, n := range arr {
		count, ok := m[n]

		if ok {
			m[n] = count + 1
		} else {
			m[n] = 1
		}
	}

	return m
}

func blink(countMap map[uint]int) map[uint]int {
	newMap := make(map[uint]int)

	for key := range maps.Keys(countMap) {
		multiplier := countMap[key]

		if key == 0 {
			ones, ok := newMap[1]

			if ok {
				newMap[1] = (ones + (1 * multiplier))
			} else {
				newMap[1] = 1 * multiplier
			}
		} else if len(strconv.FormatUint(uint64(key), 10))%2 == 0 {
			strNum := strconv.FormatUint(uint64(key), 10)

			firstHalf := strNum[:len(strNum)/2]
			secondHalf := strNum[len(strNum)/2:]
			maybeN1, _ := strconv.ParseUint(firstHalf, 10, 64)
			maybeN2, _ := strconv.ParseUint(secondHalf, 10, 64)

			n1 := uint(maybeN1)
			n2 := uint(maybeN2)

			c1, ok1 := newMap[n1]

			if ok1 {
				newMap[n1] = (c1 + (1 * multiplier))
			} else {
				newMap[n1] = 1 * multiplier
			}

			c2, ok2 := newMap[n2]

			if ok2 {
				newMap[n2] = (c2 + (1 * multiplier))
			} else {
				newMap[n2] = 1 * multiplier
			}
		} else {
			newVal := key * 2024
			count, ok := newMap[newVal]

			if ok {
				newMap[newVal] = (count + (1 * multiplier))
			} else {
				newMap[newVal] = 1 * multiplier
			}
		}
	}

	return newMap
}

func Day11Part2(lines []string) int {
	return getBlinkCount(lines, 75)
}
