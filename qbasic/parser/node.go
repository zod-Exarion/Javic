package parser

type Node interface {
	TokenLiteral() string // HACK: Debugging usage
}

type Statement interface {
	Node
	statementNode() // HACK: Dummy method for compilation security
}

type Expression interface {
	Node
	expressionNode() // HACK: Dummy method for compilation security
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
