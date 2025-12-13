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

const (
	FlagFFT = 1 << iota
	FlagDAC
)

type recursiveNode struct {
	val      string
	children []string
	path     []string
	flags    int
}

func (n *recursiveNode) buildKey() string {
	key := fmt.Sprintf("%s|%d", n.val, n.flags)

	return key
}

// 290219757077250
func p2Recursive(nodeMap map[string][]string, currNode recursiveNode, visited map[string]int) int {
	sum := 0

	key := currNode.buildKey()

	if val, ok := visited[key]; ok {
		return val
	}

	if currNode.children[0] == "out" {
		res := 0
		if currNode.flags&(FlagFFT|FlagDAC) == (FlagFFT | FlagDAC) {
			res = 1
		}

		visited[key] = res

		return res
	}

	for _, child := range currNode.children {
		flags := currNode.flags
		if currNode.val == "fft" {
			flags |= FlagFFT
		}
		if currNode.val == "dac" {
			flags |= FlagDAC
		}

		childNode := recursiveNode{
			val:      child,
			children: nodeMap[child],
			path:     append(currNode.path, currNode.val),
			flags:    flags,
		}

		sum += p2Recursive(nodeMap, childNode, visited)
	}

	visited[key] = sum

	return sum
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

	root := recursiveNode{
		val:      "svr",
		children: nodeMap["svr"],
		path:     []string{},
	}

	visited := make(map[string]int)

	result := p2Recursive(nodeMap, root, visited)

	timer.End()

	fmt.Printf("day 11, part 2 result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
