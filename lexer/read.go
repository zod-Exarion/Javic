package lexer

/* INFO: Reading runes/characters from the input string
*
* Core text reader functions which make the lexing possible
* Most functions are easy to decipher upon reading hence further explanations are not required
 */

import "unicode/utf8"

// PERF: Instead of simple ASCII, we take the character and its runewidth to introduce unicode support
// Reads the current character and stores it in lex.ch [current character field]
func (lex *Lexer) readNext() {
	if lex.readPos >= len(lex.input) {
		// WARN: Make sure to check both if ch == 0 and width == 0 for EOF

		lex.ch = 0 // Indicates lack of runes to read OR End of File
		lex.runeWidth = 0
		return
	}
	lex.ch, lex.runeWidth = utf8.DecodeRuneInString(lex.input[lex.readPos:])

	lex.pos = lex.readPos
	lex.readPos += lex.runeWidth // Since unicode chars can span through multiple indices, we have to account for the width
}

// checks out the next character [readPos is always one step ahead of the current read character]
func (lex *Lexer) peekNext() byte {
	if lex.readPos >= len(lex.input) {
		return 0
	} else {
		return lex.input[lex.readPos]
	}
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '!' && ch <= ')')
}

func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9') || ch == '.'
}

func (lex *Lexer) readIdentifier() string {
	initPos := lex.pos

	for isLetter(lex.ch) || isDigit(lex.ch) {
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

	return lex.input[initPos : lex.pos+1] // +1 because we need to include the ending quote "
}
