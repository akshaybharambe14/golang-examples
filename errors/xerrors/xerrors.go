package main

import (
	"fmt"

	"golang.org/x/xerrors"
)

var (
	errorFromLevelOne = xerrors.New("level one failed!")
	errorFromLevelTwo = xerrors.New("level two failed!")
)

func main() {
	printError(levelOne(true))
	printError(levelOne(false))
	printError(nil)
}

func levelOne(ok bool) error {
	if !ok {
		return errorFromLevelOne
	}
	return levelTwo("")
}

func levelTwo(s string) error {
	if s == "" {
		// this format is important. don't remove ':' before '%w', fails to identify the error
		return xerrors.Errorf("%q: %w", s, errorFromLevelTwo)
	}
	return nil
}

func printError(err error) {
	switch {
	case err == nil:
		fmt.Println("no error occurred")
	case xerrors.Is(err, errorFromLevelOne):
		fmt.Println("Level one failed: ", err)
	case xerrors.Is(err, errorFromLevelTwo):
		fmt.Println("Level two failed: ", xerrors.Unwrap(err), " Actual: ", err)
	default:
		fmt.Println("unknown error occurred")
	}
}
