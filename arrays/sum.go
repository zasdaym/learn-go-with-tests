package main

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

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
