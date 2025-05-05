package lexer

import "javic/core/tokenizer"

func (lex *Lexer) getToken() tokenizer.Token {
	var tok tokenizer.Token

	tok = newToken(tokenizer.TokenType(tokenizer.SearchToken(string(lex.ch))), lex.ch)

	return tok
}
