/*
Sum square difference
Problem 6
The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

==>
*/

package problems

import "fmt"

// DiffInSumOfSqAndSqOfSum - l is limit
func DiffInSumOfSqAndSqOfSum(l int) {
	diff, sumOfSq, sqOfSum, s := 0, 0, 0, 1
	for s <= l {
		sumOfSq += (s * s)
		sqOfSum += s
		s++
	}

	if d := sumOfSq - (sqOfSum * sqOfSum); d < 0 {
		diff = -d
	} else {
		diff = d
	}

	fmt.Println("P006 - difference between the sum of the squares of the first ", l, "natural numbers and the square of the sum: ", diff)
}
