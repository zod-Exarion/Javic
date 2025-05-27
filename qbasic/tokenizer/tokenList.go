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
	CARET    = "^"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	NEQ      = "<>"
	EX       = "!"
	POUND    = "$"
	SQUOTE   = "'"

	// Delimiters
	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	NLINE     = "\\n"

	// String
	DQUOTE = `"`
	STRING = "STRING"

	// Keywords (common in QBASIC)
	MOD    = "MOD"
	NOT    = "NOT"
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
	SELECT = "SELECT"
	CASE   = "CASE"
	IS     = "IS"
)

var keuywords = map[string]TokenType{
	"LET":   LET,
	"PRINT": PRINT,
	"INPUT": INPUT,
	"IF":    IF,
	"THEN":  THEN,
	"ELSE":  ELSE,
	"END":   END,
	"FOR":   FOR,
	"TO":    TO,
	"NEXT":  NEXT,
	"NOT":   NOT,
	"MOD":   MOD,

	// WARN: Dangerous Territory
	"GOTO":  GOTO,
	"GOSUB": GOSUB,

	"RETURN": RETURN,
	"WHILE":  WHILE,
	"WEND":   WEND,
	"DIM":    DIM,
	"REM":    REM,
}

var singleton = map[string]TokenType{
	// Operators
	"=":  ASSIGN,
	"+":  PLUS,
	"-":  MINUS,
	"*":  ASTERISK,
	"^":  CARET,
	"/":  SLASH,
	"<":  LT,
	">":  GT,
	"<>": NEQ,
	"!":  EX,
	"$":  POUND,
	"'":  SQUOTE,

	// String
	`"`: DQUOTE,

	// Delimiters
	",":   COMMA,
	":":   COLON,
	";":   SEMICOLON,
	"(":   LPAREN,
	")":   RPAREN,
	"\\n": NLINE,
}
