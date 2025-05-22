package tokenizer

// Token type
type TokenType string

type Token struct {
	Type TokenType
	Lit  string
}

func CheckKeyword(ident string) TokenType {
	if tok, ok := keuywords[ident]; ok {
		return tok
	}
	return IDENT
}
