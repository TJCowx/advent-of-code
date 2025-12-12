package day10

import (
	"advent-of-code/go_utils"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// https://adventofcode.com/2025/day/10

func Run(part *string) {
	go_utils.RunParts(part, "day10/input.txt", part1, part2)
}

type machine struct {
	indicatorGoal       []bool
	buttons             [][]int
	joltageRequirements []int
}

func parseInput(path string) []machine {
	lines, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatalf("Error parsing input: %s", err)
	}

	var machines []machine

	re := regexp.MustCompile(`\[(.*?)\]\s(.*)\s{([0-9,]+)}`)
	for _, line := range lines {
		var m machine
		matches := re.FindStringSubmatch(line)
		for _, r := range matches[1] {
			m.indicatorGoal = append(m.indicatorGoal, r == '#')
		}

		wiringGroups := strings.Fields(matches[2])
		m.buttons = make([][]int, len(wiringGroups))

		for i, wiringGroup := range wiringGroups {
			buttonsChanged := strings.Split(wiringGroup[1:len(wiringGroup)-1], ",")
			for _, btnStr := range buttonsChanged {
				btnInt, err := strconv.Atoi(btnStr)
				if err != nil {
					log.Fatalf("Error converting button wiring %s: %s", btnStr, err)
				}

				m.buttons[i] = append(m.buttons[i], btnInt)
			}
		}

		for _, joltage := range strings.Split(matches[3], ",") {
			jInt, err := strconv.Atoi(joltage)
			if err != nil {
				log.Fatalf("Error converting joltage %s: %s", joltage, err)
			}
			m.joltageRequirements = append(m.joltageRequirements, jInt)
		}

		machines = append(machines, m)
	}

	return machines
}

func printIndicators(indicators []bool) string {
	out := ""
	for _, val := range indicators {
		if val {
			out += "#"
		} else {
			out += "."
		}
	}
	return out
}

func printBtnGrouping(grouping []int) string {
	out := ""

	for _, val := range grouping {
		if len(out) != 0 {
			out += ","
		}
		out += strconv.Itoa(val)
	}

	return out
}

func (m *machine) doIndicatorsMatch(testIndicators []bool) bool {
	for i, val := range testIndicators {
		if val != m.indicatorGoal[i] {
			return false
		}
	}

	return true
}

func blinkIndicators(curr []bool, flipIs []int) []bool {
	cp := make([]bool, len(curr))
	copy(cp, curr)
	for _, i := range flipIs {
		cp[i] = !cp[i]
	}

	return cp
}

// Iterative DFS to find the shortest path
func (m *machine) getFastestIndicatorMatch() int {
	type node struct {
		prev       []int
		btns       []int
		pattern    []bool
		numPresses int
	}

	// Initiate all values on the stack
	queue := []node{}
	visited := make(map[string]bool)
	offIndicators := make([]bool, len(m.indicatorGoal))
	for _, val := range m.buttons {
		queue = append(queue, node{
			btns:       val,
			pattern:    offIndicators,
			numPresses: 0,
		})
	}

	res := math.MaxInt

	fmt.Println(queue)

	for len(queue) > 0 {
		newQueue, curr, err := go_utils.Pop(queue)
		if err != nil {
			log.Fatalf("Error popping from queue: %s", err)
		}

		queue = newQueue

		if go_utils.AreSlicesEqual(m.indicatorGoal, curr.pattern) && curr.numPresses < res {
			res = curr.numPresses
		}

		for _, next := range m.buttons {
			blinked := blinkIndicators(curr.pattern, next)

			key := fmt.Sprintf("%s|%s", printBtnGrouping(curr.btns), printIndicators(blinked))
			if visited[key] {
				continue
			}

			visited[key] = true

			queue = append(queue, node{
				btns:       next,
				pattern:    blinked,
				numPresses: curr.numPresses + 1,
				prev:       curr.btns,
			})
		}

	}

	return res
}

func part1(path string) int {
	fmt.Println("Day 10, Part 1: START")
	timer := go_utils.Timer{}
	result := 0

	machines := parseInput(path)

	timer.Start()

	for i, m := range machines {
		minSteps := m.getFastestIndicatorMatch()
		fmt.Printf("Machine %d = %d\n", i, minSteps)
		if minSteps < 0 {
			log.Fatalf("Have negative steps!")
		}

		result += minSteps
	}

	timer.End()

	fmt.Printf("day 10, part 1 result: %d | %s\n", result, timer.TimeLapsed())
	return result
}

func part2(path string) int {
	fmt.Println("Day 10, Part 2: START")
	result := 0

	timer := go_utils.Timer{}

	timer.Start()

	timer.End()

	fmt.Printf("day 10, part 1 result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
