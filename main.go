package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	if len(args) < 2 || len(args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nExample: go run . something standard")
		os.Exit(0)
	}
	input := args[1]
	style := "standard"
	if len(args) == 3 {
		style = args[2]
	}
	output := PrintAscii(input, style)
	fmt.Print(output)

}
