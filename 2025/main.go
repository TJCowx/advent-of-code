package main

import (
	"advent-of-code/2025/day01"
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
	}

	if runFunc, exists := dayFuncs[input.Day]; exists {
		runFunc(input.Part)
	} else {
		fmt.Printf("Day %s not found", input.Day)
	}
}
