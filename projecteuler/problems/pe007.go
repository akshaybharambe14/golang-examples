/*
10001st prime
Problem 7
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?

==> 104743
*/

package problems

import (
	"fmt"
	"golang-examples/projecteuler/utils"
)

// FindPrimeNumberAtLocation - s is start position, p is required prime position
func FindPrimeNumberAtLocation(s, p int) {
	if s < 2 {
		s = 2
	}
	pc, n := 0, 0 // pc is prime counter, n is result
	for pc < p {
		if utils.CheckPrime(s) {
			pc++
		}
		s++
	}
	n = s - 1
	fmt.Println("P007 - the ", p, " prime number: ", n)
}
