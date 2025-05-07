package test

import (
	"javic/core/lexer"
	"javic/core/tokenizer"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(),;`
	tests := []struct {
		expectedType    tokenizer.TokenType
		expectedLiteral string
	}{
		{tokenizer.ASSIGN, "="},
		{tokenizer.PLUS, "+"},
		{tokenizer.LPAREN, "("},
		{tokenizer.RPAREN, ")"},
		{tokenizer.COMMA, ","},
		{tokenizer.SEMICOLON, ";"},
		{tokenizer.EOF, ""},
	}

	l := lexer.NewLexer(input)
	for i, tt := range tests {
		tok := l.GetToken()
		if tok.Type != tt.expectedType {
			t.Logf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Lit != tt.expectedLiteral {
			t.Logf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Lit)
		}
	}
}
