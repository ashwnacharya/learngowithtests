package generic_arrays_slices

func Sum(numbers []int) (sum int) {
	
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}


func SumAll(numbersToSum ...[]int) []int {

	var sums []int

	for _, numbers := range numbersToSum {
		sums= append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {

	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(numbers, sumTail, []int{})
}
