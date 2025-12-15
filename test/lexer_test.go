package test

import (
	"javic/cmd/javic"
	"javic/qbasic/tokenizer"
	"testing"
)

func TestNextToken(t *testing.T) {
	tp := javic.NewTranspiler("qb/parse1.bas", false)

	tests := []struct {
		expectedType    tokenizer.TokenType
		expectedLiteral string
	}{
		{tokenizer.LET, "LET"},
		{tokenizer.IDENT, "X"},
		{tokenizer.ASSIGN, "="},
		{tokenizer.NUMBER, "5"},
		{tokenizer.NLINE, ""},
		{tokenizer.LET, "LET"},
		{tokenizer.IDENT, "Y"},
		{tokenizer.ASSIGN, "="},
		{tokenizer.NUMBER, "10"},
		{tokenizer.NLINE, ""},
		{tokenizer.LET, "LET"},
		{tokenizer.IDENT, "FOOBAR"},
		{tokenizer.ASSIGN, "="},
		{tokenizer.NUMBER, "838383"},
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
