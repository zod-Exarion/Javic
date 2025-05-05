package tokenizer

// Token names
const (
	ILLEGAL TokenType = iota
	EOF

	// Identifiers + literals
	IDENT  // main
	NUMBER // 12345

	// Operators
	ASSIGN   // =
	PLUS     // +
	MINUS    // -
	BANG     // !
	ASTERISK // *
	SLASH    // /

	LT     // <
	GT     // >
	EQ     // ==
	NOT_EQ // !=

	// Delimiters
	COMMA     // ,
	SEMICOLON // ;

	LPAREN   // (
	RPAREN   // )
	LBRACE   // {
	RBRACE   // }
	LBRACKET // [
	RBRACKET // ]

	// Keywords
	FUNCTION
	SUB
	END
	IF
	THEN
	ELSE
	ELSEIF
	FOR
	TO
	STEP
	NEXT
	WHILE
	WEND
	DO
	LOOP
	UNTIL
	EXIT
	SELECT
	CASE
	IS
	TYPE
	AS
	DIM
	SHARED
	REDIM
	CONST
	GOTO
	GOSUB
	RETURN
	PRINT
	INPUT
	LET

	REM      // Comment
	STRING   // "hello world"
	SINGLE   // Single-precision floating-point number
	DOUBLE   // Double-precision floating-point number
	INTEGER  // Integer number
	LONG     // Long integer number
	CURRENCY // Currency number

	AND // and
	OR  // or
	XOR // xor
	MOD // mod
	NOT // not

	// Data types
	INTEGER_TYPE
	LONG_TYPE
	SINGLE_TYPE
	DOUBLE_TYPE
	STRING_TYPE
	CURRENCY_TYPE
	BOOLEAN_TYPE
	DATE_TYPE

	// QB keywords that are also commands
	CLS
	COLOR
	LOCATE

	// QB system variables/functions
	TIMER
)

var tokens = []string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	IDENT:  "IDENT",
	NUMBER: "NUMBER",

	ASSIGN:   "=",
	PLUS:     "+",
	MINUS:    "-",
	BANG:     "!",
	ASTERISK: "*",
	SLASH:    "/",

	LT:     "<",
	GT:     ">",
	EQ:     "==",
	NOT_EQ: "!=",

	COMMA:     ",",
	SEMICOLON: ";",

	LPAREN:   "(",
	RPAREN:   ")",
	LBRACE:   "{",
	RBRACE:   "}",
	LBRACKET: "[",
	RBRACKET: "]",

	FUNCTION: "FUNCTION",
	SUB:      "SUB",
	END:      "END",
	IF:       "IF",
	THEN:     "THEN",
	ELSE:     "ELSE",
	ELSEIF:   "ELSEIF",
	FOR:      "FOR",
	TO:       "TO",
	STEP:     "STEP",
	NEXT:     "NEXT",
	WHILE:    "WHILE",
	WEND:     "WEND",
	DO:       "DO",
	LOOP:     "LOOP",
	UNTIL:    "UNTIL",
	EXIT:     "EXIT",
	SELECT:   "SELECT",
	CASE:     "CASE",
	IS:       "IS",
	TYPE:     "TYPE",
	AS:       "AS",
	DIM:      "DIM",
	SHARED:   "SHARED",
	REDIM:    "REDIM",
	CONST:    "CONST",
	GOTO:     "GOTO",
	GOSUB:    "GOSUB",
	RETURN:   "RETURN",
	PRINT:    "PRINT",
	INPUT:    "INPUT",
	LET:      "LET",

	REM:      "REM",
	STRING:   "STRING",
	SINGLE:   "SINGLE",
	DOUBLE:   "DOUBLE",
	INTEGER:  "INTEGER",
	LONG:     "LONG",
	CURRENCY: "CURRENCY",

	AND: "AND",
	OR:  "OR",
	XOR: "XOR",
	MOD: "MOD",
	NOT: "NOT",

	INTEGER_TYPE:  "INTEGER_TYPE",
	LONG_TYPE:     "LONG_TYPE",
	SINGLE_TYPE:   "SINGLE_TYPE",
	DOUBLE_TYPE:   "DOUBLE_TYPE",
	STRING_TYPE:   "STRING_TYPE",
	CURRENCY_TYPE: "CURRENCY_TYPE",
	BOOLEAN_TYPE:  "BOOLEAN_TYPE",
	DATE_TYPE:     "DATE_TYPE",

	CLS:    "CLS",
	COLOR:  "COLOR",
	LOCATE: "LOCATE",

	TIMER: "TIMER",
}

func (t TokenType) String() string {
	return tokens[t]
}

var keywords = map[string]TokenType{
	"function": FUNCTION,
	"sub":      SUB,
	"end":      END,
	"if":       IF,
	"then":     THEN,
	"else":     ELSE,
	"elseif":   ELSEIF,
	"for":      FOR,
	"to":       TO,
	"step":     STEP,
	"next":     NEXT,
	"while":    WHILE,
	"wend":     WEND,
	"do":       DO,
	"loop":     LOOP,
	"until":    UNTIL,
	"exit":     EXIT,
	"select":   SELECT,
	"case":     CASE,
	"is":       IS,
	"type":     TYPE,
	"as":       AS,
	"dim":      DIM,
	"shared":   SHARED,
	"redim":    REDIM,
	"const":    CONST,
	"goto":     GOTO,
	"gosub":    GOSUB,
	"return":   RETURN,
	"print":    PRINT,
	"input":    INPUT,
	"let":      LET,
	"rem":      REM,

	"integer":  INTEGER_TYPE,
	"long":     LONG_TYPE,
	"single":   SINGLE_TYPE,
	"double":   DOUBLE_TYPE,
	"string":   STRING_TYPE,
	"currency": CURRENCY_TYPE,

	"and": AND,
	"or":  OR,
	"xor": XOR,
	"mod": MOD,
	"not": NOT,

	"cls":    CLS,
	"color":  COLOR,
	"locate": LOCATE,

	"timer": TIMER,
}
