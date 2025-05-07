package lexer

import (
	"javic/core/tokenizer"
	"unicode/utf8"
)

type Lexer struct {
	input     string
	pos       int
	readPos   int
	ch        rune
	runeWidth int
}

func NewLexer(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readNext()
	return lex
}

func (lex *Lexer) readNext() {
	if lex.readPos >= len(lex.input) {
		// INFO: Make sure to check both if ch == 0 and width == 0 for EOF
		lex.ch = 0 // Indicates lack of runes to read OR End of File
		lex.runeWidth = 0
	} else {
		// Unicdoe Support as opposed to ASCII
		lex.ch, lex.runeWidth = utf8.DecodeRuneInString(lex.input[lex.readPos:])
	}

	lex.pos = lex.readPos
	lex.readPos += lex.runeWidth
}

func (lex *Lexer) GetToken() tokenizer.Token {
	var tok tokenizer.Token

	switch lex.ch {
	case '=':
		tok = newToken(tokenizer.ASSIGN, lex.ch)
	case ';':
		tok = newToken(tokenizer.SEMICOLON, lex.ch)
	case '(':
		tok = newToken(tokenizer.LPAREN, lex.ch)
	case ')':
		tok = newToken(tokenizer.RPAREN, lex.ch)
	case ',':
		tok = newToken(tokenizer.COMMA, lex.ch)
	case '+':
		tok = newToken(tokenizer.PLUS, lex.ch)
	case 0:
		tok.Lit = ""
		tok.Type = tokenizer.EOF
	}

	lex.readNext()
	return tok
}

func newToken(typ tokenizer.TokenType, lit rune) tokenizer.Token {
	return tokenizer.Token{Type: typ, Lit: string(lit)}
}
