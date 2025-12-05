package day01

import (
	"advent-of-code/go_utils"
	"fmt"
	"log"
	"math"
	"strconv"
)

// https://adventofcode.com/2025/day/1

func Run(part *string) {
	go_utils.RunParts(part, "day01/input.txt", part1, part2)
}

func normalize(n int) int {
	if n >= 100 {
		return n % 100
	}

	if n <= -100 {
		return -(-n % 100)
	}

	return n
}

func solve(path string, timer *go_utils.Timer, countPasses bool) int {
	intructions, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatal(err)
	}

	timer.Start()

	dial := 50
	result := 0

	for _, instruction := range intructions {
		direction := instruction[:1]
		distance, err := strconv.Atoi(instruction[1:])

		if err != nil {
			log.Fatalf("Error parsing input, %s", err)
		}

		movement := distance
		if direction == "L" {
			movement = -distance
		}

		oldDial := dial

		dial += movement

		passes := 0

		if countPasses {
			// Get the number of passes (counts for maginitudes, eg +100 does not count for smaller go over)
			passes = int(math.Abs(float64(dial)) / 100)
			leftover := int(float64(movement)) % 100

			// Handle smaller increments (14 -> -3)
			if oldDial > 0 && oldDial+leftover < 0 {
				passes += 1
			} else if oldDial < 0 && oldDial+leftover > 0 {
				passes += 1
			}
		}

		// Normalize it so we can play with it after
		dial = ((dial % 100) + 100) % 100

		if passes > 0 {
			result += passes
		} else if dial == 0 {
			result += 1
		}

		fmt.Printf("INSTRUCTION: %s | START: %d | MOVE: %d | FINISH: %d | PASSES: %d | RESULT: %d\n", instruction, oldDial, movement, dial, passes, result)

	}

	timer.End()

	return result
}

func part1(path string) int {
	fmt.Println("Day 01, Part 1: START")
	timer := go_utils.Timer{}

	result := solve(path, &timer, false)

	fmt.Printf("Day 01, Part 1 Result: %d | %s\n", result, timer.TimeLapsed())

	return result
}

func part2(path string) int {
	fmt.Println("Day 01, Part 2: START")

	timer := go_utils.Timer{}

	result := solve(path, &timer, true)

	fmt.Printf("Day 01, Part 2 Result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
