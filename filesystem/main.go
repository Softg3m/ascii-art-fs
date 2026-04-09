package main

import (
	"fmt"
	"os"
	"filesystem/ascii"
)

func main() {
	args := os.Args[1:]
	text := ""
	banner := "standard"

	if len(args) == 1 {
		text = args[0]
	} else if len(args) == 2 {
		text = args[0]
		banner = args[1]
	} else {
		Usage()
		return
	}

	result := ascii.GenerateAscii(text, banner)
	fmt.Print(result)
}

func Usage() {
	fmt.Println("Usage: go run . [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("EX: go run . something standard")
}
