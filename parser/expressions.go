package parser

import (
	"log"
	"strconv"
	"strings"

	"github.com/zod-Exarion/javic/lexer"
)

type ExpressionType int // another wrapper around int for the ease of the developer

// constant wrappers for the ease of the developer
const (
	Number StatementType = iota
	Decimal
	Ident
	String
	Binary
	Unary
	And
	Or
	Not
)

// HACK: Same logic as Statement struct: interfaces too hard :( just create struct where only the respective
// (detected expression type's respective field) is non-nil
type Expression struct {
	Kind StatementType

	NumberExpression  int64
	DecimalExpression float64
	IdentExpression   string
	StringExpression  string
	BinaryExpression  *BinaryExpression
	UnaryExpression   *UnaryExpression
}

// INFO: Here we define the expressions
type BinaryExpression struct {
	Left  Expression
	Op    string
	Right Expression
}

type UnaryExpression struct {
	Op         string
	Expression Expression
}

// INFO: Here we define the functions to parse the expressions

// PERF: Signature Pratt-Parsing recursive descent (lot of fancy words)
//
// 1. call our parsePrimary function to prase the type of expression
//
// 2. If the next token is a newline or if the precedence of it is =< our current token, we simply return the expression
// e.g: let x = 20, so our 'Expression' part, [Value] in LetStatmenet is the expression
// Numbers have the precedence zero, so it simply returns 20 without recursion
//
// 3. now it gets juicy lol. If it isn't an nline and the precedence is greater than current.
// store the operator and RECURSE over the right side
// e.g let x = 20 + 30, it would first get 20, but then create binaryExpresion {left: 20, op: +, right: 20}
// If its somethjing like 20 + 30 + 40 + 50, then it keeps recursing the right side from [30 + 40 + 50] to [40 + 50] to [50]
func (p *Parser) parseExpression(prec int) Expression {
	left := p.parsePrimary()

	for !p.expressionEnd() && p.peekPrecedence() > prec {
		p.next() // skips to operator
		op := p.cur.Lit

		currPrec := p.curPrecedence()

		p.next() // skips to right expression
		right := p.parseExpression(currPrec)

		left = Expression{
			Kind: Binary,
			BinaryExpression: &BinaryExpression{
				Left:  left,
				Op:    op,
				Right: right,
			},
		}
	}

	return left
}

// INFO: checks what type of token it is and returns the respective expression struct
func (p *Parser) parsePrimary() Expression {
	switch p.cur.Type {
	case lexer.NUMBER:
		if strings.Contains(p.cur.Lit, ".") {
			f, err := strconv.ParseFloat(p.cur.Lit, 64)
			if err != nil {
				log.Fatalf("Invalid decimal literal: %s", p.cur.Lit)
			}
			return Expression{Kind: Decimal, DecimalExpression: f}
		}

		i, err := strconv.ParseInt(p.cur.Lit, 10, 64)
		if err != nil {
			log.Fatalf("Invalid integer literal: %s", p.cur.Lit)
		}

		return Expression{Kind: Number, NumberExpression: i}

	case lexer.IDENT:
		return Expression{Kind: Ident, IdentExpression: p.cur.Lit} // simply return identifier
	case lexer.LPAREN:
		return p.parseGroupedExpression()
	case lexer.STRING:
		return Expression{Kind: String, StringExpression: p.cur.Lit}
	case lexer.NOT:
		return p.parseNotExpression()
	default:
		log.Fatal("Unrecognized expression")
		return Expression{} // unrecognizable expression, doesnt get added in the final list of expressions
	}
}

// PERF: Prone to off-by-one errors, DO NOT TOUCH.
func (p *Parser) parseGroupedExpression() Expression {
	p.next() // skips to expression
	expr := p.parseExpression(0)

	p.expectToken(lexer.RPAREN) // expect the right parenthesis

	// p.next() // consumes ')' // WARN: If we skip to next, creates parser off by one error
	return expr
}

func (p *Parser) parseNotExpression() Expression {
	p.next()
	expr := p.parseExpression(PREFIX)
	return Expression{
		Kind: Unary,
		UnaryExpression: &UnaryExpression{
			Op:         "NOT",
			Expression: expr,
		},
	}
}

func (p *Parser) expressionEnd() bool {
	switch p.peek.Type {
	case lexer.THEN,
		lexer.ELSE,
		lexer.WEND,
		lexer.END,
		// lexer.NLINE, removed for while parsing issues
		lexer.EOF:
		return true
	}
	return false
}
