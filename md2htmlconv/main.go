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
	outputFile := os.Args[2]
	title := os.Args[3]

	content, err := os.ReadFile(inputFile)

	if err != nil {
		log.Fatalf("Error Reading file: %v", err)
	}

	htmlContent := generateHTMLHead(title) + Convert(string(content)) + wrapHTMLBody()

	err = os.WriteFile(outputFile, []byte(htmlContent), 0644)
	if err != nil {
		log.Fatalf("Error writing file: %v", err);
	}

	fmt.Printf("HTML content saved to %s\n", outputFile)

	openBrowser(outputFile)
}
