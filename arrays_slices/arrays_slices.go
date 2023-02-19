package arraysslices

func Sum(n []int) int {
	sum := 0
	for _, number := range n {
		sum += number
	}

	return sum
}

func SumAll(n ...[]int) (sums []int) {
	numbersQuantity := len(n)
	sums = make([]int, numbersQuantity)

	for _, numbers := range n {
		sums = append(sums, Sum(numbers))
	}

	return
}
