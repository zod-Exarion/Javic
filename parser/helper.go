package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zod-Exarion/javic/lexer"
)

// INFO: Basic helper functions for making the parsed product human readable

func (e Expression) String() string {
	switch e.kind {
	case Number:
		return strconv.FormatInt(e.numberExpression, 10)

	case Ident:
		return e.identExpression

	case String:
		return e.stringExpression

	case Binary:
		return "(" +
			e.binaryExpression.left.String() +
			" " + e.binaryExpression.op + " " +
			e.binaryExpression.right.String() +
			")"

	case Unary:
		return "(" +
			e.unaryExpression.op +
			e.unaryExpression.expression.String() +
			")"

	default:
		return "<unknown expr>"
	}
}

func (s Statement) String() string {
	switch s.kind {
	case Let:
		return "LET " + s.letStatement.name + " = " + s.letStatement.value.String()

	case Print:
		return "PRINT " + s.printStatement.value.String()

	case Input:
		return "INPUT " + s.inputStatement.prompt + " , " + s.inputStatement.name

	case If:
		return s.ifStatement.String()

	case For:
		return s.forStatement.String()

	case Cls:
		return "CLS"

	case End:
		return "END"

	default:
		return "<unknown stmt>"
	}
}

func (i *IfStatement) String() string {
	var out strings.Builder

	out.WriteString("IF ")
	out.WriteString(i.lhscondition.String())
	out.WriteString(" ")
	out.WriteString(i.comparison)
	out.WriteString(" ")
	out.WriteString(i.rhscondition.String())
	out.WriteString(" THEN\n")

	for _, stmt := range i.consequence {
		out.WriteString(stmt.String())
		out.WriteString("\n")
	}

	if len(i.antecendent) > 0 {
		out.WriteString("ELSE\n")
		for _, stmt := range i.antecendent {
			out.WriteString(stmt.String())
			out.WriteString("\n")
		}
	}

	out.WriteString("END IF")

	return out.String()
}

func (f *ForStatement) String() string {
	var out strings.Builder

	out.WriteString("FOR ")
	out.WriteString(f.name)
	out.WriteString(" = ")
	out.WriteString(f.init.String())
	out.WriteString(" TO ")
	out.WriteString(f.final.String())
	out.WriteString(" STEP ")
	out.WriteString(f.step.String() + "\n")

	for _, stmt := range f.body {
		out.WriteString(stmt.String())
		out.WriteString("\n")
	}

	out.WriteString("NEXT")

	return out.String()
}

func DisplayStatements(statements []Statement) {
	for i := range statements {
		fmt.Printf("%v\n", statements[i])
	}
}

func (p *Parser) skipNewlines() {
	for p.cur.Type == lexer.NLINE {
		p.next()
	}
}
