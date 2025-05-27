package lexer

import (
	"javic/qbasic/tokenizer"
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

func (lex *Lexer) GetToken() tokenizer.Token {
	var tok tokenizer.Token

	lex.eatWhitespace() // Ensure this does NOT skip newline characters

	switch lex.ch {
	case '\n':
		tok.Type = tokenizer.NLINE
		tok.Lit = ""
	case 0:
		tok.Lit = ""
		tok.Type = tokenizer.EOF
	case '"':
		tok.Type = tokenizer.STRING
		tok.Lit = lex.readString('"')

	default:
		if isLetter(lex.ch) {
			tok.Lit = lex.readIdentifier()
			tok.Type = tokenizer.CheckKeyword(tok.Lit)
			return tok
		} else if isDigit(lex.ch) {
			tok.Lit = lex.readNumber()
			tok.Type = tokenizer.NUMBER
			return tok
		} else {
			tok = newToken(tokenizer.CheckSingleton(string(lex.ch)), lex.ch)
		}
	}

	lex.readNext()
	return tok
}

func newToken(typ tokenizer.TokenType, lit rune) tokenizer.Token {
	return tokenizer.Token{Type: typ, Lit: string(lit)}
}

func (lex *Lexer) eatWhitespace() {
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\r' {
		lex.readNext()
	}
}
