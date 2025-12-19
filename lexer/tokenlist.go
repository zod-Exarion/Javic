package lexer

/* INFO: Absolute massive list of all the tokens to be used in the lexer
*
*  WARN: Inefficient due to the use to strings as the constants, byte/int is much more computationally efficient
*
* Constants -> Typing lexer.ILLEGAL essentially just substitutes "ILLEGAL", this is just for convenince: nothing of substance here other than ease of development for the programmer/debugger
*
* Maps -> We create two maps - Keywords/Singletons, in order to segregate them for future benefits.
* Also, by assigning the string value to the same string value (Wrapped as TokenType here) -> we can easily check if the current scanned series of runes is indeed a valid keyword or singleton.
 */

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
	DQUOTE    = `"`
	STRING    = "STRING"

	// DONE
	LET   = "LET"
	PRINT = "PRINT"
	INPUT = "INPUT"
	IF    = "IF"
	THEN  = "THEN"
	ELSE  = "ELSE"
	END   = "END"
	FOR   = "FOR"
	TO    = "TO"
	NEXT  = "NEXT"
	STEP  = "STEP"
	CLS   = "CLS"
	WHILE = "WHILE"
	WEND  = "WEND"
	REM   = "REM"

	// WILL WE EVER GET TO THESE I DONT KNOW RAHH
	MOD    = "MOD"
	NOT    = "NOT"
	DIM    = "DIM"
	SELECT = "SELECT"
	CASE   = "CASE"
	IS     = "IS"

	// I dont really know if I want to
	FUNCTION = "FUNCTION"
	RETURN   = "RETURN"

	// GOTO   = "GOTO"
	// GOSUB  = "GOSUB"
)

var keywords = map[string]TokenType{
	"CLS":      CLS,
	"LET":      LET,
	"PRINT":    PRINT,
	"INPUT":    INPUT,
	"IF":       IF,
	"THEN":     THEN,
	"ELSE":     ELSE,
	"END":      END,
	"FOR":      FOR,
	"TO":       TO,
	"NEXT":     NEXT,
	"NOT":      NOT,
	"MOD":      MOD,
	"STEP":     STEP,
	"RETURN":   RETURN,
	"WHILE":    WHILE,
	"WEND":     WEND,
	"DIM":      DIM,
	"REM":      REM,
	"FUNCTION": FUNCTION,

	// "GOTO":  GOTO,
	// "GOSUB": GOSUB,

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
