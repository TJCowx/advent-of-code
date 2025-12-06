package day06

import (
	"advent-of-code/go_utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// https://adventofcode.com/2025/day/6

func Run(part *string) {
	go_utils.RunParts(part, "day06/input.txt", part1, part2)
}

type mathProblem struct {
	values      []int
	deliminator string
}

type column struct {
	startI      int
	endI        int
	deliminator string
}

func parseInput(rows []string) []mathProblem {
	problems := make([]mathProblem, len(strings.Fields(rows[0])))

	for i, row := range rows {
		for j, value := range strings.Fields(row) {
			if i == len(rows)-1 {
				// Add a deliminator to it since it is the last row
				problems[j].deliminator = value
				continue
			}
			valInt, err := strconv.Atoi(value)

			if err != nil {
				log.Fatalf("Value %s could not be processed: %s", value, err)
			}

			problems[j].values = append(problems[j].values, valInt)
		}
	}

	return problems
}

func solveProblem(problem mathProblem) int {
	if problem.deliminator == "+" {
		return go_utils.SumArr(problem.values)
	}

	return go_utils.MultiplyArr(problem.values)
}

func part1(path string) int {
	fmt.Println("Day 06, Part 1: START")
	result := 0
	timer := go_utils.Timer{}

	rows, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatalf("Error reading input: %s", err)
	}

	timer.Start()

	problems := parseInput(rows)

	for _, p := range problems {
		result += solveProblem(p)
	}

	timer.End()

	fmt.Printf("day 06, part 1 result: %d | %s\n", result, timer.TimeLapsed())
	return result
}

// All deliminators are left aligned to the problem so we know the start of each problem
// And if there is 1 space between each column, the end of it is 2 spaces back from current
func findOperators(delimRow string) []int {
	var indexes []int

	for i, r := range delimRow {
		if r == '+' || r == '*' {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

func parseInputP2(rows []string) []mathProblem {
	var problems []mathProblem

	operatorIndexes := findOperators(rows[len(rows)-1])

	// Start going right to left
	for i := len(operatorIndexes) - 1; i >= 0; i-- {
		currOpI := operatorIndexes[i]
		var colLastI int
		// Guard against last op i, the end of the column would be the length
		if i == len(operatorIndexes)-1 {
			colLastI = len(rows[0]) - 1
		} else {
			// The last column would be 2 previous from the next operator
			// But the last colI is exclusive
			colLastI = operatorIndexes[i+1] - 2
		}

		problem := mathProblem{
			values:      []int{},
			deliminator: string(rows[len(rows)-1][currOpI]),
		}

		// Iterate through the columns
		for colI := colLastI; colI >= currOpI; colI-- {
			// Iterate through the columns, not including last and build the string
			// Ignoring the last I of the row as it is the deliminator
			var built string
			for rowI := 0; rowI < len(rows)-1; rowI++ {
				charVal := string(rows[rowI][colI])
				if charVal != " " {
					built += charVal
				}
			}
			newVal, err := strconv.Atoi(built)

			if err != nil {
				log.Fatalf("Error converting (%s): %s\n", built, err)
			}

			problem.values = append(problem.values, newVal)
		}

		problems = append(problems, problem)
	}

	return problems
}

func part2(path string) int {
	fmt.Println("Day 06, Part 2: START")
	result := 0

	timer := go_utils.Timer{}
	rows, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatalf("Error reading input: %s", err)
	}

	timer.Start()

	problems := parseInputP2(rows)

	for _, p := range problems {
		result += solveProblem(p)
	}

	timer.End()

	fmt.Printf("day 06, part 2 result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
