package parser

import (
	"javic/qbasic/tokenizer"
)

type Identifier struct {
	Token tokenizer.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Lit }

type AssignmentStatement struct {
	Token tokenizer.Token
	Name  *Identifier
	Value *Expression
}

func (as *AssignmentStatement) statementNode()       {}
func (as *AssignmentStatement) TokenLiteral() string { return as.Token.Lit }

func (p *Parser) parseStatement() Statement {
	switch p.curToken.Type {
	case tokenizer.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *AssignmentStatement {
	statement := &AssignmentStatement{Token: p.curToken}

	if !p.expectToken(tokenizer.IDENT) {
		return nil
	}

	statement.Name = &Identifier{p.curToken, p.curToken.Lit}

	if !p.expectToken(tokenizer.ASSIGN) {
		return nil
	}

	// TODO: Skipping all literal/expressions until newline for now
	for p.curToken.Type != tokenizer.NLINE {
		p.getNextToken()
	}

	return statement
}

func (p *Parser) expectToken(t tokenizer.TokenType) bool {
	if p.nextToken.Type == t {
		p.getNextToken()
		return true
	} else {
		return false
	}
}
