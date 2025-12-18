// Package parser is the Pratt-parser for Javic
//
// the Parser has four fields:
//
// tokens -> the list of tokens to be parsed
//
// pos -> the current position in the token list
//
// cur -> the current token being parsed
//
// peek -> the next token being parsed
//
// Usage:
// call Parse() and it does all the magic for you
package parser

import (
	"log"

	"github.com/zod-Exarion/javic/lexer"
)

// again, constants for the ease of the developer
const (
	IllegalStatement StatementType = iota
	Let
	Print
	Input
	If
	For
	Cls
	End
)

type Parser struct {
	tokens []lexer.Token
	pos    int

	cur  lexer.Token
	peek lexer.Token
}

func Parse(tokens []lexer.Token) []Statement {
	p := &Parser{tokens: tokens}
	// INFO: Impetus to initiate default states into the tokens
	p.next() // current token = nothing and peek = tokens[0]
	p.next() // current token = tokens[0] and peek = tokens[1]

	statements := make([]Statement, 0)

	// INFO: Loop though all the tokens and append it to the list of statements
	for p.cur.Type != lexer.EOF {
		stmt := p.ParseStatement()
		if stmt.kind != 0 { // If parseStatment() returns empty statement(unrecognized) -> dont add it to the list
			statements = append(statements, stmt)
		}
		p.next()

	}

	return statements
}

// 1. Sets current token to the previously peeked token
// 2. If safe, assigns peek to the next token [since pos is always one step ahead]
// 3. pos: 0 when peek = nil, pos:1 when peek = token[0]
func (p *Parser) next() {
	p.cur = p.peek
	if p.pos < len(p.tokens) {
		p.peek = p.tokens[p.pos]
	}
	p.pos++
}

func (p *Parser) ParseStatement() Statement {
	// p.skipNewlines()
	switch p.cur.Type {
	case lexer.LET:
		return Statement{kind: Let, letStatement: p.parseLetStatement()}
	case lexer.PRINT:
		return Statement{kind: Print, printStatement: p.parsePrintStatement()}
	case lexer.INPUT:
		return Statement{kind: Input, inputStatement: p.parseInputStatement()}
	case lexer.IF:
		return Statement{kind: If, ifStatement: p.parseIfStatement()}
	case lexer.FOR:
		return Statement{kind: For, forStatement: p.parseForStatement()}
	case lexer.CLS:
		return Statement{kind: Cls, clsStatement: p.parseClsStatement()}
	case lexer.END:
		return Statement{kind: End, endStatement: p.parseEndStatement()}
	default:
		return Statement{} // unrecognizable statmenet, doesnt get added in the final list of statemnets
	}
}

// Checks if the next token is the proper type expected
// The param is specific to each statement, e.g: LET expects an identifier and a =
func (p *Parser) expectToken(t lexer.TokenType) bool {
	if p.peek.Type == t {
		p.next()
		return true
	} else {
		log.Fatalf("Expected token %v, got %v", t, p.peek.Type)
		return false
	}
}
