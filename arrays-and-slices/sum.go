package arraysandslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	var sumSlice []int
	for _, numbers := range slices {
		sumSlice = append(sumSlice, Sum(numbers))
	}
	return sumSlice
}

func SumAllTails(slices ...[]int) []int {
	var tailSumSlice []int
	for _, numbers := range slices {
		tail := numbers[1:]
		tailSumSlice = append(tailSumSlice, Sum(tail))
	}
	return tailSumSlice
}
