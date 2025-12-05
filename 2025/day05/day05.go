package day05

import (
	"advent-of-code/go_utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

// https://adventofcode.com/2025/day/5

func Run(part *string) {
	go_utils.RunParts(part, "day05/input.txt", part1, part2)
}

type freshRange struct {
	min int
	max int
}

// Parsed the input, returns the array of ranges and array of ingredient Ids
func parseInput(path string) ([]freshRange, []int) {
	rows, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatalf("Error reading input: %s", err)
	}

	var ranges []freshRange
	var ids []int

	for _, row := range rows {
		if row == "" {
			// skip the empty one
			continue
		}
		if strings.Contains(row, "-") {
			rng := strings.Split(row, "-")
			rngMin, _ := strconv.Atoi(rng[0])
			rngMax, _ := strconv.Atoi(rng[1])
			ranges = append(ranges, freshRange{min: rngMin, max: rngMax})
		} else {
			newId, _ := strconv.Atoi(row)
			ids = append(ids, newId)
		}
	}

	return ranges, ids
}

func isIngredientFresh(id int, freshRanges []freshRange) bool {
	for _, rng := range freshRanges {
		if id <= rng.max && id >= rng.min {
			return true
		}
	}

	return false
}

func part1(path string) int {
	fmt.Println("Day 05, Part 1: START")
	result := 0

	timer := go_utils.Timer{}

	timer.Start()

	freshRanges, ingredientIds := parseInput(path)

	for _, id := range ingredientIds {
		if isIngredientFresh(id, freshRanges) {
			result += 1
		}
	}

	timer.End()

	fmt.Printf("day 05, part 1 result: %d | %s\n", result, timer.TimeLapsed())
	return result
}

func combineOverlap(ranges []freshRange) []freshRange {
	combined := []freshRange{ranges[0]}
	for _, curr := range ranges[1:] {
		prev := &combined[len(combined)-1]

		if curr.min <= prev.max {
			if curr.max > prev.max {
				prev.max = curr.max
			}

		} else {
			combined = append(combined, curr)
		}
	}

	return combined
}

func part2(path string) int {
	fmt.Println("Day 05, Part 2: START")
	result := 0

	timer := go_utils.Timer{}

	freshRanges, _ := parseInput(path)

	timer.Start()

	sort.Slice(freshRanges, func(i, j int) bool {
		return freshRanges[i].min < freshRanges[j].min
	})

	combined := combineOverlap(freshRanges)

	for _, rng := range combined {
		result += (rng.max - rng.min) + 1
	}

	timer.End()

	fmt.Printf("day 05, part 1 result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
