package go_utils

import (
	"fmt"
	"os"
	"strings"
)

type ParsedInput struct {
	day  string
	part *string
}

// Parses the console input for the runner, returns if the input is correct
func ParseUserInput() (ParsedInput, error) {
	args := os.Args

	if len(args) != 2 {
		return ParsedInput{}, fmt.Errorf("Invalid Input, Only 1 Argument. Format should be in 'day-part' or just 'day', eg (1-1 or 1)")
	}

	parts := strings.Split(args[1], "-")

	if len(parts) == 0 || len(parts) > 2 {
		return ParsedInput{}, fmt.Errorf("Incorrect format, must be 'day-part' format or just 'day' (1-1 or 1)")
	}

	if len(parts) == 1 {
		return ParsedInput{
			day:  parts[0],
			part: nil,
		}, nil
	}

	part := parts[1]
	return ParsedInput{
		day:  parts[0],
		part: &part,
	}, nil
}
