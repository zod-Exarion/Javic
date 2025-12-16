package parser

import "github.com/zod-Exarion/javic/lexer"

var precedences = map[lexer.TokenType]int{
	lexer.PLUS:     1,
	lexer.MINUS:    1,
	lexer.ASTERISK: 2,
	lexer.SLASH:    2,
}

func (p *Parser) peekPrecedence() int {
	if prec, ok := precedences[p.peek.Type]; ok {
		return prec
	}
	return 0
}

func (p *Parser) curPrecedence() int {
	if prec, ok := precedences[p.cur.Type]; ok {
		return prec
	}
	return 0
}
