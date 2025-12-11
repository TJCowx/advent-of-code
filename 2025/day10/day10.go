package day10

import (
	"advent-of-code/go_utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// https://adventofcode.com/2025/day/10

func Run(part *string) {
	go_utils.RunParts(part, "day10/input.txt", part1, part2)
}

type machine struct {
	indicatorLights     []bool
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
			m.indicatorLights = append(m.indicatorLights, r == '#')
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

func part1(path string) int {
	fmt.Println("Day 10, Part 1: START")
	result := 0
	timer := go_utils.Timer{}

	machines := parseInput(path)

	fmt.Println(machines)
	timer.Start()

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
