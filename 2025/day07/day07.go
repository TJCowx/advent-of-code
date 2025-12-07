package day07

import (
	"advent-of-code/go_utils"
	"fmt"
	"log"
	"strings"
)

// https://adventofcode.com/2025/day/7

func Run(part *string) {
	go_utils.RunParts(part, "day07/input.txt", part1, part2)
}

func part1(path string) int {
	fmt.Println("Day 07, Part 1: START")
	result := 0
	timer := go_utils.Timer{}

	rows, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatalf("Error reading input: %s", err)
	}

	timer.Start()

	// get start pos which is directly under the "S"
	startX := strings.Index(rows[0], "S")

	beams := map[int]struct{}{}
	beams[startX] = struct{}{}
	// Ignore the first row
	for y := 1; y < len(rows); y++ {
		// Now loop through the beams
		for beam := range beams {
			// Get the position of the beam and if it's a string we got to add more
			if rows[y][beam] == '^' {
				// Split the beam up
				beams[beam+1] = struct{}{}
				beams[beam-1] = struct{}{}
				delete(beams, beam)
				result += 1
			}
		}
	}

	timer.End()

	fmt.Printf("day 07, part 1 result: %d | %s\n", result, timer.TimeLapsed())
	return result
}

func part2(path string) int {
	fmt.Println("Day 07, Part 2: START")
	result := 0

	timer := go_utils.Timer{}

	timer.Start()

	timer.End()

	fmt.Printf("day 07, part 1 result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
