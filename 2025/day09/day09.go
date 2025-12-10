package day09

import (
	"advent-of-code/go_utils"
	"fmt"
	"image"
	"log"
	"math"
	"strconv"
	"strings"
)

// https://adventofcode.com/2025/day/9

func Run(part *string) {
	go_utils.RunParts(part, "day09/input.txt", part1, part2)
}

type edge struct {
	start, end image.Point
	horizontal bool
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

func isPointInPoly(polyPoints []image.Point, point image.Point, visted map[image.Point]bool) bool {
	if isInPoly, ok := visted[point]; ok {
		return isInPoly
	}

	intersectionCount := 0
	for i := 0; i < len(polyPoints)-1; i++ {
		start := polyPoints[i]
		end := polyPoints[i+1]

		xp, yp := float64(point.X), float64(point.Y)
		x1, y1 := float64(start.X), float64(start.Y)
		x2, y2 := float64(end.X), float64(end.Y)

		// Skip horizontal edges
		if y1 == y2 {
			continue
		}

		// Half-open interval: count intersection if yp in (min, max]
		lowerY := math.Min(y1, y2)
		upperY := math.Max(y1, y2)

		if yp > lowerY && yp <= upperY {
			// Compute intersection X
			xint := x1 + (yp-y1)*(x2-x1)/(y2-y1)
			if xp < xint {
				intersectionCount++
			}
		}
	}

	// If it intersects with an odd amount of edges, its in the polygon
	isInPoly := intersectionCount%2 == 1
	visted[point] = isInPoly

	return isInPoly
}

func isEdgeInPoly(polyPoints []image.Point, rectEdge edge, visited map[image.Point]bool) bool {
	if rectEdge.horizontal {
		// Navigate across the x axis
		for i := rectEdge.start.X; i <= rectEdge.end.X; i++ {
			if !isPointInPoly(polyPoints, image.Pt(i, rectEdge.start.Y), visited) {
				return false
			}
		}
	} else {
		// Navigate down the y axis
		for i := rectEdge.start.Y; i <= rectEdge.end.Y; i++ {
			if !isPointInPoly(polyPoints, image.Pt(rectEdge.start.X, i), visited) {
				return false
			}
		}
	}

	return true
}

func isRectInPoly(polyPoints []image.Point, rect image.Rectangle, visited map[image.Point]bool) bool {
	// First check if all 4 corners at in the polygon
	corners := []image.Point{
		{X: rect.Min.X, Y: rect.Min.Y},
		{X: rect.Max.X, Y: rect.Min.Y},
		{X: rect.Min.X, Y: rect.Max.Y},
		{X: rect.Max.X, Y: rect.Max.Y},
	}

	for _, corner := range corners {
		if !isPointInPoly(polyPoints, corner, visited) {
			return false
		}
	}

	edges := []edge{
		{start: image.Pt(rect.Min.X, rect.Min.Y), end: image.Pt(rect.Max.X, rect.Min.Y), horizontal: true},
		{start: image.Pt(rect.Min.X, rect.Max.Y), end: image.Pt(rect.Max.X, rect.Max.Y), horizontal: true},
		{start: image.Pt(rect.Min.X, rect.Min.Y), end: image.Pt(rect.Min.X, rect.Max.Y), horizontal: false},
		{start: image.Pt(rect.Max.X, rect.Min.Y), end: image.Pt(rect.Max.X, rect.Max.Y), horizontal: false},
	}
	for _, edge := range edges {
		if !isEdgeInPoly(polyPoints, edge, visited) {
			return false
		}
	}

	return true
}

func getLargestArea(positions []image.Point, checkInPoly bool) int {
	largestArea := 0
	visited := make(map[image.Point]bool)
	for i := 0; i < len(positions)-1; i++ {
		pointA := positions[i]
		for j := i + 1; j < len(positions); j++ {
			pointB := positions[j]
			rect := image.Rect(pointA.X, pointA.Y, pointB.X, pointB.Y)

			area := (rect.Dx() + 1) * (rect.Dy() + 1)
			if largestArea < area {
				if !checkInPoly || isRectInPoly(positions, rect, visited) {
					largestArea = area
				}
			}
		}
	}

	return largestArea
}

func part1(path string) int {
	fmt.Println("Day 09, Part 1: START")
	timer := go_utils.Timer{}

	positions := parseInput(path)

	timer.Start()

	result := getLargestArea(positions, false)

	timer.End()

	fmt.Printf("day 09, part 1 result: %d | %s\n", result, timer.TimeLapsed())
	return result
}

func part2(path string) int {
	fmt.Println("Day 09, Part 2: START")

	timer := go_utils.Timer{}

	polygon := parseInput(path)

	timer.Start()

	result := getLargestArea(polygon, true)

	timer.End()

	fmt.Printf("day 09, part 1 result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
