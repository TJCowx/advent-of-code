package day03

import (
	"advent-of-code/go_utils"
	"fmt"
	"log"
	"strconv"
)

// https://adventofcode.com/2025/day/3

func Run(part *string) {
	go_utils.RunParts(part, "day03/input.txt", part1, part2)
}

func part1(path string) int {
	fmt.Println("Day 03, Part 1: START")
	result := 0

	banks, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatal(err)
	}

	timer := go_utils.Timer{}
	timer.Start()

	for _, bank := range banks {
		highest := 0
		for i := 0; i < len(bank)-1; i++ {
			for j := i + 1; j < len(bank); j++ {
				currStr := string(bank[i]) + string(bank[j])
				curr, err := strconv.Atoi(currStr)

				if err != nil {
					log.Fatalf("Error converting %s |%s", currStr, err)
				}

				if curr > highest {
					highest = curr
				}
			}
		}

		fmt.Printf("High value %d\n", highest)
		result += highest
	}

	timer.End()

	fmt.Printf("day 03, part 1 result: %d | %s\n", result, timer.TimeLapsed())
	return result
}

func getHighestStarts(bank []int) (int, []int) {
	highestStart, matchedIs := 0, []int{}
	for i := 0; i < len(bank)-12; i++ {
		val := bank[i]

		if val > highestStart {
			highestStart = val
			matchedIs = []int{i}
		} else if val == highestStart {
			matchedIs = append(matchedIs, i)
		}
	}

	return highestStart, matchedIs
}

func getComboVal(bank []int, currI int, remaining int, known map[string]int) int {
	if remaining == 1 {
		return bank[currI]
	}

	mapKey := fmt.Sprintf("%d|%d", currI, remaining)
	if val, ok := known[mapKey]; ok {
		return val
	}

	highestVal := 0
	for next := currI + 1; next <= len(bank)-(remaining-1); next++ {
		val := getComboVal(bank, next, remaining-1, known)
		comboValStr := fmt.Sprintf("%d%d", bank[currI], val)
		comboVal, _ := strconv.Atoi(comboValStr)

		if comboVal > highestVal {
			highestVal = comboVal
		}
	}

	known[mapKey] = highestVal
	return highestVal
}

func getHighestCombo(bank []int, startI int) int {
	highest := 0
	known := make(map[string]int)

	for i := startI; i <= len(bank)-12; i++ {
		val := getComboVal(bank, i, 12, known)
		if val > highest {
			highest = val
		}
	}

	return highest
}

func part2(path string) int {
	fmt.Println("Day 03, Part 2: START")
	result := 0

	banks, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatal(err)
	}

	timer := go_utils.Timer{}
	timer.Start()

	for _, bankStr := range banks {
		bank := go_utils.StringAsNumArray(bankStr)
		_, matchedIs := getHighestStarts(bank)

		highestCombo := 0

		for _, matchedI := range matchedIs {
			combo := getHighestCombo(bank, matchedI)
			if combo > highestCombo {
				highestCombo = combo
			}
		}

		result += highestCombo
	}

	timer.End()
	fmt.Printf("day 03, part 1 result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
