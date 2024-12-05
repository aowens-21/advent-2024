package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day1() (int, int) {
	return Day1Part1(), Day1Part2()
}

func Day1Part1() int {
	numPairs := readNums("./input/day1.txt")
	var leftNums []int
	var rightNums []int

	for _, pair := range numPairs {
		numsAsStr := strings.Split(pair, "   ")
		leftNum, err1 := strconv.Atoi(numsAsStr[0])
		rightNum, err2 := strconv.Atoi(numsAsStr[1])

		if err1 != nil || err2 != nil {
			panic("Failed to convert an integer")
		}

		leftNums = append(leftNums, leftNum)
		rightNums = append(rightNums, rightNum)
	}

	sort.Ints(leftNums)
	sort.Ints(rightNums)

	sum := 0

	for i := 0; i < len(leftNums); i++ {
		sum += int(math.Abs(float64(leftNums[i]) - float64(rightNums[i])))
	}

	return sum
}

func Day1Part2() int {
	numPairs := readNums("./input/day1.txt")
	var leftNums []int
	var rightNums []int

	for _, pair := range numPairs {
		numsAsStr := strings.Split(pair, "   ")
		leftNum, err1 := strconv.Atoi(numsAsStr[0])
		rightNum, err2 := strconv.Atoi(numsAsStr[1])

		if err1 != nil || err2 != nil {
			panic("Failed to convert an integer")
		}

		leftNums = append(leftNums, leftNum)
		rightNums = append(rightNums, rightNum)
	}

	rightMap := make(map[int]int)

	for _, num := range rightNums {
		if _, ok := rightMap[num]; ok {
			rightMap[num]++
		} else {
			rightMap[num] = 1
		}
	}

	sum := 0

	for _, num := range leftNums {
		rightCount, ok := rightMap[num]

		if ok {
			sum += num * rightCount
		}
	}

	return sum
}

func readNums(fileName string) []string {
	f, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	var numPairs []string

	scanner := bufio.NewScanner(bufio.NewReader(f))

	for scanner.Scan() {
		line := scanner.Text()
		numPairs = append(numPairs, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading input failed: ", err)
	}

	return numPairs
}
