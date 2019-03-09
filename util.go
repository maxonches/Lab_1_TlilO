package brown_robinson

import "math/rand"

func Max(array []float64) float64 {
	if len(array) == 0 {
		return 0
	}
	max := array[0]
	for _, elem := range array {
		if elem > max {
			max = elem
		}
	}
	return max
}

func Min(array []float64) float64 {
	if len(array) == 0 {
		return 0
	}
	min := array[0]
	for _, elem := range array {
		if elem < min {
			min = elem
		}
	}
	return min
}

func RandFromRange(min, max int) int {
	return rand.Intn(max - min) + min
}
