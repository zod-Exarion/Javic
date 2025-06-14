package parser

import "javic/qbasic/tokenizer"

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
