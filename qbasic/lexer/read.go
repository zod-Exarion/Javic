package lexer

import "unicode/utf8"

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

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9') || ch == '.'
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

func (lex *Lexer) readString(delim rune) string {
	initPos := lex.pos

	lex.readNext()
	for lex.ch != delim {
		lex.readNext()
	}

	return lex.input[initPos : lex.pos+1]
}
