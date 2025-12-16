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
	"fmt"
	"log"

	"github.com/zod-Exarion/javic/lexer"
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
		stmt := p.parseStatement()
		if stmt.kind != 0 { // If parseStatment() returns empty statement(unrecognized) -> dont add it to the list
			statements = append(statements, stmt)
		}
		p.next()

	}

	return statements
}

func (p *Parser) next() {
	p.cur = p.peek
	if p.pos < len(p.tokens) {
		p.peek = p.tokens[p.pos]
	}
	p.pos++
}

func (p *Parser) parseStatement() Statement {
	switch p.cur.Type {
	case lexer.LET:
		return Statement{kind: Let, letStatement: p.parseLetStatement()}
	case lexer.PRINT:
		return Statement{kind: Print, printStatement: p.parsePrintStatement()}
	default:
		return Statement{} // unrecognizable statmenet, doesnt get added in the final list of statemnets
	}
}

func DisplayStatements(statements []Statement) {
	for i := range statements {
		fmt.Printf("%v\n", statements[i])
	}
}

func (p *Parser) expectToken(t lexer.TokenType) bool {
	if p.peek.Type == t {
		p.next()
		return true
	} else {
		log.Fatalf("Expected token %v, got %v", t, p.peek.Type)
		return false
	}
}
