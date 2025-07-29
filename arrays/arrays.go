package iteration

func Sum(values []int) int {
	var ret int
	for _, number := range values {
		ret += number
	}
	return ret
}

func SumAll(arrays ...[]int) []int{
	var sums []int
	for _, numbers := range arrays {
		sums = append(sums, Sum(numbers))
	}

	return sums
}