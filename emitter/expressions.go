package emitter

import (
	"strconv"

	"github.com/zod-Exarion/javic/parser"
)

func (e *Emitter) emitExpr(expr parser.Expression) string {
	switch expr.Kind {
	case parser.Number:
		return strconv.FormatInt(expr.NumberExpression, 10)
	case parser.Decimal:
		return strconv.FormatFloat(expr.DecimalExpression, 'f', -1, 64)
	case parser.String:
		return expr.StringExpression
	case parser.Ident:
		return sanitize(expr.IdentExpression)
	case parser.Unary:
		return expr.UnaryExpression.Op + e.emitExpr(expr.UnaryExpression.Expression)
	case parser.Binary:
		op := expr.BinaryExpression.Op
		if op == "MOD" {
			op = "%"
		}
		if op == "=" {
			op = "=="
		}
		return "(" +
			e.emitExpr(expr.BinaryExpression.Left) +
			" " + op + " " +
			e.emitExpr(expr.BinaryExpression.Right) +
			")"
	default:
		panic("unknown expression")
	}
}
