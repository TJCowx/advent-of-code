package day03

import (
	"advent-of-code/go_utils"
	"fmt"
	"log"
	"strconv"
)

// https://adventofcode.com/2025/day/3

func Run(part *string) {
	go_utils.RunParts(part, "day03/input.txt", part1, part2)
}

func part1(path string) int {
	fmt.Println("Day 03, Part 1: START")
	result := 0

	banks, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatal(err)
	}

	timer := go_utils.Timer{}
	timer.Start()

	for _, bank := range banks {
		highest := 0
		for i := 0; i < len(bank)-1; i++ {
			for j := i + 1; j < len(bank); j++ {
				currStr := string(bank[i]) + string(bank[j])
				curr, err := strconv.Atoi(currStr)

				if err != nil {
					log.Fatalf("Error converting %s |%s", currStr, err)
				}

				if curr > highest {
					highest = curr
				}
			}
		}

		fmt.Printf("High value %d\n", highest)
		result += highest
	}

	timer.End()

	fmt.Printf("day 03, part 1 result: %d | %s\n", result, timer.TimeLapsed())
	return result
}

func part2(path string) int {
	fmt.Println("Day 03, Part 2: START")
	result := 0

	timer := go_utils.Timer{}
	timer.Start()

	timer.End()
	fmt.Printf("day 03, part 1 result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
