package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zod-Exarion/javic/lexer"
)

// INFO: Reads the file, tosses it to the lexer
func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: javic <file.bas>")
	}

	filename := os.Args[1]

	if filename[len(filename)-4:] != ".bas" {
		log.Fatal("Invalid File: Must end with .bas!")
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	lexer.Lex(string(content))
}
