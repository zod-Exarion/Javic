package parser

import (
	"fmt"
	"log"

	"github.com/zod-Exarion/javic/lexer"
)

type StatementType int // for the ease of the developer, wrapper of int

// HACK: My small brain too stupid for interfaces, just have a struct where only the proper field is non-nil
type Statement struct {
	kind StatementType

	letStatement   *LetStatement
	printStatement *PrintStatement
	inputStatement *InputStatement
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

type InputStatement struct {
	prompt string
	name   string
}

type IfStatement struct {
	lhscondition Expression
	rhscondition Expression
	comparison   string
	consequence  []Statement
	antecendent  []Statement
}

type ForStatement struct {
	name  string
	init  Expression
	final Expression
	step  Expression
	body  []Statement
}

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

func (p *Parser) parseInputStatement() *InputStatement {
	p.expectToken(lexer.STRING) // expect string

	prompt := p.cur.Lit

	p.expectToken(lexer.COMMA) // except a ,

	p.expectToken(lexer.IDENT) // expect identifier

	name := p.cur.Lit
	p.next()

	return &InputStatement{prompt: prompt, name: name}
}

func (p *Parser) parseIfStatement() *IfStatement {
	stmt := &IfStatement{}
	p.next() // Current was IF
	stmt.lhscondition = p.parseExpression(0)

	p.next()
	stmt.comparison = p.parseComparisonOperator()

	p.next()
	stmt.rhscondition = p.parseExpression(0)

	p.expectToken(lexer.THEN) // expect THEN
	p.next()

	// INFO: Then block
	for p.cur.Type != lexer.ELSE &&
		!(p.cur.Type == lexer.END && p.peek.Type == lexer.IF) {

		if p.cur.Type != lexer.NLINE {
			stmt.consequence = append(stmt.consequence, p.ParseStatement())
		}
		p.next()
	}

	// INFO: Else block
	if p.cur.Type == lexer.ELSE {
		p.next()

		for !(p.cur.Type == lexer.END && p.peek.Type == lexer.IF) {

			if p.cur.Type != lexer.NLINE {
				stmt.antecendent = append(stmt.antecendent, p.ParseStatement())
			}
			p.next()
		}
	}

	p.next()
	p.next() // consume END IF

	return stmt
}

func (p *Parser) parseComparisonOperator() string {
	if p.cur.Type != lexer.LT &&
		p.cur.Type != lexer.GT &&
		p.cur.Type != lexer.ASSIGN {
		log.Fatalf("expected comparison operator, got %v", p.cur.Type)
	}

	op := p.cur.Lit

	// Lookahead for <= >= <>
	if p.peek.Type == lexer.ASSIGN || p.peek.Type == lexer.GT {
		p.next()
		op += p.cur.Lit
	}

	switch op {
	case "=", "<", ">", "<=", ">=", "<>":
		return op
	default:
		log.Fatalf("invalid QBASIC comparison operator: %s", op)
	}

	return ""
}

func (p *Parser) parseForStatement() *ForStatement {
	/*
		*FOR I = 1.5 TO 5.6 STEP 1
			PRINT I
		 NEXT
	*/
	stmt := &ForStatement{}
	p.expectToken(lexer.IDENT) // expect identifier
	stmt.name = p.cur.Lit
	p.expectToken(lexer.ASSIGN) // expect =
	p.next()                    // start of expression
	stmt.init = p.parseExpression(0)
	p.expectToken(lexer.TO) // expect TO
	p.next()                // start of expression
	stmt.final = p.parseExpression(0)
	p.next() // end the expression

	if p.cur.Type == lexer.STEP {
		p.next()
		stmt.step = p.parseExpression(0)
		p.next() // cosume number
	} else {
		stmt.step = Expression{kind: Number, numberExpression: 1}
	}

	p.skipNewlines()

	for p.cur.Type != lexer.NEXT {
		if p.cur.Type != lexer.NLINE {
			fmt.Printf("Entered loop: got [%v -> %v]\n", p.cur.Type, p.cur.Lit)
			stmt.body = append(stmt.body, p.ParseStatement())
		}
		p.next()
	}

	p.next() // consume NEXT

	if p.cur.Type == lexer.IDENT {
		p.next()
	}

	p.skipNewlines()

	return stmt
}
