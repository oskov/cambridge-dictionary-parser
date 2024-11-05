package main

import (
	"fmt"
	"github.com/oskov/cambridge-dictionary-parser/parser"
	"log"
)

func main() {
	// Example usage with default settings
	parser := parser.NewDictionaryParser()
	wordData, err := parser.ParseWord("example")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Print the scraped definitions and examples for the word
	fmt.Printf("Definitions and examples for the word '%s':\n", wordData.Word)
	for i, def := range wordData.Definitions {
		fmt.Printf("%d: %s\n", i+1, def.Definition)
		if len(def.Examples) > 0 {
			fmt.Println("  Examples:")
			for _, example := range def.Examples {
				fmt.Printf("    - %s\n", example)
			}
		}
	}
}
