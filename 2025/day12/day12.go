package day12

import (
	"advent-of-code/go_utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// https://adventofcode.com/2025/day/12

func Run(part *string) {
	go_utils.RunParts(part, "day12/input.txt", part1, part2)
}

func solve(path string) int {
	lines, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatal(err)
	}

	shapeLines := lines[:29]
	regionLines := lines[30:]

	shapeSizes := make(map[int]int)
	currShapeI := 0
	for i := 0; i < len(shapeLines); i++ {
		if len(shapeLines[i]) == 0 {
			continue
		}

		if strings.HasSuffix(shapeLines[i], ":") {
			shapeNum, err := strconv.Atoi(string(shapeLines[i][0]))
			if err != nil {
				log.Fatal("Failed parsing shapeline", shapeLines[i])
			}
			currShapeI = shapeNum
			shapeSizes[currShapeI] = 0
		} else {
			shapeSizes[currShapeI] += strings.Count(shapeLines[i], "#")
		}

	}

	canFitCount := 0

	for _, line := range regionLines {
		splitSects := strings.Split(line, ":")

		areaDimensions := strings.Split(splitSects[0], "x")

		w, _ := strconv.Atoi(areaDimensions[0])
		h, _ := strconv.Atoi(areaDimensions[1])

		maxArea := w * h
		currArea := 0

		for i, countStr := range strings.Fields(splitSects[1]) {
			count, _ := strconv.Atoi(countStr)
			currArea += (shapeSizes[i] * count)
		}

		if currArea <= maxArea {
			canFitCount++
		}
	}

	return canFitCount
}

func part1(path string) int {
	fmt.Println("Day 12, Part 1: START")
	timer := go_utils.Timer{}

	timer.Start()

	result := solve(path)

	timer.End()

	fmt.Printf("day 12, part 1 result: %d | %s\n", result, timer.TimeLapsed())
	return result
}

func part2(path string) int {
	fmt.Println("Day 12, Part 2: START")
	result := 0

	timer := go_utils.Timer{}

	timer.Start()

	timer.End()

	fmt.Printf("day 12, part 1 result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
