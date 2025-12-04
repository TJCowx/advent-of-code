package go_utils

// Converts a string of numbers into an array of numbers
// I don't validate for numerical values, so input better be good
func StringAsNumArray(s string) []int {
	nums := make([]int, len(s))

	for i := range s {
		nums[i] = int(s[i] - '0')
	}

	return nums
}
