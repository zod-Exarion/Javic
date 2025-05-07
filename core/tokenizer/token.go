package tokenizer

// Token type
type TokenType string

type Token struct {
	Type TokenType
	Lit  string
}
