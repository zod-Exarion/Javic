package parser

import (
	"bytes"
	"javic/qbasic/tokenizer"
)

type SingleOperandStatmenet struct {
	Token tokenizer.Token
	Value Expression
}

func (ss *SingleOperandStatmenet) statementNode()       {}
func (ss *SingleOperandStatmenet) TokenLiteral() string { return ss.Token.Lit }

type DoubleOperandStatmenet struct {
	SingleOperandStatmenet
	Name *Identifier
}

func (s *SingleOperandStatmenet) String() string {
	var out bytes.Buffer

	out.WriteString(s.TokenLiteral() + " ")

	if s.Value != nil {
		out.WriteString(s.Value.String())
	}

	return out.String()
}

func (s *DoubleOperandStatmenet) String() string {
	var out bytes.Buffer

	out.WriteString(s.TokenLiteral() + " ")
	out.WriteString(s.Name.String())
	out.WriteString(" = ")

	if s.Value != nil {
		out.WriteString(s.Value.String())
	}

	return out.String()
}

func (s *InputStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.TokenLiteral() + " ")

	if s.Value != nil {
		out.WriteString(s.Value.String())
	}

	out.WriteString(" , ")
	out.WriteString(s.Name.String())

	return out.String()
}

type (
	AssignmentStatement struct{ DoubleOperandStatmenet }
	ReturnStatement     struct{ SingleOperandStatmenet }
	PrintStatement      struct{ SingleOperandStatmenet }
	NotStatement        struct{ SingleOperandStatmenet }
	InputStatement      struct{ DoubleOperandStatmenet }
)
