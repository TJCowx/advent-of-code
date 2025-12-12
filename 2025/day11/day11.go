package day11

import (
	"advent-of-code/go_utils"
	"fmt"
	"log"
	"strings"
)

// https://adventofcode.com/2025/day/11

func Run(part *string) {
	go_utils.RunParts(part, "day11/input.txt", part1, part2)
}

func parseInput(path string) map[string][]string {
	lines, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatalf("Error reading input: %s", err)
	}

	nodeMap := make(map[string][]string)

	for _, line := range lines {
		parts := strings.Split(line, ":")

		val := parts[0]
		connected := strings.Fields(parts[1])

		nodeMap[val] = connected
	}

	return nodeMap
}

func countPaths(nodeMap map[string][]string, start string) int {
	type item struct {
		val       string
		processed bool
		hitFFT    bool
		hitDAC    bool
	}

	visited := make(map[string]int)

	stack := []item{{val: start, processed: false, hitFFT: false, hitDAC: false}}

	for len(stack) > 0 {
		remaining, node, err := go_utils.Pop(stack)
		stack = remaining

		if err != nil {
			log.Fatal("Failed at popping stack", err)
		}

		if node.processed {
			sum := 0
			for _, child := range nodeMap[node.val] {
				sum += visited[child]
			}
			if node.val == "out" {
				sum++
			}

			visited[node.val] = sum
		} else {
			stack = append(stack, item{
				val:       node.val,
				processed: true,
			})

			for _, child := range nodeMap[node.val] {
				if _, ok := visited[child]; !ok {

					stack = append(stack, item{
						val:       child,
						processed: false,
					})
				}
			}
		}

	}

	return visited[start]
}

func countPathsP2(nodeMap map[string][]string, start string) int {
	type item struct {
		val       string
		processed bool
		hitFFT    bool
		hitDAC    bool
	}

	count := 0
	visited := make(map[string]map[int]int)

	stack := []item{{val: start, processed: false, hitFFT: false, hitDAC: false}}

	for len(stack) > 0 {
		remaining, node, err := go_utils.Pop(stack)
		stack = remaining

		if err != nil {
			log.Fatal("Failed at popping stack", err)
		}

		if node.val == "fft" {
			node.hitFFT = true
		} else if node.val == "dac" {
			node.hitDAC = true
		}

		flags := 0
		if node.hitFFT {
			flags |= 1
		}
		if node.hitDAC {
			flags |= 2
		}

		if visited[node.val] == nil {
			visited[node.val] = make(map[int]int)
		}

		if _, ok := visited[node.val][flags]; ok {
			continue
		}

		if node.processed {
			sum := 0
			for _, child := range nodeMap[node.val] {
				sum += visited[child][flags]
			}

			if node.hitFFT && node.hitDAC {
				sum++
			} else if node.val == "out" {
				sum++
			}

			visited[node.val][flags] = sum
			count += sum
		} else {
			stack = append(stack, item{
				val:       node.val,
				processed: true,
				hitFFT:    node.hitFFT,
				hitDAC:    node.hitDAC,
			})

			for _, child := range nodeMap[node.val] {
				if _, ok := visited[child]; !ok {

					stack = append(stack, item{
						val:       child,
						processed: false,
						hitFFT:    node.hitFFT,
						hitDAC:    node.hitDAC,
					})
				}
			}
		}

	}

	fmt.Println(visited[start][0])

	return count
}

func part1(path string) int {
	fmt.Println("Day 11, Part 1: START")
	timer := go_utils.Timer{}

	nodeMap := parseInput(path)

	timer.Start()

	result := countPaths(nodeMap, "you")

	timer.End()

	fmt.Printf("day 11, part 1 result: %d | %s\n", result, timer.TimeLapsed())
	return result
}

func part2(path string) int {
	fmt.Println("Day 11, Part 2: START")
	timer := go_utils.Timer{}

	nodeMap := parseInput(path)

	timer.Start()

	result := countPathsP2(nodeMap, "svr")

	timer.End()

	fmt.Printf("day 11, part 1 result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
