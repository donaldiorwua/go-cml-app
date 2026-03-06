package main

import (
	"fmt"
	"go-reloaded/internal/file"
	"go-reloaded/internal/processor"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	text := file.ReadFile(inputFile)

	result := processor.Process(text)

	file.WriteFile(outputFile, result)

	fmt.Println("Input: ", inputFile)
	fmt.Println("Output: ", outputFile)

}
