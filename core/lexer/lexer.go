package lexer

import (
	"javic/core/tokenizer"
	"strings"
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
	lex := &Lexer{input: strings.ToUpper(input)}
	lex.readNext()
	return lex
}

func (lex *Lexer) readNext() {
	if lex.readPos >= len(lex.input) {
		// INFO: Make sure to check both if ch == 0 and width == 0 for EOF
		lex.ch = 0 // Indicates lack of runes to read OR End of File
		lex.runeWidth = 0
		return
	}
	// Unicdoe Support as opposed to ASCII
	lex.ch, lex.runeWidth = utf8.DecodeRuneInString(lex.input[lex.readPos:])

	lex.pos = lex.readPos
	lex.readPos += lex.runeWidth
}

func (lex *Lexer) GetToken() tokenizer.Token {
	var tok tokenizer.Token

	lex.skipWhitespace() // Ensure this does NOT skip newline characters

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
	case '\n':
		tok.Type = tokenizer.NLINE
		tok.Lit = ""

	case 0:
		tok.Lit = ""
		tok.Type = tokenizer.EOF
	default:
		if isLetter(lex.ch) {
			tok.Lit = lex.readIdentifier()
			tok.Type = tokenizer.CheckKeyword(tok.Lit)
			return tok
		}
		if isDigit(lex.ch) {
			tok.Lit = lex.readNumber()
			tok.Type = tokenizer.NUMBER
			return tok
		} else {
			tok.Lit = ""
			tok.Type = tokenizer.ILLEGAL
		}
	}

	lex.readNext()
	return tok
}

func newToken(typ tokenizer.TokenType, lit rune) tokenizer.Token {
	return tokenizer.Token{Type: typ, Lit: string(lit)}
}

func (lex *Lexer) skipWhitespace() {
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\r' {
		lex.readNext()
	}
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func (lex *Lexer) readIdentifier() string {
	initPos := lex.pos

	for isLetter(lex.ch) {
		lex.readNext()
	}

	return lex.input[initPos:lex.pos]
}

func (lex *Lexer) readNumber() string {
	initPos := lex.pos

	for isDigit(lex.ch) {
		lex.readNext()
	}

	return lex.input[initPos:lex.pos]
}
