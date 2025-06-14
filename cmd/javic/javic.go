package javic

import (
	"fmt"
	"javic/qbasic/lexer"
	"javic/qbasic/parser"
	"javic/qbasic/tokenizer"
	"os"
)

type Transpiler struct {
	FileName string
	Content  string
	Tokens   []tokenizer.Token
	Lexer    *lexer.Lexer
	Parser   *parser.Parser
}

func Javic(fileName string) {
	tp := NewTranspiler(fileName, true)
	tp.Parser.ParseProgram()
}

func NewTranspiler(fileName string, advparser bool) Transpiler {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	tp := Transpiler{
		Content:  string(content),
		FileName: fileName,
	}

	tp.initTranspiler(advparser)
	return tp
}

func (t *Transpiler) initTranspiler(flag bool) {
	t.initLex()
	t.initParse(flag)
}

func (t *Transpiler) initLex() *Transpiler {
	t.Lexer = lexer.NewLexer(t.Content)
	return t
}

func (t *Transpiler) initParse(flag bool) *Transpiler {
	t.Parser = parser.NewParser(t.Lexer, flag)
	return t
}

func (tp *Transpiler) tokenize() *Transpiler {
	for range tp.Content {
		tp.Tokens = append(tp.Tokens, tp.Lexer.GetToken())
	}
	return tp
}
