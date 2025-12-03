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

func part1(path string) int {
	fmt.Println("Day 01, Part 1: START")
	intructions, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatal(err)
	}

	timer := go_utils.Timer{}
	timer.Start()

	dial := 50
	result := 0

	for _, instruction := range intructions {
		direction := instruction[:1]
		distance, err := strconv.Atoi(instruction[1:])

		if err != nil {
			log.Fatalf("Error parsing input, %s", err)
		}

		if direction == "L" {
			dial -= distance
		} else {
			dial += distance
		}

		if dial >= 100 {
			dial = normalize(dial)
		} else if dial < 0 {
			dial = ((dial % 100) + 100) % 100
		}

		// fmt.Printf("%s | %d\n", instruction, dial)

		if dial == 0 {
			result += 1
		}
	}

	timer.End()

	fmt.Printf("Day 01, Part 1 Result: %d | %s\n", result, timer.TimeLapsed())

	return result
}

func part2(path string) int {
	fmt.Println("Day 01, Part 2: START")

	result := 0

	fmt.Printf("Day 01, Part 2 Result: %d\n", result)

	return 0
}
