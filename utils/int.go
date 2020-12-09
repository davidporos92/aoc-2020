package utils

func SumInts(a []int) int {
	sum := 0

	for i := 0; i < len(a); i++ {
		sum += a[i]
	}

	return sum
}

func MinMax(array []int) (int, int) {
	max := array[0]
	min := array[0]

	for _, value := range array {
		if max < value {
			max = value
		}

		if min > value {
			min = value
		}
	}

	return min, max
}
