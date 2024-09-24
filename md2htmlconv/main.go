package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: markdown-to-html <input.md>")
		return
	}

	inputFile := os.Args[1]
	content, err := os.ReadFile(inputFile)

	if err != nil {
		log.Fatalf("Error Reading file: %v", err)
	}

	htmlContent := Convert(string(content))

	fmt.Println(htmlContent)
}
