package go_utils

import (
	"fmt"
	"strconv"
	"strings"
)

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

func Copy2DSlice[T string | int | bool](s [][]T) [][]T {
	c := make([][]T, len(s))

	copy(c, s)

	return c
}

func Combinations[T any](pool []T, r int) [][]T {
	n := len(pool)
	if r > n || r <= 0 {
		return nil
	}

	var result [][]T
	indices := make([]int, r)
	for i := 0; i < r; i++ {
		indices[i] = i
	}

	for {
		comb := make([]T, r)
		for i, idx := range indices {
			comb[i] = pool[idx]
		}
		result = append(result, comb)

		// Find the rightmost index that can be incremented
		i := r - 1
		for i >= 0 && indices[i] == i+n-r {
			i--
		}
		if i < 0 {
			break
		}

		indices[i]++
		for j := i + 1; j < r; j++ {
			indices[j] = indices[j-1] + 1
		}
	}

	return result
}

// Equiv of python's range function
func Range(n int) []int {
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = i
	}

	return nums
}

func ParseIntoIntSlice(s string) []int {
	s = strings.Trim(s, "[]")
	if s == "" {
		return []int{}
	}

	parts := strings.Fields(s)
	result := make([]int, len(parts))

	for i, p := range parts {
		val, err := strconv.Atoi(p)
		if err != nil {
			panic(fmt.Sprintf("ParseSlice: invalid number %q in %q", p, s))
		}
		result[i] = val
	}

	return result
}
