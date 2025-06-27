package parser

import (
	"javic/qbasic/lexer"
	"javic/qbasic/tokenizer"
)

type Parser struct {
	lex *lexer.Lexer

	curToken  tokenizer.Token
	nextToken tokenizer.Token
}

func NewParser(lex *lexer.Lexer, flag bool) *Parser {
	p := &Parser{lex: lex}

	if flag {
		p.getNextToken()
		p.getNextToken()
	}

	return p
}

func (p *Parser) getNextToken() {
	p.curToken = p.nextToken
	p.nextToken = p.lex.GetToken()
}

func (p *Parser) ParseProgram() *Program {
	program := &Program{}
	program.Statements = []Statement{}

	for p.curToken.Type != tokenizer.EOF {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.getNextToken()
	}

	return program
}
