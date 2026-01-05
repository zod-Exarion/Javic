package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zod-Exarion/javic/lexer"
)

// INFO: Basic helper functions for making the parsed product human readable

func (e Expression) String() string {
	switch e.Kind {
	case Number:
		return strconv.FormatInt(e.NumberExpression, 10)

	case Decimal:
		return strconv.FormatFloat(e.DecimalExpression, 'f', -1, 64)

	case Ident:
		return e.IdentExpression

	case String:
		return e.StringExpression

	case Binary:
		return "(" +
			e.BinaryExpression.Left.String() +
			" " + e.BinaryExpression.Op + " " +
			e.BinaryExpression.Right.String() +
			")"

	case Unary:
		return "(" +
			e.UnaryExpression.Op +
			e.UnaryExpression.Expression.String() +
			")"

	default:
		return "<unknown expr>"
	}
}

func (s Statement) String() string {
	switch s.Kind {
	case Rem:
		return "REM " + s.RemStatement.Body
	case Let:
		return "LET " + s.LetStatement.Name + " = " + s.LetStatement.Value.String()

	case Print:
		return s.PrintStatement.String()

	case Input:
		return "INPUT " + s.InputStatement.Prompt + " , " + s.InputStatement.Name

	case If:
		return s.IfStatement.String()

	case For:
		return s.ForStatement.String()

	case Cls:
		return "CLS"

	case End:
		return "END"

	case While:
		return s.WhileStatement.String()

	case Dim:
		return "DIM " + s.DimStatement.Name + " AS " + s.DimStatement.Datatype

	case Redim:
		return "REDIM " + s.RedimStatement.Name

	case Select:
		return s.SelectStatement.String()

	default:
		return "<unknown stmt>"
	}
}

func (i *IfStatement) String() string {
	var out strings.Builder

	out.WriteString("IF ")
	out.WriteString(i.Condition.String())
	out.WriteString(" THEN\n")

	for _, stmt := range i.Consequence {
		out.WriteString(stmt.String())
		out.WriteString("\n")
	}

	if len(i.Antecendent) > 0 {
		out.WriteString("ELSE\n")
		for _, stmt := range i.Antecendent {
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
	out.WriteString(f.Name)
	out.WriteString(" = ")
	out.WriteString(f.Init.String())
	out.WriteString(" TO ")
	out.WriteString(f.Final.String())
	out.WriteString(" STEP ")
	out.WriteString(f.Step.String() + "\n")

	for _, stmt := range f.Body {
		out.WriteString(stmt.String())
		out.WriteString("\n")
	}

	out.WriteString("NEXT")

	return out.String()
}

func (w *WhileStatement) String() string {
	var out strings.Builder

	out.WriteString("WHILE ")
	out.WriteString(w.Condition.String() + "\n")
	for _, stmt := range w.Body {
		out.WriteString(stmt.String())
		out.WriteString("\n")
	}
	out.WriteString("WEND")

	return out.String()
}

func (s *SelectStatement) String() string {
	var out strings.Builder

	out.WriteString("SELECT CASE ")
	out.WriteString(s.Match.String())
	out.WriteString("\n")

	for _, c := range s.Cases {
		out.WriteString("CASE ")
		out.WriteString(c.Value.String())
		out.WriteString("\n")
		for _, stmt := range c.Body {
			out.WriteString("  ")
			out.WriteString(stmt.String())
			out.WriteString("\n")
		}
	}

	if s.Defaultcase != nil {
		out.WriteString("CASE ELSE\n")
		for _, stmt := range s.Defaultcase {
			out.WriteString("  ")
			out.WriteString(stmt.String())
			out.WriteString("\n")
		}
	}

	out.WriteString("END SELECT")
	return out.String()
}

func (p *PrintStatement) String() string {
	var out strings.Builder

	out.WriteString("PRINT ")
	for _, expr := range p.Value {
		out.WriteString(expr.String())
		out.WriteString(", ")
	}

	return out.String()
}

func DisplayStatements(statements []Statement) {
	for i := range statements {
		fmt.Printf("%v\n", statements[i])
	}
	fmt.Println("\n\n\n")
}

func (p *Parser) skipNewlines() {
	for p.cur.Type == lexer.NLINE {
		p.next()
	}
}
