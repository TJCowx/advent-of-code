package day01

import (
	"advent-of-code/go_utils"
	"fmt"
	"log"
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

func solveP1(path string, timer *go_utils.Timer) int {
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

		dial += movement

		// Normalize it so we can play with it after
		dial = ((dial % 100) + 100) % 100

		if dial == 0 {
			result += 1
		}

		// fmt.Printf("INSTRUCTION: %s | START: %d | MOVE: %d | FINISH: %d | RESULT: %d\n", instruction, oldDial, movement, dial, result)

	}

	timer.End()

	return result
}

func part1(path string) int {
	fmt.Println("Day 01, Part 1: START")
	timer := go_utils.Timer{}

	result := solveP1(path, &timer)

	fmt.Printf("Day 01, Part 1 Result: %d | %s\n", result, timer.TimeLapsed())

	return result
}

func part2(path string) int {
	fmt.Println("Day 01, Part 2: START")

	timer := go_utils.Timer{}
	result := 0

	// result := solve(path, &timer)

	fmt.Printf("Day 01, Part 2 Result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
