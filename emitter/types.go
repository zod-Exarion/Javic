package emitter

import (
	"strings"

	"github.com/zod-Exarion/javic/parser"
)

type VarType int

const (
	VarString VarType = iota
	VarInt
	VarFloat
	VarLong
	VarDouble
)

func inferVarTypeFromName(name string) VarType {
	switch name[len(name)-1] {
	case '$':
		return VarString
	case '%':
		return VarInt
	case '!':
		return VarFloat
	case '#':
		return VarDouble
	case '&':
		return VarLong
	}
	return VarInt // default numeric
}

func (e *Emitter) inferExprType(expr parser.Expression) VarType {
	switch expr.Kind {
	case parser.Number:
		return VarInt
	case parser.Decimal:
		return VarDouble
	case parser.Ident:
		if t, ok := e.symbols[expr.IdentExpression]; ok {
			return t
		}
		return inferVarTypeFromName(expr.IdentExpression)

	case parser.Unary:
		return e.inferExprType(expr.UnaryExpression.Expression)

	case parser.Binary:
		left := e.inferExprType(expr.BinaryExpression.Left)
		right := e.inferExprType(expr.BinaryExpression.Right)

		if left == VarString || right == VarString {
			return VarString
		}
		if left == VarDouble || right == VarDouble {
			return VarDouble
		}
		return VarInt

	default:
		return VarString
	}
}

func javaType(t VarType) string {
	switch t {
	case VarString:
		return "String"
	case VarDouble:
		return "double"
	case VarInt:
		return "int"
	case VarLong:
		return "long"
	case VarFloat:
		return "float"
	default:
		return "String"
	}
}

func sanitize(name string) string {
	if len(name) == 0 {
		return name
	}

	if strings.Contains("$#%!&", string(name[len(name)-1])) {
		switch name[len(name)-1] {
		case '$':
			return name[:len(name)-1] + "_str"
		case '%':
			return name[:len(name)-1] + "_int"
		case '!':
			return name[:len(name)-1] + "_float"
		case '#':
			return name[:len(name)-1] + "_double"
		case '&':
			return name[:len(name)-1] + "_long"
		default:
			return name
		}
	}

	return name
}
