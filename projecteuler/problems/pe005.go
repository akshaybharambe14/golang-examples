/*
Smallest multiple
Problem 5
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

==> 232792560
*/

package problems

import (
	"fmt"
	"golang-examples/projecteuler/utils"
)

// SmallestDivisibleByRangeSimple -
func SmallestDivisibleByRangeSimple(s, e int) {
	n, r, d := e, e-s+1, 0
N:
	for {
		d = 0
		for i := e; i >= s; i-- {
			if n%i == 0 {
				d++
			} else {
				n++
				continue N
			}
		}

		if d == r {
			break
		}
	}

	fmt.Println("P003 - smallest positive number that is evenly divisible by all of the numbers from ", s, " to ", e, ": ", n)
}

// SmallestDivisibleByRange -
func SmallestDivisibleByRange(s, e int) {
	primes, n, d := utils.GetAllPrimeNumbersInRange(s, e), 1, 0 // get all possible prime factors, n is final value, d is prime factor
	for i := range primes {
		d = primes[i]
		for d <= e {
			d *= primes[i] // check max possible value with current prime factor. ex. for 20, we can have 2*2*2*2
		}
		d /= primes[i] // revert extra multiplied value
		n *= d
	}
	fmt.Println(n)
}
