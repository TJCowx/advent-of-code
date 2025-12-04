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
	fmt.Println(bank)
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

func getHighestCombo(bank []int, currI int, currCombo []int, bankSize int) int {
	combo := append(currCombo, bank[currI])
	if len(combo) == 12 {
		res, err := go_utils.CombineSimpleInts(combo)

		if err != nil {
			log.Fatalf("Invalid combo (%d) %s", currCombo, err)
		}

		return res
	}

	// Size guard, if we need more numbers than we have left, lets just leave it
	remaining := 12 - len(combo)
	if remaining > (bankSize - currI) {
		return -1
	}

	highestCombo := 0

	for i := currI + 1; i < bankSize; i++ {
		comboVal := getHighestCombo(bank, i, combo, bankSize)
		if comboVal > highestCombo {
			highestCombo = comboVal
		}
	}

	return highestCombo
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
		highestStart, matchedIs := getHighestStarts(bank)

		highestCombo := 0

		for _, matchedI := range matchedIs {
			combo := getHighestCombo(bank, matchedI, []int{}, len(bank))
			if combo > highestCombo {
				highestCombo = combo
			}
		}

		fmt.Printf("Highest: %d | Indexes: %v | Highest combo: %d\n", highestStart, matchedIs, highestCombo)
		result += highestCombo
	}

	timer.End()
	fmt.Printf("day 03, part 1 result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
