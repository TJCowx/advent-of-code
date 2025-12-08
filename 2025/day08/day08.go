package day08

import (
	"advent-of-code/go_utils"
	"fmt"
)

// https://adventofcode.com/2025/day/8

func Run(part *string) {
	go_utils.RunParts(part, "day08/input.txt", part1, part2)
}

func part1(path string) int {
	fmt.Println("Day 08, Part 1: START")
	result := 0
	timer := go_utils.Timer{}

	timer.Start()

	timer.End()

	fmt.Printf("day 08, part 1 result: %d | %s\n", result, timer.TimeLapsed())
	return result
}

func part2(path string) int {
	fmt.Println("Day 08, Part 2: START")
	result := 0

	timer := go_utils.Timer{}

	timer.Start()

	timer.End()

	fmt.Printf("day 08, part 1 result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
