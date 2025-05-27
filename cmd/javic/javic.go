package javic

import (
	"fmt"
	"javic/qbasic/lexer"
	"javic/qbasic/tokenizer"
	"os"
	"strings"
)

type Transpiler struct {
	FileName string
	Content  string
	Tokens   []tokenizer.Token
	Lexer    *lexer.Lexer
}

func Javic(fileName string) {
	tp := NewTranspiler(fileName)
	tp.Tokenize()

	tokenizer.DisplayTokens(tp.Tokens)
}

func NewTranspiler(fileName string) Transpiler {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	tp := Transpiler{
		Content:  (strings.ToUpper(string(content))),
		FileName: fileName,
	}
	tp.InitLex()
	return tp
}

func (t *Transpiler) InitLex() *Transpiler {
	t.Lexer = lexer.NewLexer(t.Content)
	return t
}

func (tp *Transpiler) Tokenize() *Transpiler {
	for range tp.Content {
		tp.Tokens = append(tp.Tokens, tp.Lexer.GetToken())
	}
	return tp
}
