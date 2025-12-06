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
	for _, r := range ranges_in {
		grouping := strings.Split(r, "-")
		start, _ := strconv.Atoi(grouping[0])
		endOfRange, err := strconv.Atoi(grouping[1])

		if err != nil {
			log.Fatal(err)
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
		}
	}

	return validIdSum
}

func splitIntoPieces(str string, numParts int) []string {
	partSize := len(str) / numParts
	parts := []string{}

	for i := 0; i < len(str); i += partSize {
		end := i + partSize

		parts = append(parts, str[i:end])
	}

	return parts
}

func areAllEqual(parts []string) bool {
	baseElement := parts[0]
	for i := 1; i < len(parts); i++ {
		if baseElement != parts[i] {
			return false
		}
	}

	return true
}

func isInvalidIdP2(id int) bool {
	idAsStr := strconv.Itoa(id)

	// Single digits can't be invalid
	if len(idAsStr) < 2 {
		return false
	}

	// If all are matching
	if len(idAsStr) >= 2 && strings.Count(idAsStr, string(idAsStr[0])) == len(idAsStr) {
		return true
	}

	// Get all even possibilities
	if len(idAsStr)%2 == 0 {
		pieces := splitIntoPieces(idAsStr, 2)

		if areAllEqual(pieces) {
			return true
		}
	}

	// Every odd number up until half the size (which is probably still overkill)
	for i := 3; i <= (len(idAsStr) / 2); i += 2 {
		if len(idAsStr)%i == 0 {
			pieces := splitIntoPieces(idAsStr, i)
			if areAllEqual(pieces) {
				return true
			}
		}
	}

	return false
}

func sumInvalidIdsInRange(r Range) int {
	validIdSum := 0

	for id := r.start; id <= r.end; id++ {
		if isInvalidIdP2(id) {
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

// INVALID NUMS
// 33832678425
func part2(path string) int {
	fmt.Println("Day 02, Part 2: START")
	timer := go_utils.Timer{}
	result := 0
	timer.Start()

	ranges := getInput(path)

	for _, r := range ranges {
		result += sumInvalidIdsInRange(r)
	}

	timer.End()
	fmt.Printf("day 02, part 1 result: %d | %s\n", result, timer.TimeLapsed())
	return result
}
