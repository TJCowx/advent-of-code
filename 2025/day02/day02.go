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
	input, err := go_utils.Read(path)

	if err != nil {
		log.Fatal(err)
	}

	ranges_in := strings.Split(input, ",")

	ranges_out := []Range{}
	for _, r := range ranges_in {
		grouping := strings.Split(r, "-")
		start, _ := strconv.Atoi(grouping[0])
		end, _ := strconv.Atoi(grouping[1])
		ranges_out = append(ranges_out, Range{
			start: start,
			end:   end,
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
			fmt.Printf("INVALID ID: %d\n", id)
			validIdSum += id
		}
	}

	return validIdSum
}

func part1(path string) int {
	fmt.Println("Day 02, Part 1: START")
	ranges := getInput(path)
	timer := go_utils.Timer{}
	result := 0
	timer.Start()

	for _, r := range ranges {
		result += bruteForceRangeP1(r)
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
