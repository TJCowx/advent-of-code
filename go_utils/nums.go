package go_utils

import "fmt"

// Combines an array of 0-9 integers into a number
// EG: [2,3,1,3] -> 2313
func CombineSimpleInts(digits []int) (int, error) {
	n := 0
	for _, d := range digits {
		if d > 9 || d < 0 {
			return -1, fmt.Errorf("invalid digit: %d (must be 0-9)", d)
		}
		n = n*10 + d
	}

	return n, nil
}
