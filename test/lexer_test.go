package test

import (
	"javic/cmd/javic"
	"javic/qbasic/tokenizer"
	"testing"
)

func TestNextToken(t *testing.T) {
	tp := javic.NewTranspiler("qb/dummy.bas")

	tests := []struct {
		expectedType    tokenizer.TokenType
		expectedLiteral string
	}{
		{tokenizer.IDENT, "X"},
		{tokenizer.ASSIGN, "="},
		{tokenizer.NUMBER, "55"},
		{tokenizer.NLINE, ""},
		{tokenizer.IDENT, "Y"},
		{tokenizer.ASSIGN, "="},
		{tokenizer.NUMBER, "8.2"},
		{tokenizer.NLINE, ""},
		{tokenizer.PRINT, "PRINT"},
		{tokenizer.IDENT, "Y"},
		{tokenizer.COMMA, ","},
		{tokenizer.IDENT, "X"},
		{tokenizer.COMMA, ","},
		{tokenizer.NUMBER, "0"},
		{tokenizer.NLINE, ""},

		{tokenizer.IDENT, "E"},
		{tokenizer.ASSIGN, "="},
		{tokenizer.NUMBER, "1"},
		{tokenizer.NLINE, ""},

		{tokenizer.IF, "IF"},
		{tokenizer.NOT, "NOT"},
		{tokenizer.IDENT, "E"},
		{tokenizer.THEN, "THEN"},
		{tokenizer.NLINE, ""},

		{tokenizer.IF, "IF"},
		{tokenizer.IDENT, "E"},
		{tokenizer.ASSIGN, "="},
		{tokenizer.NUMBER, "0"},
		{tokenizer.THEN, "THEN"},
		{tokenizer.NLINE, ""},
		{tokenizer.PRINT, "PRINT"},
		{tokenizer.MINUS, "-"},
		{tokenizer.NUMBER, "69"},
		{tokenizer.NLINE, ""},
		{tokenizer.ELSE, "ELSE"},
		{tokenizer.NLINE, ""},
		{tokenizer.PRINT, "PRINT"},
		{tokenizer.IDENT, "E"},
		{tokenizer.NLINE, ""},

		{tokenizer.END, "END"},
		{tokenizer.IF, "IF"},
		{tokenizer.NLINE, ""},

		{tokenizer.END, "END"},
		{tokenizer.IF, "IF"},
		{tokenizer.NLINE, ""},
	}

	l := tp.Lexer
	for i, tt := range tests {
		tok := l.GetToken()
		if tok.Type != tt.expectedType {
			t.Errorf("tests[%d] - token type wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Lit != tt.expectedLiteral {
			t.Errorf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Lit)
		}
	}
}
