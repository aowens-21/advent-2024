package main

import (
	"fmt"
	"os"
)

var thing map[string]func() (int, int) = make(map[string]func() (int, int))

func main() {
	thing["1"] = Day1
	thing["2"] = Day2
	thing["3"] = Day3
	thing["4"] = Day4
	thing["5"] = Day5

	// Trim off program name
	args := os.Args[1:]

	if len(args) < 1 {
		panic("Please provide a day as a command line argument")
	}

	dayFn, ok := thing[args[0]]

	if ok {
		part1, part2 := dayFn()
		fmt.Printf("Part1: %d\nPart2: %d\n", part1, part2)
	} else {
		fmt.Println("Not implemented!")
	}

}
