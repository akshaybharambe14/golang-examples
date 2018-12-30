/*
Smallest multiple
Problem 5
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package problems

import "fmt"

// SmallestDivisibleByRange -
func SmallestDivisibleByRange(s, e int) {
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
