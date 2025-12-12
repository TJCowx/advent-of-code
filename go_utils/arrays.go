package go_utils

import "fmt"

// Converts a string of numbers into an array of numbers
// I don't validate for numerical values, so input better be good
func StringAsNumArray(s string) []int {
	nums := make([]int, len(s))

	for i := range s {
		nums[i] = int(s[i] - '0')
	}

	return nums
}

// Removes the last element from an array
// Returns the array, the element, and an error if there is no element to be able to pop
func Pop[T any](s []T) ([]T, T, error) {
	if len(s) == 0 {
		var zero T
		return s, zero, fmt.Errorf("No stack to pop from")
	}

	first := s[len(s)-1]
	return s[:len(s)-1], first, nil
}

func AreSlicesEqual[T string | int | bool](s1 []T, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func CopySlice[T string | int | bool](s []T) []T {
	c := make([]T, len(s))

	copy(c, s)

	return c
}
