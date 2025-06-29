package parser

import (
	"bytes"
	"fmt"
	"javic/qbasic/tokenizer"
)

type Node interface {
	TokenLiteral() string // HACK: Debugging usage
	String() string
}

type Statement interface {
	Node
	statementNode() // HACK: Dummy method for compilation security
}

type Expression interface {
	Node
	expressionNode() // HACK: Dummy method for compilation security
}

type Identifier struct {
	Token tokenizer.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Lit }
func (i *Identifier) String() string       { return i.Value }

type IntegerLiteral struct {
	Token tokenizer.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Lit }
func (il *IntegerLiteral) String() string {
	return fmt.Sprintf("%d", il.Value)
}

type InfixExpression struct {
	Token    tokenizer.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Lit }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

type PrefixExpression struct {
	Token    tokenizer.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Lit }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
		out.WriteString("\n")
	}

	return out.String()
}
