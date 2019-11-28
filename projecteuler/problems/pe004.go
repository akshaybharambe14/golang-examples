/*
Largest palindrome product
Problem 4
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

==> 906609
*/

package problems

import (
	"fmt"
	"math"

	"github.com/akshaybharambe14/golang-examples/projecteuler/utils"
)

// LargestPalindromeFromDigit -
func LargestPalindromeFromDigit(dc int) {
	ul := int(math.Pow(10, float64(dc))) - 1
	ll := int(math.Pow(10, float64(dc-1)))
	i, j, m, lp := ul, ul, 0, 0
	for i >= ll {
		for j >= ll {
			m = i * j
			if utils.CheckNumberPalindrome(m) {
				if m > lp {
					lp = m
				}
			}
			j--
		}
		i--
		j = i
	}
	fmt.Println("P004 - largest palindrome made from the product of two ", dc, "-digit numbers: ", lp)
}
