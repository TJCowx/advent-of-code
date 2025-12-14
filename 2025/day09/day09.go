package day09

import (
	"advent-of-code/go_utils"
	"fmt"
	"image"
	"log"
	"math"
	"sort"
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

type grid struct {
	gridPoints             map[image.Point]rune
	minY, maxY, minX, maxX int
	edgesForY              map[int][]int
}

func (g *grid) buildOutline(points []image.Point) {
	for _, p := range points {
		if p.X < g.minX {
			g.minX = p.X
		}
		if p.X > g.maxX {
			g.maxX = p.X
		}
		if p.Y < g.minY {
			g.minY = p.Y
		}
		if p.Y > g.maxY {
			g.maxY = p.Y
		}
	}

	for i := 0; i < len(points); i++ {
		start := points[i]
		end := points[(i+1)%len(points)]
		g.gridPoints[start] = '#'
		g.gridPoints[end] = '#'

		if start.X == end.X {
			// Vertical, fill it vertically
			lowY := min(start.Y, end.Y)
			highY := max(start.Y, end.Y)
			for y := lowY; y <= highY; y++ {
				pt := image.Pt(start.X, y)
				g.gridPoints[pt] = '#'
				if g.edgesForY[y] == nil {
					g.edgesForY[y] = make([]int, 0)
				}

				g.edgesForY[y] = append(g.edgesForY[y], start.X)
			}

		} else {
			lowX := min(start.X, end.X)
			highX := max(start.X, end.X)

			for x := lowX; x <= highX; x++ {
				pt := image.Pt(x, start.Y)
				g.gridPoints[pt] = '#'
			}
		}
	}

	// Just sort them
	for k := range g.edgesForY {
		sort.Ints(g.edgesForY[k])
	}
}

func (g *grid) isEdgeInPoly(rectEdge edge) bool {
	if rectEdge.horizontal {
		// Navigate across the x axis
		for i := rectEdge.start.X; i <= rectEdge.end.X; i++ {
			if !g.pointIsInPoly(image.Pt(i, rectEdge.start.Y)) {
				return false
			}
		}
	} else {
		// Navigate down the y axis
		for i := rectEdge.start.Y; i <= rectEdge.end.Y; i++ {
			if !g.pointIsInPoly(image.Pt(rectEdge.start.X, i)) {
				return false
			}
		}
	}

	return true
}

// Since the shape isn't in a U or n shape, or even more complicated this
// works to see if it is just in between the points, too lazy to make it for
// all edge cases
func (g *grid) pointIsInPoly(pt image.Point) bool {
	startX := g.edgesForY[pt.Y][0]
	endX := g.edgesForY[pt.Y][len(g.edgesForY[pt.Y])-1]

	return pt.X >= startX && pt.X <= endX
}

func (g *grid) isRectInPoly(rect image.Rectangle) bool {
	// First check if all 4 corners at in the polygon
	corners := []image.Point{
		{X: rect.Min.X, Y: rect.Min.Y},
		{X: rect.Max.X, Y: rect.Min.Y},
		{X: rect.Min.X, Y: rect.Max.Y},
		{X: rect.Max.X, Y: rect.Max.Y},
	}

	for _, corner := range corners {
		if !g.pointIsInPoly(corner) {
			return false
		}
	}

	// Looking at the shape, I don't have to worry abount U or n shapes or even more complicated
	edges := []edge{
		{start: image.Pt(rect.Min.X, rect.Min.Y), end: image.Pt(rect.Min.X, rect.Max.Y), horizontal: false},
		{start: image.Pt(rect.Max.X, rect.Min.Y), end: image.Pt(rect.Max.X, rect.Max.Y), horizontal: false},
	}
	for _, edge := range edges {
		if !g.isEdgeInPoly(edge) {
			return false
		}
	}

	return true
}

func (g *grid) largestArea(positions []image.Point) int {
	largestArea := 0
	for i := 0; i < len(positions)-1; i++ {
		pointA := positions[i]
		for j := i + 1; j < len(positions); j++ {
			pointB := positions[j]
			rect := image.Rect(pointA.X, pointA.Y, pointB.X, pointB.Y)

			area := (rect.Dx() + 1) * (rect.Dy() + 1)
			if largestArea < area {
				if g.isRectInPoly(rect) {
					largestArea = area
				}
			}
		}
	}

	return largestArea
}

func part2(path string) int {
	fmt.Println("Day 09, Part 2: START")

	timer := go_utils.Timer{}

	polygon := parseInput(path)

	timer.Start()

	g := grid{
		gridPoints: make(map[image.Point]rune),
		minY:       math.MaxInt,
		maxY:       0,
		minX:       math.MaxInt,
		maxX:       0,
		edgesForY:  make(map[int][]int),
	}

	g.buildOutline(polygon)

	result := g.largestArea(polygon)

	timer.End()

	fmt.Printf("day 09, part 1 result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
