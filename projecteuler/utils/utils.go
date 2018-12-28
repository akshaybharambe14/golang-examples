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

// CheckNumberPalindrome -
func CheckNumberPalindrome(i int) bool {
	i = GetPositive(i)
	is := GetNumberInSlice(i)
	l := len(is)
	for i := 0; i < l/2; i++ {
		if is[i] != is[l-i-1] {
			return false
		}
	}
	return true
}

// CheckStringPalindrome -
func CheckStringPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

// GetNumberInSlice -
func GetNumberInSlice(i int) []int {
	i = GetPositive(i)
	lenInt := GetIntLen(i)
	intSlice := make([]int, lenInt)
	for i > 0 {
		intSlice[lenInt-1] = (i % 10)
		i /= 10
		lenInt--
	}

	return intSlice
}

// GetIntLen -
func GetIntLen(i int) int {
	i = GetPositive(i)
	count := 0
	for i > 0 {
		i /= 10
		count++
	}
	return count
}

// GetPositive -
func GetPositive(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
