package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zod-Exarion/javic/emitter"
	"github.com/zod-Exarion/javic/external"
	"github.com/zod-Exarion/javic/lexer"
	"github.com/zod-Exarion/javic/parser"
)

// INFO: Reads the file, tosses it to the lexer, tosses list of tokens to parser
func main() {
	argslength := len(os.Args)
	if argslength < 2 {
		log.Fatal("Usage: javic <file.bas>      OR     javic gui")
	}

	if os.Args[1] == "gui" {
		external.NewGUI()
		return
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
	if argslength >= 3 && os.Args[2] == "tokens" {
		fmt.Println(lexer.DisplayTokens(tokens))
	}

	statements := parser.Parse(tokens)
	if argslength >= 4 && os.Args[3] == "ast" {
		fmt.Println(parser.DisplayStatements(statements))
	}

	result := emitter.Emit(statements)

	fmt.Println(result)
}
