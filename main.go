package main

import (
	"fmt"
	"os"
	"strings"
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

	data, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	fileLines := strings.Split(string(data), "\n")
	result := ""
	words := strings.Split(text, "\\n")

	for _, word := range words {
		for i := 1; i <= 8; i++ {
			for _, char := range word {
				asciiIndex := int(char) - 32
				start := asciiIndex * 9

				if start+i >= len(fileLines) {
					continue
				}

				result += fileLines[start+i]
			}
			result += "\n"
		}
	}
	fmt.Println(result) 
}

func Usage() {
	fmt.Println("Usage: go run . [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("EX: go run . something standard")
}
