package day06

import (
	"advent-of-code/go_utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// https://adventofcode.com/2025/day/6

func Run(part *string) {
	go_utils.RunParts(part, "day06/input.txt", part1, part2)
}

type mathProblem struct {
	values      []int
	deliminator string
}

func parseInput(rows []string) []mathProblem {
	problems := make([]mathProblem, len(strings.Fields(rows[0])))

	for i, row := range rows {
		for j, value := range strings.Fields(row) {
			if i == len(rows)-1 {
				// Add a deliminator to it since it is the last row
				problems[j].deliminator = value
				continue
			}
			valInt, err := strconv.Atoi(value)

			if err != nil {
				log.Fatalf("Value %s could not be processed: %s", value, err)
			}

			problems[j].values = append(problems[j].values, valInt)
		}
	}

	return problems
}

func solveProblem(problem mathProblem) int {
	if problem.deliminator == "+" {
		return go_utils.SumArr(problem.values)
	}

	return go_utils.MultiplyArr(problem.values)
}

func part1(path string) int {
	fmt.Println("Day 06, Part 1: START")
	result := 0
	timer := go_utils.Timer{}

	rows, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatalf("Error reading input: %s", err)
	}

	timer.Start()

	problems := parseInput(rows)

	for _, p := range problems {
		result += solveProblem(p)
	}

	timer.End()

	fmt.Printf("day 06, part 1 result: %d | %s\n", result, timer.TimeLapsed())
	return result
}

func part2(path string) int {
	fmt.Println("Day 06, Part 2: START")
	result := 0

	timer := go_utils.Timer{}

	timer.Start()

	timer.End()

	fmt.Printf("day 06, part 1 result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
