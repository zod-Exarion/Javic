package javic

import (
	"fmt"
	"javic/core/tokenizer"
	"os"
	"strings"
)

type Transpiler struct {
	FileName string
	Content  string
	Tokens   []tokenizer.Token
}

func Javic(fileName string) {
	// lex := lexer.NewLexer(string(content))
	// for range content {
	// 	tok := lex.GetToken()
	// 	if tok.Type == tokenizer.EOF {
	// 		break
	// 	}
	// 	fmt.Printf("%v -> %v\n", tok.Type, tok.Lit)
	// }
}

func NewTranspiler(fileName string) Transpiler {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	return Transpiler{
		Content:  (strings.ToUpper(string(content))),
		FileName: fileName,
	}
}
