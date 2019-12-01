package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
Given an integer,n, perform the following conditional actions:

If  is odd, print Weird
If  is even and in the inclusive range of 2 to 5, print Not Weird
If  is even and in the inclusive range of 5 to 20, print Weird
If  is even and greater than 20, print Not Weird
*/

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)

	if N < 1 || N > 100 {
		return
	}

	r := N % 2

	if r == 1 {
		fmt.Println("Weird")
		return
	}

	if r == 0 {
		if (N >= 2 && N <= 5) || N > 20 {
			fmt.Println("Not Weird")
			return
		}

		if N >= 6 && N <= 20 {
			fmt.Println("Weird")
			return
		}
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
