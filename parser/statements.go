package parser

import "github.com/zod-Exarion/javic/lexer"

type StatementType int

const (
	IllegalStatement StatementType = iota
	Let
	Print
	If
	For
)

type Statement struct {
	kind StatementType

	letStatement   *LetStatement
	printStatement *PrintStatement
	ifStatement    *IfStatement
	forStatement   *ForStatement
}

type LetStatement struct {
	name  string
	value Expression
}

type PrintStatement struct {
	value Expression
}

type IfStatement struct{}

type ForStatement struct{}

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

func (s Statement) String() string {
	switch s.kind {
	case Let:
		return "LET " + s.letStatement.name + " = " + s.letStatement.value.String()

	case Print:
		return "PRINT " + s.printStatement.value.String()

	default:
		return "<unknown stmt>"
	}
}
