package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the solve function below.
func solve(meal_cost float64, tip_percent int32, tax_percent int32) {
	// meal_cost 10.25
	// tip_percent 17
	// tax_percent 5
	// gives 12.504999999999999 i.e with int32(2.504999999999999) returns 12, which mast be 13. Use math.Round()
	tip := (meal_cost * float64(tip_percent)) / 100

	tax := (meal_cost * float64(tax_percent)) / 100

	total := meal_cost + tip + tax

	fmt.Println(total)
	fmt.Println(math.Round(total))
	fmt.Println(int32(math.Round(total)))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	meal_cost, err := strconv.ParseFloat(readLine(reader), 64)
	checkError(err)

	tip_percentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	tip_percent := int32(tip_percentTemp)

	tax_percentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	tax_percent := int32(tax_percentTemp)

	solve(meal_cost, tip_percent, tax_percent)
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
