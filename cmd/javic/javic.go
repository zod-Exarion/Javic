package javic

import (
	"fmt"
	"javic/core/lexer"
	"javic/core/tokenizer"
	"os"
)

func Javic(fileName string) {
	// Read the QBASIC File
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	lex := lexer.NewLexer(string(content))
	for range content {
		tok := lex.GetToken()
		if tok.Type == tokenizer.EOF {
			break
		}
		fmt.Printf("%v -> %v\n", tok.Type, tok.Lit)
	}

	// test.TestNextToken(&testing.T{}, string(content))
}
