package parser

import (
	"fmt"
	"strconv"
)

// INFO: Basic helper functions for making the parsed product human readable

func (e Expression) String() string {
	switch e.kind {
	case Number:
		return strconv.FormatInt(e.numberExpression, 10)

	case Ident:
		return e.identExpression

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

	default:
		return "<unknown stmt>"
	}
}

func DisplayStatements(statements []Statement) {
	for i := range statements {
		fmt.Printf("%v\n", statements[i])
	}
}
