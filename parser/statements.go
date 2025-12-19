package parser

import (
	"log"

	"github.com/zod-Exarion/javic/lexer"
)

type StatementType int // for the ease of the developer, wrapper of int

// HACK: My small brain too stupid for interfaces, just have a struct where only the proper field is non-nil
type Statement struct {
	kind StatementType

	remStatement   *RemStatement
	letStatement   *LetStatement
	printStatement *PrintStatement
	inputStatement *InputStatement
	ifStatement    *IfStatement
	forStatement   *ForStatement
	clsStatement   *ClsStatement
	endStatement   *EndStatement
	whileStatement *WhileStatement
}

// INFO: Here we define each and every statement
type LetStatement struct {
	name  string
	value Expression
}

type PrintStatement struct {
	value []Expression
}

type InputStatement struct {
	prompt string
	name   string
}

type IfStatement struct {
	condition   Expression
	consequence []Statement
	antecendent []Statement
}

type ForStatement struct {
	name  string
	init  Expression
	final Expression
	step  Expression
	body  []Statement
}

type WhileStatement struct {
	condition Expression
	body      []Statement
}

type (
	ClsStatement struct{}
	EndStatement struct{}
)

type RemStatement struct {
	body string
}

// INFO: Here we create functions to parse each and every statment

func (p *Parser) parseRemStatement() *RemStatement {
	var body string
	for p.peek.Type != lexer.NLINE {
		p.next()
		body += string(p.cur.Lit) + " " // Stringify the token until we hit newline
	}

	return &RemStatement{body: body}
}

func (p *Parser) parseLetStatement() *LetStatement {
	if p.cur.Type != lexer.IDENT {
		p.expectToken(lexer.IDENT) // expect identifier
	}

	name := p.cur.Lit

	p.expectToken(lexer.ASSIGN) // expect =

	p.next() // start of expression

	value := p.parseExpression(0)

	return &LetStatement{name: name, value: value}
}

func (p *Parser) parsePrintStatement() *PrintStatement {
	stmt := &PrintStatement{}
	p.next() // current token was 'PRINT', skip over that to reach the expression

	stmt.value = append(stmt.value, p.parseExpression(0))

	for p.peek.Type == lexer.COMMA || p.peek.Type == lexer.SEMICOLON {
		p.next()
		p.next()
		stmt.value = append(stmt.value, p.parseExpression(0))
	}

	return stmt
}

func (p *Parser) parseInputStatement() *InputStatement {
	p.expectToken(lexer.STRING) // expect string

	prompt := p.cur.Lit

	if p.peek.Type == lexer.COMMA || p.peek.Type == lexer.SEMICOLON {
		p.next()
	} else {
		log.Fatalf("expected comma or semicolon, got %v", p.cur.Type)
	}

	p.expectToken(lexer.IDENT) // expect identifier

	name := p.cur.Lit
	// p.next()

	return &InputStatement{prompt: prompt, name: name}
}

func (p *Parser) parseIfStatement() *IfStatement {
	stmt := &IfStatement{}
	p.next() // Current was IF

	stmt.condition = p.parseExpression(0)

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
			stmt.body = append(stmt.body, p.ParseStatement())
		}
		p.next()
	}

	// p.next() // consume NEXT

	if p.cur.Type == lexer.IDENT {
		p.next()
	}

	// p.skipNewlines()

	return stmt
}

func (p *Parser) parseClsStatement() *ClsStatement {
	p.next()

	return &ClsStatement{}
}

func (p *Parser) parseEndStatement() *EndStatement {
	p.next()

	return &EndStatement{}
}

func (p *Parser) parseWhileStatement() *WhileStatement {
	stmt := &WhileStatement{}
	/*WHILE i<=n
		s=s+i
		i=i+1
	WEND*/

	p.next() // Current was WHILE

	stmt.condition = p.parseExpression(0)
	p.next() // skip past expression

	p.skipNewlines()

	for p.cur.Type != lexer.WEND && p.cur.Type != lexer.EOF {
		if p.cur.Type != lexer.NLINE {
			stmt.body = append(stmt.body, p.ParseStatement())
		}
		p.next()
	}

	p.next() // consume WEND

	return stmt
}
