package arrays

// Sum calculates the total from a slice of integers
func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

// SumAllTails calculates the sums of all but the first number given a collection of slices
func SumAllTails(numbersToSum ...[]int) (result []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) > 1 {
			tail := numbers[1:]
			result = append(result, Sum(tail))
		} else {
			result = append(result, 0)
		}
	}
	return
}
