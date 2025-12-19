package parser

import "github.com/zod-Exarion/javic/lexer"

const (
	LOWEST  = iota
	OR      // OR
	AND     // AND
	EQUALS  // == <>
	COMPARE //  <= >=
	SUM     // + -
	PRODUCT // * / ^
	PREFIX  // NOT
)

// INFO: Here we define the precedence for all the operators to their respective numberical heirarchy
var precedences = map[lexer.TokenType]int{
	lexer.OR:  OR,
	lexer.AND: AND,

	lexer.ASSIGN: EQUALS,
	lexer.NEQ:    EQUALS,

	lexer.LT:  COMPARE,
	lexer.GT:  COMPARE,
	lexer.LTE: COMPARE,
	lexer.GTE: COMPARE,

	lexer.PLUS:  SUM,
	lexer.MINUS: SUM,

	lexer.ASTERISK: PRODUCT,
	lexer.SLASH:    PRODUCT,
	lexer.CARET:    PRODUCT,
	lexer.MOD:      PRODUCT,
}

// INFO: return the precende of *Parser.peek
func (p *Parser) peekPrecedence() int {
	if prec, ok := precedences[p.peek.Type]; ok {
		return prec
	}
	return 0
}

// INFO: return the precende of *Parser.cur
func (p *Parser) curPrecedence() int {
	if prec, ok := precedences[p.cur.Type]; ok {
		return prec
	}
	return 0
}
