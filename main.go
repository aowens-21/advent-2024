package main

import (
	"bufio"
	"fmt"
	"os"
)

var thing map[string]func([]string) (int, int) = make(map[string]func([]string) (int, int))

func main() {
	thing["1"] = Day1
	thing["2"] = Day2
	thing["3"] = Day3
	thing["4"] = Day4
	thing["5"] = Day5
	thing["6"] = Day6

	// Trim off program name
	args := os.Args[1:]

	if len(args) < 1 {
		panic("Please provide a day as a command line argument")
	}

	dayFn, ok := thing[args[0]]

	if ok {
		fileName := fmt.Sprintf("./input/day%s.txt", args[0])

		part1, part2 := dayFn(read(fileName))
		fmt.Printf("Part1: %d\nPart2: %d\n", part1, part2)
	} else {
		fmt.Println("Not implemented!")
	}

}

func read(fileName string) []string {
	f, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	var lines []string

	scanner := bufio.NewScanner(bufio.NewReader(f))

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading input failed: ", err)
	}

	return lines
}
