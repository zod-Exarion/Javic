package parser

import (
	"fmt"
	"javic/qbasic/tokenizer"
	"strconv"
)

const (
	LOWEST = iota
	EQUALS
	LESSGREAT
	SUM
	PRODUCT
	EXPONENT
	PREFIX
	CALL
)

var precedences = map[tokenizer.TokenType]int{
	tokenizer.ASSIGN:   EQUALS,
	tokenizer.LT:       LESSGREAT,
	tokenizer.GT:       LESSGREAT,
	tokenizer.PLUS:     SUM,
	tokenizer.MINUS:    SUM,
	tokenizer.SLASH:    PRODUCT,
	tokenizer.ASTERISK: PRODUCT,
	tokenizer.CARET:    EXPONENT,
}

func (p *Parser) parseExpression(precedence int) Expression {
	leftExp := p.parsePrimary()
	if leftExp == nil {
		return nil
	}

	for p.curToken.Type != tokenizer.NLINE && precedence < p.peekPrecedence() {
		infix := p.parseInfixExpression
		// if infix == nil {
		// 	return leftExp
		// }

		p.getNextToken() // consume operator
		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parsePrimary() Expression {
	switch p.curToken.Type {
	case tokenizer.NUMBER:
		return p.parseIntegerLiteral()
	case tokenizer.IDENT:
		return p.parseIdentifier()
	case tokenizer.MINUS:
		return p.parsePrefixExpression()
	default:
		return nil
	}
}

func (p *Parser) parsePrefixExpression() Expression {
	expr := &PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Lit,
	}

	p.getNextToken()

	expr.Right = p.parseExpression(PREFIX)

	return expr
}

func (p *Parser) parseInfixExpression(left Expression) Expression {
	expr := &InfixExpression{
		Token: p.curToken,
		Left:  left,
	}

	// HACK: Exception case for equality/assignment operator
	if p.curToken.Type != tokenizer.ASSIGN {
		expr.Operator = p.curToken.Lit
	} else {
		expr.Operator = "=="
	}

	precedence := p.curPrecedence()
	p.getNextToken() // move to right operand

	expr.Right = p.parseExpression(precedence)
	return expr
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.nextToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) parseIdentifier() Expression {
	return &Identifier{Token: p.curToken, Value: p.curToken.Lit}
}

func (p *Parser) parseIntegerLiteral() *IntegerLiteral {
	lit := &IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Lit, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("parser error: couldnt parse %q as integer", p.curToken.Lit)
		p.errors = append(p.errors, msg)
	}
	lit.Value = value

	return lit
}
