package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		take input from user. Input can contain spaces.
		fmt.Scan() and ft.Scanln don't work with input containing spaces.
	*/

	scn := bufio.NewScanner(os.Stdin)
	scn.Scan() // use `for scanner.Scan()` to keep reading
	line := scn.Text()
	fmt.Println("captured:", line)
}
