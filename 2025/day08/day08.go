package day08

import (
	"advent-of-code/go_utils"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

// https://adventofcode.com/2025/day/8

func p1Wrapper(path string) int {
	return part1(path, 1000)
}

func Run(part *string) {
	go_utils.RunParts(part, "day08/input.txt", p1Wrapper, part2)
}

type junctionBox struct {
	X, Y, Z int
}

type distancePair struct {
	Distance float64
	PointA   junctionBox
	PointB   junctionBox
}

type distanceMapKey struct {
	pointA junctionBox
	pointB junctionBox
}

// Map key is always the smallest first, then the biggest
func makeMapKey(a junctionBox, b junctionBox) distanceMapKey {
	if isLessThan(a, b) {
		return distanceMapKey{pointA: a, pointB: b}
	}

	return distanceMapKey{pointA: b, pointB: a}
}

func isLessThan(a junctionBox, b junctionBox) bool {
	if a.X != b.X {
		return a.X < b.X
	}

	if a.Y != b.Y {
		return a.Y < b.Y
	}

	return a.Z < b.Z
}

func getJunctionBoxes(path string) []junctionBox {
	rows, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatalf("Error parsing input: %s", err)
	}

	var points []junctionBox
	for _, row := range rows {
		valsStrs := strings.Split(row, ",")

		x, err := strconv.Atoi(valsStrs[0])
		if err != nil {
			log.Fatalf("Error parsing value %s: %s", valsStrs[0], err)
		}

		y, err := strconv.Atoi(valsStrs[1])
		if err != nil {
			log.Fatalf("Error parsing value %s: %s", valsStrs[1], err)
		}

		z, err := strconv.Atoi(valsStrs[2])
		if err != nil {
			log.Fatalf("Error parsing value %s: %s", valsStrs[2], err)
		}

		points = append(points, junctionBox{X: x, Y: y, Z: z})

	}

	return points
}

func distancesBetweenBoxes(a junctionBox, b junctionBox) float64 {
	distance := math.Pow(float64(b.X-a.X), 2) +
		math.Pow(float64(b.Y-a.Y), 2) +
		math.Pow(float64(b.Z-a.Z), 2)

	return distance
}

func getDistances(boxes []junctionBox) []distancePair {
	visited := make(map[distanceMapKey]float64)
	var distances []distancePair

	for i := 0; i < len(boxes)-1; i++ {
		for j := i + 1; j < len(boxes); j++ {
			distanceMapKey := makeMapKey(boxes[i], boxes[j])
			if _, ok := visited[distanceMapKey]; ok {
				continue
			}
			distance := distancesBetweenBoxes(boxes[i], boxes[j])
			visited[distanceMapKey] = distance
			distances = append(distances, distancePair{
				Distance: distance,
				PointA:   boxes[i],
				PointB:   boxes[j],
			})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	return distances
}

func getTopNFromMap(in map[int][]junctionBox, topCount int) []int {
	var circuitSizes []int
	for _, boxes := range in {
		circuitSizes = append(circuitSizes, len(boxes))
	}

	sort.Ints(circuitSizes)

	return circuitSizes[len(circuitSizes)-(topCount):]
}

func part1(path string, connectionCount int) int {
	fmt.Println("Day 08, Part 1: START")

	boxes := getJunctionBoxes(path)

	result := 1
	timer := go_utils.Timer{}

	timer.Start()

	distances := getDistances(boxes)

	connections := 0
	// What circuit the junctionbox is connected to
	circuits := make(map[int][]junctionBox)
	visited := make(map[junctionBox]int)
	nextCircuit := 0

	for _, d := range distances {
		aCircuit, aVisited := visited[d.PointA]
		bCircuit, bVisited := visited[d.PointB]

		if aVisited && bVisited {
			if aCircuit != bCircuit {
				// Merge the circuits
				circuits[aCircuit] = append(circuits[aCircuit], circuits[bCircuit]...)
				for _, box := range circuits[bCircuit] {
					visited[box] = aCircuit
				}

				delete(circuits, bCircuit)
			}
		} else if aVisited {
			circuits[aCircuit] = append(circuits[aCircuit], d.PointB)
			visited[d.PointB] = aCircuit
		} else if bVisited {
			circuits[bCircuit] = append(circuits[bCircuit], d.PointA)
			visited[d.PointA] = bCircuit
		} else {
			circuits[nextCircuit] = []junctionBox{d.PointA, d.PointB}
			visited[d.PointA] = nextCircuit
			visited[d.PointB] = nextCircuit

			nextCircuit++
		}

		connections++
		// Break when we do 10 connections
		if connections >= connectionCount {
			break
		}
	}

	// Sort then get the 3 highest values
	for _, val := range getTopNFromMap(circuits, 3) {
		result *= val
	}

	timer.End()

	fmt.Printf("day 08, part 1 result: %d | %s\n", result, timer.TimeLapsed())
	return result
}

func part2(path string) int {
	fmt.Println("Day 08, Part 2: START")
	result := 0

	timer := go_utils.Timer{}

	timer.Start()

	timer.End()

	fmt.Printf("day 08, part 1 result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
