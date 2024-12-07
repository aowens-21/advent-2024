package main

import (
	"slices"
	"strconv"
	"strings"
)

func Day7(lines []string) (int, int) {
	return Day7Part1(lines), Day7Part2(lines)
}

func Day7Part1(lines []string) int {
	sum := 0

	for _, l := range lines {
		sum += checkEquation(l, []string{"+", "*"})
	}

	return sum
}

func Day7Part2(lines []string) int {
	sum := 0

	for _, l := range lines {
		sum += checkEquation(l, []string{"+", "*", "||"})
	}

	return sum
}

func checkEquation(eq string, possibleOperators []string) int {
	if eq == "" {
		return 0
	}

	testVal, nums := getParts(eq)
	operatorPermutations := getOperatorPermutations(len(nums), possibleOperators)

	for _, perm := range operatorPermutations {
		if producesCorrectResult(nums, perm, testVal) {
			return testVal
		}
	}

	return 0
}

func getParts(eq string) (int, []int) {
	parts := strings.Split(eq, ":")

	if len(parts) != 2 {
		panic("Invalid input")
	}

	numsInEqAsStrs := strings.Split(strings.TrimSpace(parts[1]), " ")
	var numsInEq []int
	for _, numStr := range numsInEqAsStrs {
		num, _ := strconv.Atoi(numStr)
		numsInEq = append(numsInEq, num)
	}

	testVal, _ := strconv.Atoi(parts[0])

	return testVal, numsInEq
}

func getOperatorPermutations(operandsLength int, possibleOperators []string) [][]string {
	var permutations [][]string
	var helper func([]string, int)

	helper = func(operands []string, targetLength int) {
		if len(operands) == targetLength {
			permutations = append(permutations, operands)
		} else {
			for i := 0; i < len(possibleOperators); i++ {
				copyOfOperands := slices.Clone(operands)
				helper(append(copyOfOperands, possibleOperators[i]), targetLength)
			}
		}
	}

	helper([]string{}, operandsLength-1)

	return permutations
}

func producesCorrectResult(nums []int, operators []string, targetValue int) bool {
	total := 0

	for i := 0; i < len(operators); i++ {
		nextOperator := operators[i]
		var operand1, operand2 int

		if i == 0 {
			operand1 = nums[0]
			operand2 = nums[1]
		} else {
			operand1 = total
			operand2 = nums[i+1]
		}

		if nextOperator == "+" {
			total = (operand1 + operand2)
		} else if nextOperator == "*" {
			total = (operand1 * operand2)
		} else if nextOperator == "||" {
			n, _ := strconv.Atoi(strconv.Itoa(operand1) + strconv.Itoa(operand2))
			total = n
		}
	}

	return total == targetValue
}
