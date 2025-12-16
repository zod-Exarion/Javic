package parser

import (
	"log"
	"strconv"

	"github.com/zod-Exarion/javic/lexer"
)

type ExpressionType int // another wrapper around int for the ease of the developer

// constant wrappers for the ease of the developer
const (
	Number StatementType = iota
	Ident
	Binary
	Unary
)

// HACK: Same logic as Statement struct: interfaces too hard :( just create struct where only the respective
// (detected expression type's respective field) is non-nil
type Expression struct {
	kind StatementType

	numberExpression int64
	identExpression  string
	binaryExpression *BinaryExpression
	unaryExpression  *UnaryExpression
}

// INFO: Here we define the expressions
type BinaryExpression struct {
	left  Expression
	op    string
	right Expression
}

type UnaryExpression struct {
	op         string
	expression Expression
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

	for p.peek.Type != lexer.NLINE && p.peekPrecedence() > prec {
		p.next() // skips to operator
		op := p.cur.Lit

		currPrec := p.curPrecedence()

		p.next() // skips to right expression
		right := p.parseExpression(currPrec)

		left = Expression{
			kind: Binary,
			binaryExpression: &BinaryExpression{
				left:  left,
				op:    op,
				right: right,
			},
		}
	}

	return left
}

// INFO: checks what type of token it is and returns the respective expression struct
func (p *Parser) parsePrimary() Expression {
	switch p.cur.Type {
	case lexer.NUMBER:
		n, _ := strconv.ParseInt(p.cur.Lit, 10, 64) // converts string to number
		return Expression{kind: Number, numberExpression: n}
	case lexer.IDENT:
		return Expression{kind: Ident, identExpression: p.cur.Lit} // simply return identifier
	case lexer.LPAREN:
		p.next() // skips to expression
		expr := p.parseExpression(0)

		if p.cur.Type != lexer.RPAREN {
			log.Fatal("Expected ')' after expression")
		}

		p.next() // consumes ')'
		return expr
	default:
		log.Fatal("Unrecognized expression")
		return Expression{} // unrecognizable expression, doesnt get added in the final list of expressions
	}
}
