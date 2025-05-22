package test

import (
	"javic/cmd/javic"
	"javic/core/lexer"
	"javic/core/tokenizer"
	"testing"
)

func TestNextToken(t *testing.T) {
	tp := javic.NewTranspiler("qb/dummy.bas")
	input := tp.Content

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
		{tokenizer.NUMBER, "8"},
		{tokenizer.NLINE, ""},
		{tokenizer.PRINT, "PRINT"},
		{tokenizer.IDENT, "Y"},
		{tokenizer.COMMA, ","},
		{tokenizer.IDENT, "X"},
		{tokenizer.COMMA, ","},
		{tokenizer.NUMBER, "0"},
	}

	l := lexer.NewLexer(input)
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
