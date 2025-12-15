package lexer

/* INFO: Token: Defining a token and simple helper functions
*
* Token: First of all, we create a wrapper TokenType just for better readability, has no effect on the meaning of the code
* Then, we have a Type field and a Lit field to record both the type of the token detected and the literal value [useful for identifiers]
*
* Functions: Simple helper functions which return the TokenType of the series of runes passed
 */

import "fmt"

type TokenType string

type Token struct {
	Type TokenType
	Lit  string
}

// If series of runes/characters is a valid keyword, fetch the value from the map. Otherwise it must be an identifier
func checkKeyword(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

// If series of runes/characters is a valid singleton operator, fetch the value from the map. Otherwise it's invalid char.
func checkSingleton(ident string) TokenType {
	if tok, ok := singleton[ident]; ok {
		return tok
	}
	return ILLEGAL
}

// DisplayTokens public function to print a slice(list) of tokens
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
