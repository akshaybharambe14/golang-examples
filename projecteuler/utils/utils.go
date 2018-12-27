package utils

import (
	"math"
)

// CheckEven -
func CheckEven(i int) bool {
	return i%2 == 0
}

// CheckPrime -
func CheckPrime(i int) bool {
	// max := i / 2
	max := GetSqRoot(i)
	for j := 2; j <= max; j++ {
		if i%j == 0 {
			return false
		}
	}
	return i > 1
}

// GetSqRoot -
func GetSqRoot(i int) int {
	return int(math.Floor(math.Sqrt(float64(i))))
}
