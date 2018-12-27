/*
Largest prime factor
Problem 3
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

*/

package problems

import (
	"fmt"
	"golang-examples/projecteuler/utils"
)

// LargestPrimeFactor -
func LargestPrimeFactor(number int) {
	max := utils.GetSqRoot(number)
	for i := max; i > 1; i-- {
		if number%i == 0 && utils.CheckPrime(i) {
			fmt.Println("P003 - largest prime factor: ", i)
			break
		}
	}
}
