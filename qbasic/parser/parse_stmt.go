package parser

import "javic/qbasic/tokenizer"

func (p *Parser) parseStatement() Statement {
	switch p.curToken.Type {
	case tokenizer.LET:
		return p.parseLetStatement()
	case tokenizer.RETURN:
		return p.parseReturnStatement()
	case tokenizer.PRINT:
		return p.parsePrintStatement()
	case tokenizer.INPUT:
		return p.parseInputStatement()
	case tokenizer.NOT:
		return p.parseNotStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *AssignmentStatement {
	statement := &AssignmentStatement{
		DoubleOperandStatmenet: DoubleOperandStatmenet{
			SingleOperandStatmenet: SingleOperandStatmenet{
				Token: p.curToken,
			},
		},
	}

	if !p.expectToken(tokenizer.IDENT) {
		return nil
	}

	statement.Name = &Identifier{p.curToken, p.curToken.Lit}

	if !p.expectToken(tokenizer.ASSIGN) {
		return nil
	}

	p.getNextToken()

	statement.Value = p.parseExpression(LOWEST)

	// TODO: Skipping all literal/expressions until newline for now
	for p.curToken.Type != tokenizer.NLINE {
		p.getNextToken()
	}

	return statement
}

func (p *Parser) parseReturnStatement() *ReturnStatement {
	statement := &ReturnStatement{
		SingleOperandStatmenet: SingleOperandStatmenet{
			Token: p.curToken,
		},
	}

	p.getNextToken()

	statement.Value = p.parseExpression(LOWEST)

	// TODO: Skipping all literal/expressions until newline for now
	for p.curToken.Type != tokenizer.NLINE {
		p.getNextToken()
	}

	return statement
}

func (p *Parser) parsePrintStatement() *PrintStatement {
	statement := &PrintStatement{
		SingleOperandStatmenet: SingleOperandStatmenet{
			Token: p.curToken,
		},
	}

	p.getNextToken()

	statement.Value = p.parseExpression(LOWEST)

	// TODO: Skipping all literal/expressions until newline for now
	for p.curToken.Type != tokenizer.NLINE {
		p.getNextToken()
	}

	return statement
}

func (p *Parser) parseInputStatement() *InputStatement {
	statement := &InputStatement{
		DoubleOperandStatmenet: DoubleOperandStatmenet{
			SingleOperandStatmenet: SingleOperandStatmenet{
				Token: p.curToken,
			},
		},
	}

	// TODO: Input Prompt here.

	if !p.expectToken(tokenizer.COMMA) {
		return nil
	}

	if !p.expectToken(tokenizer.IDENT) {
		return nil
	}

	statement.Name = &Identifier{p.curToken, p.curToken.Lit}

	// TODO: Skipping all literal/expressions until newline for now
	for p.curToken.Type != tokenizer.NLINE {
		p.getNextToken()
	}

	return statement
}

// TODO: Introduce support for NOT statements nested within other statements.
func (p *Parser) parseNotStatement() *NotStatement {
	statement := &NotStatement{
		SingleOperandStatmenet: SingleOperandStatmenet{
			Token: p.curToken,
		},
	}

	p.getNextToken()

	statement.Value = p.parseExpression(LOWEST)

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
		p.nextError(t)
		return false
	}
}
