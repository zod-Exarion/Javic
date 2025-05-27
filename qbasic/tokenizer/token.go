package tokenizer

import "fmt"

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

func CheckSingleton(ident string) TokenType {
	if tok, ok := singleton[ident]; ok {
		return tok
	}
	return ILLEGAL
}

func DisplayTokens(toks []Token) {
	for _, tok := range toks {
		if tok.Type == EOF {
			break
		} else if tok.Type == NLINE {
			fmt.Println()
		} else {
			fmt.Printf("[%v -> %v] ", tok.Type, tok.Lit)
		}
	}
}
