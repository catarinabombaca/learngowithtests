package sum

func Sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

func SumAllTails(numbersToSum ...[]int) []int {
	var totalSum []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			continue
		}

		tail := numbers[1:]
		totalSum = append(totalSum, Sum(tail))
	}

	return totalSum
}
