package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zod-Exarion/javic/lexer"
	"github.com/zod-Exarion/javic/parser"
)

// INFO: Reads the file, tosses it to the lexer, tosses list of tokens to parser
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

	tokens := lexer.Lex(string(content))
	lexer.DisplayTokens(tokens)
	parser.DisplayStatements(parser.Parse(tokens))
}
