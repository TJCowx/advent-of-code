package go_utils

func SumArr(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}

	return sum
}

func MultiplyArr(numbers []int) int {
	total := 1

	for _, n := range numbers {
		total *= n
	}

	return total
}
