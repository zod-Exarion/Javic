package parser

import "github.com/zod-Exarion/javic/lexer"

type StatementType int // for the ease of the developer, wrapper of int

// again, constants for the ease of the developer
const (
	IllegalStatement StatementType = iota
	Let
	Print
	If
	For
)

// HACK: My small brain too stupid for interfaces, just have a struct where only the proper field is non-nil
type Statement struct {
	kind StatementType

	letStatement   *LetStatement
	printStatement *PrintStatement
	ifStatement    *IfStatement
	forStatement   *ForStatement
}

// INFO: Here we define each and every statement
type LetStatement struct {
	name  string
	value Expression
}

type PrintStatement struct {
	value Expression
}

type IfStatement struct{}

type ForStatement struct{}

// INFO: Here we create functions to parse each and every statment
func (p *Parser) parseLetStatement() *LetStatement {
	p.expectToken(lexer.IDENT) // expect identifier

	name := p.cur.Lit

	p.expectToken(lexer.ASSIGN) // expect =

	p.next() // start of expression

	value := p.parseExpression(0)

	return &LetStatement{name: name, value: value}
}

func (p *Parser) parsePrintStatement() *PrintStatement {
	p.next() // current token was 'PRINT', skip over that to reach the expression

	value := p.parseExpression(0)

	return &PrintStatement{value: value}
}
