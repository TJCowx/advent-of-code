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

		startedZero := dial == 0

		if err != nil {
			log.Fatalf("Error parsing input, %s", err)
		}

		if direction == "L" {
			dial -= distance
		} else {
			dial += distance
		}

		passes := 0
		if countPasses && (dial > 100 || dial < -100) {
			passes = int(math.Floor(math.Abs(float64(dial)) / 100.0))
		}

		if dial >= 100 {
			dial = normalize(dial)
		} else if dial < 0 {
			dial = ((dial % 100) + 100) % 100
			if passes == 0 && !startedZero {
				passes += 1
			}
		}

		if dial == 0 {
			result += 1
		}
		result += passes

		fmt.Printf("%s | Pos %d | Result: %d (Started Zero: %v) | Passes: %d\n", instruction, dial, result, startedZero, passes)

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
