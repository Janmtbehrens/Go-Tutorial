package iteration

func Sum(values []int) int {
	var ret int
	for _, number := range values {
		ret += number
	}
	return ret
}