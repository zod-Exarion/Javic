package tokenizer

const (
	// Special tokens
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"  // variable names, e.g. x, total
	NUMBER = "NUMBER" // numeric literals

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NEQ      = "<>"

	// Delimiters
	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"

	// Keywords (common in QBASIC)
	LET    = "LET"
	PRINT  = "PRINT"
	INPUT  = "INPUT"
	IF     = "IF"
	THEN   = "THEN"
	ELSE   = "ELSE"
	END    = "END"
	FOR    = "FOR"
	TO     = "TO"
	NEXT   = "NEXT"
	GOTO   = "GOTO"
	GOSUB  = "GOSUB"
	RETURN = "RETURN"
	WHILE  = "WHILE"
	WEND   = "WEND"
	DIM    = "DIM"
	REM    = "REM" // comment
)
