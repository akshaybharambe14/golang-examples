package main

import (
	"github.com/akshaybharambe14/golang-examples/projecteuler/problems"
)

func main() {
	problems.SumOfMultiples(10)                  //p001
	problems.FibSeqSum(4e6)                      //p002
	problems.LargestPrimeFactor(600851475143)    //p003
	problems.LargestPalindromeFromDigit(3)       //p004
	problems.SmallestDivisibleByRange(1, 20)     //p005
	problems.DiffInSumOfSqAndSqOfSum(10)         //p006
	problems.FindPrimeNumberAtLocation(1, 10001) //p007
}
