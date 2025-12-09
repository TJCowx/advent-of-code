package day09

import (
	"advent-of-code/go_utils"
	"fmt"
	"image"
	"log"
	"strconv"
	"strings"
)

// https://adventofcode.com/2025/day/9

func Run(part *string) {
	go_utils.RunParts(part, "day09/input.txt", part1, part2)
}

func parseInput(path string) []image.Point {
	rows, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatalf("Error parsing input: %s", err)
	}

	var positions []image.Point

	for _, row := range rows {
		parts := strings.Split(row, ",")
		xCoor, err := strconv.Atoi(parts[0])

		if err != nil {
			log.Fatalf("Error parsing %s: %s", parts[0], err)
		}

		yCoor, err := strconv.Atoi(parts[1])

		if err != nil {
			log.Fatalf("Error parsing %s: %s", parts[1], err)
		}

		positions = append(positions, image.Pt(xCoor, yCoor))
	}

	return positions
}

func part1(path string) int {
	fmt.Println("Day 09, Part 1: START")
	result := 0
	timer := go_utils.Timer{}

	positions := parseInput(path)

	timer.Start()

	for i := 0; i < len(positions)-1; i++ {
		pointA := positions[i]
		for j := i + 1; j < len(positions); j++ {
			pointB := positions[j]
			rect := image.Rect(pointA.X, pointA.Y, pointB.X, pointB.Y)

			area := (rect.Dx() + 1) * (rect.Dy() + 1)
			if result < area {
				result = area
			}
		}
	}

	timer.End()

	fmt.Printf("day 09, part 1 result: %d | %s\n", result, timer.TimeLapsed())
	return result
}

func part2(path string) int {
	fmt.Println("Day 09, Part 2: START")
	result := 0

	timer := go_utils.Timer{}

	timer.Start()

	timer.End()

	fmt.Printf("day 09, part 1 result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
