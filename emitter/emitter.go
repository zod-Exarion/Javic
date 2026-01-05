package emitter

import (
	"strings"

	"github.com/zod-Exarion/javic/parser"
)

type Emitter struct {
	out        strings.Builder
	indent     int
	imports    strings.Builder
	terminates strings.Builder
	symbols    map[string]VarType
}

func Emit(stmts []parser.Statement) string {
	e := &Emitter{
		symbols: make(map[string]VarType),
	}

	e.emitLine("public class Main {")
	e.indent++
	e.emitLine("public static void main(String[] args) {")
	e.indent++

	for _, stmt := range stmts {
		e.emitStatement(stmt)
	}

	e.emitLine(e.terminates.String())

	e.indent--
	e.emitLine("}")
	e.indent--
	e.emitLine("}")

	return e.imports.String() + e.out.String()
}

func (e *Emitter) emitLine(s string) {
	e.out.WriteString(strings.Repeat("    ", e.indent))
	e.out.WriteString(s)
	e.out.WriteString("\n")
}

func (e *Emitter) emitStatement(stmt parser.Statement) {
	switch stmt.Kind {
	case parser.Let:
		e.emitLet(stmt.LetStatement)
	case parser.Print:
		e.emitPrint(stmt.PrintStatement)
	case parser.Input:
		e.emitInput(stmt.InputStatement)
	case parser.If:
		e.emitIf(stmt.IfStatement)
	case parser.For:
		e.emitFor(stmt.ForStatement)
	case parser.While:
		e.emitWhile(stmt.WhileStatement)
	case parser.Select:
		e.emitSelect(stmt.SelectStatement)
	case parser.Cls:
		// ignore
	case parser.Rem:
		e.emitLine("// " + stmt.RemStatement.Body)
	case parser.End:
		// ignore
	default:
		panic("unknown statement")
	}
}

func (e *Emitter) checkScanner() {
	if !scannerFlag {
		e.imports.WriteString("import java.util.Scanner;\n")
		e.emitLine("Scanner scanner = new Scanner(System.in);")
		e.terminates.WriteString("scanner.close()")
		scannerFlag = true
	}
}
