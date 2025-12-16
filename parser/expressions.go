package parser

import (
	"log"
	"strconv"

	"github.com/zod-Exarion/javic/lexer"
)

type ExpressionType int

const (
	Number StatementType = iota
	Ident
	Binary
	Unary
)

type Expression struct {
	kind StatementType

	numberExpression int64
	identExpression  string
	binaryExpression *BinaryExpression
	unaryExpression  *UnaryExpression
}

type BinaryExpression struct {
	left  Expression
	op    string
	right Expression
}

type UnaryExpression struct {
	op         string
	expression Expression
}

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

func (p *Parser) parsePrimary() Expression {
	switch p.cur.Type {
	case lexer.NUMBER:
		n, _ := strconv.ParseInt(p.cur.Lit, 10, 64)
		return Expression{kind: Number, numberExpression: n}
	case lexer.IDENT:
		return Expression{kind: Ident, identExpression: p.cur.Lit}
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

func (e Expression) String() string {
	switch e.kind {
	case Number:
		return strconv.FormatInt(e.numberExpression, 10)

	case Ident:
		return e.identExpression

	case Binary:
		return "(" +
			e.binaryExpression.left.String() +
			" " + e.binaryExpression.op + " " +
			e.binaryExpression.right.String() +
			")"

	case Unary:
		return "(" +
			e.unaryExpression.op +
			e.unaryExpression.expression.String() +
			")"

	default:
		return "<unknown expr>"
	}
}
