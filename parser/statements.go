package parser

import (
	"log"

	"github.com/zod-Exarion/javic/lexer"
)

type StatementType int // for the ease of the developer, wrapper of int

// HACK: My small brain too stupid for interfaces, just have a struct where only the proper field is non-nil
type Statement struct {
	Kind StatementType

	RemStatement    *RemStatement
	LetStatement    *LetStatement
	PrintStatement  *PrintStatement
	InputStatement  *InputStatement
	IfStatement     *IfStatement
	ForStatement    *ForStatement
	ClsStatement    *ClsStatement
	EndStatement    *EndStatement
	WhileStatement  *WhileStatement
	DimStatement    *DimStatement
	RedimStatement  *RedimStatement
	SelectStatement *SelectStatement
}

// INFO: Here we define each and every statement
type LetStatement struct {
	Name  string
	Value Expression
}

type PrintStatement struct {
	Value []Expression
}

type InputStatement struct {
	Prompt string
	Name   string
}

type IfStatement struct {
	Condition   Expression
	Consequence []Statement
	Antecendent []Statement
}

type ForStatement struct {
	Name  string
	Init  Expression
	Final Expression
	Step  Expression
	Body  []Statement
}

type WhileStatement struct {
	Condition Expression
	Body      []Statement
}

type (
	ClsStatement struct{}
	EndStatement struct{}
)

type RemStatement struct {
	Body string
}

type DimStatement struct {
	Name     string
	Datatype string
}

type RedimStatement struct {
	Name string
}

type SelectStatement struct {
	Match       Expression
	Cases       []*CasePhrase
	Defaultcase []Statement
}

type CasePhrase struct {
	Value Expression
	Body  []Statement
}

// INFO: Here we create functions to parse each and every statment

func (p *Parser) parseRemStatement() *RemStatement {
	var body string
	for p.peek.Type != lexer.NLINE {
		p.next()
		body += string(p.cur.Lit) + " " // Stringify the token until we hit newline
	}

	return &RemStatement{Body: body}
}

func (p *Parser) parseLetStatement() *LetStatement {
	if p.cur.Type != lexer.IDENT {
		p.expectToken(lexer.IDENT) // expect identifier
	}

	name := p.cur.Lit

	p.expectToken(lexer.ASSIGN) // expect =

	p.next() // start of expression

	value := p.parseExpression(0)

	return &LetStatement{Name: name, Value: value}
}

func (p *Parser) parsePrintStatement() *PrintStatement {
	stmt := &PrintStatement{}
	p.next() // current token was 'PRINT', skip over that to reach the expression

	stmt.Value = append(stmt.Value, p.parseExpression(0))

	for p.peek.Type == lexer.COMMA || p.peek.Type == lexer.SEMICOLON {
		p.next()
		p.next()
		stmt.Value = append(stmt.Value, p.parseExpression(0))
		// p.next()
	}

	return stmt
}

func (p *Parser) parseInputStatement() *InputStatement {
	var prompt string
	if p.peek.Type == lexer.STRING {
		p.next()
		prompt = p.cur.Lit
		if p.peek.Type == lexer.COMMA || p.peek.Type == lexer.SEMICOLON {
			p.next()
		} else {
			log.Fatalf("expected comma or semicolon, got %v", p.cur.Type)
		}
	}

	p.expectToken(lexer.IDENT) // expect identifier

	name := p.cur.Lit
	// p.next()

	return &InputStatement{Prompt: prompt, Name: name}
}

func (p *Parser) parseIfStatement() *IfStatement {
	stmt := &IfStatement{}
	p.next() // Current was IF

	stmt.Condition = p.parseExpression(0)

	p.expectToken(lexer.THEN) // expect THEN
	p.next()

	// INFO: Then block
	for p.cur.Type != lexer.ELSE &&
		!(p.cur.Type == lexer.END && p.peek.Type == lexer.IF) {

		if p.cur.Type != lexer.NLINE {
			stmt.Consequence = append(stmt.Consequence, p.ParseStatement())
		}
		p.next()
	}

	// INFO: Else block
	if p.cur.Type == lexer.ELSE {
		p.next()

		for !(p.cur.Type == lexer.END && p.peek.Type == lexer.IF) {

			if p.cur.Type != lexer.NLINE {
				stmt.Antecendent = append(stmt.Antecendent, p.ParseStatement())
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
	stmt.Name = p.cur.Lit
	p.expectToken(lexer.ASSIGN) // expect =
	p.next()                    // start of expression
	stmt.Init = p.parseExpression(0)
	p.expectToken(lexer.TO) // expect TO
	p.next()                // start of expression
	stmt.Final = p.parseExpression(0)
	p.next() // end the expression

	if p.cur.Type == lexer.STEP {
		p.next()
		stmt.Step = p.parseExpression(0)
		p.next() // cosume number
	} else {
		stmt.Step = Expression{Kind: Number, NumberExpression: 1}
	}

	p.skipNewlines()

	for p.cur.Type != lexer.NEXT {
		if p.cur.Type != lexer.NLINE {
			stmt.Body = append(stmt.Body, p.ParseStatement())
		}
		p.next()
	}

	p.next() // consume NEXT

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

	stmt.Condition = p.parseExpression(0)
	p.next() // skip past expression

	p.skipNewlines()

	for p.cur.Type != lexer.WEND && p.cur.Type != lexer.EOF {
		if p.cur.Type != lexer.NLINE {
			stmt.Body = append(stmt.Body, p.ParseStatement())
		}
		p.next()
	}

	p.next() // consume WEND

	return stmt
}

func (p *Parser) parseDimStatement() *DimStatement {
	stmt := &DimStatement{}

	p.next() // current token was 'DIM', skip over that to reach the var

	stmt.Name = p.cur.Lit

	if p.peek.Type == lexer.AS {
		p.next()
		p.next()
		stmt.Datatype = p.cur.Lit
	}

	return stmt
}

func (p *Parser) parseRedimStatement() *RedimStatement {
	p.next() // current token was 'REDIM', skip over that to reach the var
	return &RedimStatement{Name: p.cur.Lit}
}

func (p *Parser) parseSelectStatement() *SelectStatement {
	stmt := &SelectStatement{}

	p.expectToken(lexer.CASE)

	p.next()
	stmt.Match = p.parseExpression(LOWEST)

	p.next()

	p.skipNewlines()

	if p.peek.Type == lexer.COMMA {
		log.Fatalf("only single equality checks are allowed, got %v", p.cur.Type)
	}

	for p.cur.Type == lexer.CASE {
		p.next()

		// CASE ELSE
		if p.cur.Type == lexer.ELSE {
			p.next()
			stmt.Defaultcase = p.parseCaseBlock()
			break
		}

		// CASE <single expression>
		clause := CasePhrase{
			Value: p.parseExpression(LOWEST),
		}

		clause.Body = p.parseCaseBlock()

		stmt.Cases = append(stmt.Cases, &clause)
	}

	p.skipNewlines()

	// END SELECT
	if p.cur.Type == lexer.END && p.peek.Type == lexer.SELECT {
		p.next()
		p.next()
		return stmt
	}

	log.Fatalf("expected END SELECT, got %v and %v", p.cur.Type, p.peek.Lit)
	return nil
}

func (p *Parser) parseCaseBlock() []Statement {
	var body []Statement

	p.next()

	p.skipNewlines()

	for p.cur.Type != lexer.CASE &&
		p.cur.Type != lexer.END && p.cur.Type != lexer.SELECT {
		if p.cur.Type != lexer.NLINE {
			body = append(body, p.ParseStatement())
		}
		p.next()
	}

	p.skipNewlines()

	return body
}
