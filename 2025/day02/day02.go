package day02

import (
	"advent-of-code/go_utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// https://adventofcode.com/2025/day/2

func Run(part *string) {
	go_utils.RunParts(part, "day02/input.txt", part1, part2)
}

type Range struct {
	start int
	end   int
}

func getInput(path string) []Range {
	input, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatal(err)
	}

	ranges_in := strings.Split(input[0], ",")

	ranges_out := []Range{}
	for i, r := range ranges_in {
		grouping := strings.Split(r, "-")
		start, _ := strconv.Atoi(grouping[0])
		endOfRange, err := strconv.Atoi(grouping[1])

		if err != nil {
			log.Fatal(err)
		}

		if i == len(ranges_in)-1 {
			fmt.Printf("Uhm okay: %s | START: %d, END: %d | %s, %s", r, start, endOfRange, grouping[0], grouping[1])
		}
		ranges_out = append(ranges_out, Range{
			start: start,
			end:   endOfRange,
		})
	}

	return ranges_out
}

func bruteForceRangeP1(r Range) int {
	validIdSum := 0
	for id := r.start; id <= r.end; id++ {
		idAsStr := strconv.Itoa(id)
		// Odd lengths don't impact this, it can't be repeated
		if len(idAsStr)%2 != 0 {
			continue
		}

		mid := len(idAsStr) / 2

		if idAsStr[:mid] == idAsStr[mid:] {
			validIdSum += id
			// fmt.Printf("INVALID ID: %d (%d)\n", id, validIdSum)
		}
	}

	return validIdSum
}

func part1(path string) int {
	fmt.Println("Day 02, Part 1: START")
	ranges := getInput(path)
	fmt.Println(ranges)
	timer := go_utils.Timer{}
	result := 0
	timer.Start()

	for _, r := range ranges {
		result += bruteForceRangeP1(r)
		// fmt.Printf("New Total: %d\n", result)
	}
	timer.End()
	fmt.Printf("day 02, part 1 result: %d | %s\n", result, timer.TimeLapsed())
	return result
}

func part2(path string) int {
	fmt.Println("Day 02, Part 2: START")
	timer := go_utils.Timer{}
	result := 0
	timer.Start()

	timer.End()
	fmt.Printf("day 02, part 1 result: %d | %s\n", result, timer.TimeLapsed())
	return 0
}
