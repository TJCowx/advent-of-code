package main

import (
	"advent-of-code/2025/day01"
	"advent-of-code/2025/day02"
	"advent-of-code/2025/day03"
	"advent-of-code/2025/day04"
	"advent-of-code/2025/day05"
	"advent-of-code/2025/day06"
	"advent-of-code/2025/day07"
	"advent-of-code/2025/day08"
	"advent-of-code/2025/day09"
	"advent-of-code/go_utils"
	"fmt"
	"log"
)

func main() {
	input, err := go_utils.ParseUserInput()

	if err != nil {
		log.Fatal(err)
	}

	dayFuncs := map[string]func(*string){
		"1": day01.Run,
		"2": day02.Run,
		"3": day03.Run,
		"4": day04.Run,
		"5": day05.Run,
		"6": day06.Run,
		"7": day07.Run,
		"8": day08.Run,
		"9": day09.Run,
	}

	if runFunc, exists := dayFuncs[input.Day]; exists {
		runFunc(input.Part)
	} else {
		fmt.Printf("Day %s not found", input.Day)
	}
}
