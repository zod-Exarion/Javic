package tokenizer

// Token type
type TokenType int

type Token struct {
	Type TokenType
	Lit  string
}

// LookupIdent checks if the string is a keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// Searches for the Token Type of the Passed String
func SearchToken(ch string) int {
	for i, j := range tokens {
		if ch == j {
			return i
		}
	}

	return 0
}
