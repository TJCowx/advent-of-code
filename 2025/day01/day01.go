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

func normalize(dial int) int {
	return ((dial % 100) + 100) % 100
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

		// Get the number of passes (counts for maginitudes, eg +100 does not count for smaller go over)
		passes := int(math.Abs(float64(dial)) / 100)
		leftover := int(float64(dial)) % 100

		if countPasses {

			// Handle smaller increments (14 -> -3)
			if oldDial > 0 && leftover < 0 {
				passes += 1
			}

		}

		// Normalize it so we can play with it after
		dial = normalize(dial)

		if passes > 0 {
			result += passes
		} else if dial == 0 {
			result += 1
		}

		if countPasses {
			fmt.Printf("Instruction: %s | Start: %d | End: %d | PassCount: %d | Leftover: %d | Live Count: %d\n", instruction, oldDial, dial, passes, leftover, result)
		}

		// fmt.Printf("INSTRUCTION: %s | START: %d | MOVE: %d | FINISH: %d | RESULT: %d\n", instruction, oldDial, movement, dial, result)

	}

	timer.End()

	return result
}

func solveP2(instructions []string) int {
	result := 0
	dial := 50

	for _, instruction := range instructions {
		direction := instruction[:1]
		distance, err := strconv.Atoi(instruction[1:])

		if err != nil {
			log.Fatalf("Error parsing input, %s", err)
		}

		dHundreds := int(distance / 100)
		distance = distance % 100

		movement := distance
		if direction == "L" {
			movement = -distance
		}

		// We have already passed nx100 times if the instruction is over 100
		passes := dHundreds

		end := dial + movement

		if movement != 0 {
			if dial > 0 && end <= 0 {
				passes++
			} else if end >= 100 {
				passes++
			}
		}

		dial = normalize(end)

		result += passes
	}

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

	instructions, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatal(err)
	}

	timer := go_utils.Timer{}

	timer.Start()

	// result := solve(path, &timer, true)
	result := solveP2(instructions)

	timer.End()
	fmt.Printf("Day 01, Part 2 Result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
