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

func getVisitedMapKey(btnsClicked []int, currPattern []bool) string {
	return fmt.Sprintf("%s|%s", printBtnGrouping(btnsClicked), printIndicators(currPattern))
}

func getFastestMatchRecursive(m machine, currPattern []bool, blinkCount int, visited map[string]int, fastest int) int {
	// Disregard if the fastest is done before this
	if blinkCount > fastest {
		return math.MaxInt
	}

	if m.doIndicatorsMatch(currPattern) {
		return blinkCount
	}

	fastestMatch := fastest
	for _, btnGroup := range m.buttons {
		key := getVisitedMapKey(btnGroup, currPattern)

		if knownCount, ok := visited[key]; ok {
			if blinkCount < fastestMatch && knownCount < fastest {
				fastestMatch = (blinkCount + knownCount)
			}
			continue
		}

		blinked := blinkIndicators(currPattern, btnGroup)
		keyBlinked := getVisitedMapKey(btnGroup, blinked)

		visited[keyBlinked] = math.MaxInt
		fastestIter := getFastestMatchRecursive(m, blinked, blinkCount+1, visited, fastestMatch)
		visited[keyBlinked] = fastestIter

		if fastestIter < fastestMatch {
			fastestMatch = fastestIter
		}
	}

	return fastestMatch
}

type state struct {
	pattern []bool
	count   int
}

func solveP1(m machine) int {
	startState := state{
		pattern: make([]bool, len(m.indicatorGoal)),
		count:   0,
	}

	queue := []state{startState}
	visited := map[string]bool{printIndicators(startState.pattern): true}

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]

		if m.doIndicatorsMatch(s.pattern) {
			return s.count
		}

		for _, btn := range m.buttons {
			next := blinkIndicators(s.pattern, btn)
			key := printIndicators(next)

			if !visited[key] {
				visited[key] = true
				queue = append(queue, state{next, s.count + 1})
			}
		}
	}

	return math.MaxInt
}

func part1(path string) int {
	fmt.Println("Day 10, Part 1: START")
	timer := go_utils.Timer{}
	result := 0

	machines := parseInput(path)

	timer.Start()

	for _, m := range machines {
		minSteps := solveP1(m)
		if minSteps < 0 {
			log.Fatalf("Have negative steps! %d", minSteps)
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
