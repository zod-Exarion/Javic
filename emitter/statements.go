package emitter

import (
	"strings"

	"github.com/zod-Exarion/javic/parser"
)

func (e *Emitter) emitLet(stmt *parser.LetStatement) {
	name := stmt.Name
	expr := stmt.Value

	javaName := sanitize(name)

	// infer type
	exprType := e.inferExprType(expr)

	if _, exists := e.symbols[name]; !exists {
		// FIRST TIME → DECLARATION
		e.symbols[name] = exprType

		e.emitLine(
			javaType(exprType) + " " +
				javaName + " = " +
				e.emitExpr(expr) + ";",
		)
	} else {
		// ALREADY DECLARED → ASSIGNMENT
		e.emitLine(
			javaName + " = " +
				e.emitExpr(expr) + ";",
		)
	}
}

func (e *Emitter) emitPrint(s *parser.PrintStatement) {
	e.imports.WriteString("import java.util.Scanner;\n")

	parts := []string{}
	for _, v := range s.Value {
		parts = append(parts, e.emitExpr(v))
	}
	e.emitLine("System.out.println(" + strings.Join(parts, " + \" \" + ") + ");")
}

func (e *Emitter) emitInput(s *parser.InputStatement) {
	name := s.Name
	varType := inferVarTypeFromName(name)

	javaName := sanitize(name)

	// Declare if needed
	if _, ok := e.symbols[name]; !ok {
		e.symbols[name] = varType
		e.emitLine(javaType(varType) + " " + javaName + ";")
	}

	if s.Prompt != "" {
		e.emitLine(`System.out.print("` + s.Prompt + `");`)
	}

	switch e.symbols[name] {
	case VarString:
		e.emitLine(javaName + " = scanner.nextLine();")
	case VarDouble:
		e.emitLine(javaName + " = scanner.nextDouble();")
	case VarInt:
		e.emitLine(javaName + " = scanner.nextInt();")
	case VarFloat:
		e.emitLine(javaName + " = scanner.nextFloat();")
	case VarLong:
		e.emitLine(javaName + " = scanner.nextLong();")
	}
}

func (e *Emitter) emitIf(s *parser.IfStatement) {
	e.emitLine("if (" + e.emitExpr(s.Condition) + ") {")
	e.indent++
	for _, st := range s.Consequence {
		e.emitStatement(st)
	}
	e.indent--
	e.emitLine("}")

	if len(s.Antecendent) > 0 {
		e.emitLine("else {")
		e.indent++
		for _, st := range s.Antecendent {
			e.emitStatement(st)
		}
		e.indent--
		e.emitLine("}")
	}
}

func (e *Emitter) emitWhile(s *parser.WhileStatement) {
	e.emitLine("while (" + e.emitExpr(s.Condition) + ") {")
	e.indent++
	for _, st := range s.Body {
		e.emitStatement(st)
	}
	e.indent--
	e.emitLine("}")
}

func (e *Emitter) emitFor(s *parser.ForStatement) {
	step := "1"
	if s.Step.Kind != 0 {
		step = e.emitExpr(s.Step)
	}

	e.emitLine(
		"for (int " + s.Name +
			" = " + e.emitExpr(s.Init) +
			"; " + s.Name + " <= " + e.emitExpr(s.Final) +
			"; " + s.Name + " += " + step + ") {",
	)

	e.indent++
	for _, st := range s.Body {
		e.emitStatement(st)
	}
	e.indent--
	e.emitLine("}")
}

func (e *Emitter) emitSelect(s *parser.SelectStatement) {
	e.emitLine("switch (" + e.emitExpr(s.Match) + ") {")
	e.indent++

	for _, c := range s.Cases {
		e.emitLine("case " + e.emitExpr(c.Value) + ":")
		e.indent++
		for _, st := range c.Body {
			e.emitStatement(st)
		}
		e.emitLine("break;")
		e.indent--
	}

	if len(s.Defaultcase) > 0 {
		e.emitLine("default:")
		e.indent++
		for _, st := range s.Defaultcase {
			e.emitStatement(st)
		}
		e.emitLine("break;")
		e.indent--
	}

	e.indent--
	e.emitLine("}")
}
