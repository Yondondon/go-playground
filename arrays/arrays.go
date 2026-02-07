package arrays

func Sum(numbers []int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}

	return total
}

func SumAll(numbersToSum ...[]int) []int {
	total := make([]int, len(numbersToSum))

	for i, numbers := range numbersToSum {
		total[i] = Sum(numbers)
	}

	return total
}
