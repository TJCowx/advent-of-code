package go_utils

import "fmt"

func RunParts(part *string, inputPath string, part1 func(string) int, part2 func(string) int) {
	if part == nil {
		part1(inputPath)
		part2(inputPath)
		return
	}

	switch *part {
	case "1":
		part1(inputPath)
	case "2":
		part2(inputPath)
	default:
		fmt.Println("Invalid Input")
	}
}
