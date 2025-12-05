package day04

import (
	"advent-of-code/go_utils"
	"fmt"
	"image"
	"log"
)

// https://adventofcode.com/2025/day/4

type grid struct {
	positions [][]string
	visited   map[image.Point]bool
	bounds    image.Rectangle
}

func (g *grid) isInBounds(p image.Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < g.bounds.Max.X && p.Y < g.bounds.Max.Y
}

func (g *grid) doesPointHaveRoll(p image.Point) bool {
	return g.isInBounds(p) && g.positions[p.Y][p.X] == "@"
}

func (g *grid) removeVisited() {
	for point := range g.visited {
		g.positions[point.Y][point.X] = "."
	}

	g.visited = make(map[image.Point]bool)
}

func (g *grid) adjRollCount(p image.Point) int {
	res := 0

	// UpLeft
	upLeft, _ := go_utils.GetNextDir(p, go_utils.UpLeft)
	if g.doesPointHaveRoll(upLeft) {
		res += 1
	}
	// Up
	up, _ := go_utils.GetNextDir(p, go_utils.Up)
	if g.doesPointHaveRoll(up) {
		res += 1
	}
	// UpRight
	upRight, _ := go_utils.GetNextDir(p, go_utils.UpRight)
	if g.doesPointHaveRoll(upRight) {
		res += 1
	}
	// Left
	left, _ := go_utils.GetNextDir(p, go_utils.Left)
	if g.doesPointHaveRoll(left) {
		res += 1
	}
	// Right
	right, _ := go_utils.GetNextDir(p, go_utils.Right)
	if g.doesPointHaveRoll(right) {
		res += 1
	}
	// DownLeft
	downLeft, _ := go_utils.GetNextDir(p, go_utils.DownLeft)
	if g.doesPointHaveRoll(downLeft) {
		res += 1
	}
	// Down
	down, _ := go_utils.GetNextDir(p, go_utils.Down)
	if g.doesPointHaveRoll(down) {
		res += 1
	}
	// DownRight
	downRight, _ := go_utils.GetNextDir(p, go_utils.DownRight)
	if g.doesPointHaveRoll(downRight) {
		res += 1
	}

	return res
}

func parseInput(rows []string) grid {
	positions := make([][]string, len(rows))
	for i, row := range rows {
		positions[i] = make([]string, len(row))
		for j, char := range row {
			positions[i][j] = string(char)
		}
	}

	return grid{
		positions: positions,
		visited:   make(map[image.Point]bool),
		bounds:    image.Rect(0, 0, len(positions[0]), len(positions)),
	}
}

func Run(part *string) {
	go_utils.RunParts(part, "day04/input.txt", part1, part2)
}

func part1(path string) int {
	fmt.Println("Day 04, Part 1: START")
	result := 0
	timer := go_utils.Timer{}

	rows, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatalf("Error reading input: %s", err)
	}

	g := parseInput(rows)

	timer.Start()

	for y, row := range g.positions {
		for x, val := range row {
			if val == "@" {
				adjRollRound := g.adjRollCount(image.Point{X: x, Y: y})

				if adjRollRound < 4 {
					result += 1
				}
			}

		}
	}

	timer.End()

	fmt.Printf("day 04, part 1 result: %d | %s\n", result, timer.TimeLapsed())
	return result
}

func part2(path string) int {
	fmt.Println("Day 04, Part 2: START")
	result := 0

	timer := go_utils.Timer{}

	rows, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatalf("Error parsing input: %s", err)
	}

	g := parseInput(rows)

	timer.Start()

	for true {
		removedRolls := 0

		for y, row := range g.positions {
			for x, val := range row {
				if val == "@" {
					p := image.Point{X: x, Y: y}
					adjRollRound := g.adjRollCount(p)

					if adjRollRound < 4 {
						g.visited[p] = true
						removedRolls += 1
					}
				}

			}
		}

		g.removeVisited()

		if removedRolls == 0 {
			break
		}
		result += removedRolls
	}

	timer.End()
	fmt.Printf("day 04, part 1 result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
