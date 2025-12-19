// Package lexer is the lexer for Javic
//
// The Lexer consists of five fields:
//
// input -> This is the content of the text file
//
// pos -> Last position which has been read from the [input]
//
// readPos -> Current position of the reader
//
// ch -> Last character stored from [pos]
//
// runeWidth -> The width of the character stored in [ch] for Unicode support
//
// Usage:
// You call Lex() and it does all the magic for you.
package lexer

import "strings"

type Lexer struct {
	input     string
	pos       int
	readPos   int
	ch        rune
	runeWidth int
}

// Lex -> "main" method of the lexer, takes an input and returns a slice of tokens : []Token
func Lex(content string) []Token {
	lexer := &Lexer{input: content}
	lexer.readNext() // HACK: Initial impetus to set the lex.ch properly for getToken() method

	var tokens []Token
	for range content {
		tokens = append(tokens, lexer.getToken())
	}

	return tokens
}

// getToken() -> uses the reader functions from [read.go] to read the current token at [pos] and return a [Token]
// 1. readNext() already has been called once in Lex()
// 2. switch case on the current character
// 3. go through all possibilities and return a token using reader functions
func (lex *Lexer) getToken() Token {
	var tok Token

	lex.eatWhitespace()

	switch lex.ch {
	case '\n':
		tok.Type = NLINE
		tok.Lit = ""
	case 0:
		tok.Lit = ""
		tok.Type = EOF
	case '"':
		tok.Type = STRING
		tok.Lit = lex.readString('"')
	case '<':
		if lex.peekNext() == '>' {
			tok.Type = NEQ
			tok.Lit = lex.readString('>')
		} else if lex.peekNext() == '=' {
			tok.Type = LTE
			tok.Lit = lex.readString('=')
		} else {
			tok = newToken(LT, lex.ch)
		}
	case '>':
		if lex.peekNext() == '<' {
			tok.Type = NEQ
			tok.Lit = lex.readString('<')
		} else if lex.peekNext() == '=' {
			tok.Type = GTE
			tok.Lit = lex.readString('=')
		} else {
			tok = newToken(GT, lex.ch)
		}

	default:
		if isLetter(lex.ch) {
			tok.Lit = strings.ToUpper(lex.readIdentifier())   // if any sort of letter read that name (string)
			tok.Type = checkKeyword(strings.ToUpper(tok.Lit)) // if name is not keyword -> identifier
			return tok
		} else if isDigit(lex.ch) {
			tok.Type = NUMBER
			tok.Lit = lex.readNumber()
			return tok
		} else {
			tok = newToken(checkSingleton(string(lex.ch)), lex.ch) // final case for operators
			if tok.Type == ASSIGN {
				if lex.peekNext() == '>' {
					tok.Type = GTE
					tok.Lit = lex.readString('>')
				} else {
					if lex.peekNext() == '<' {
						tok.Type = LTE
						tok.Lit = lex.readString('<')
					}
				}
			}
		}
	}
	lex.readNext()
	return tok
}

func newToken(typ TokenType, lit rune) Token {
	return Token{Type: typ, Lit: string(lit)}
}

// WARN: Must not remove newLine characters
func (lex *Lexer) eatWhitespace() {
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\r' {
		lex.readNext()
	}
}
