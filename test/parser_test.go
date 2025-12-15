package test

import (
	"fmt"
	"javic/cmd/javic"
	"javic/qbasic/parser"
	"strings"
	"testing"
)

func checkParseErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestLetStatements(t *testing.T) {
	tp := javic.NewTranspiler("qb/parse1.bas", true)
	p := tp.Parser

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d \n\n",
			len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"X"},
		{"Y"},
		{"FOOBAR"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s parser.Statement, name string) bool {
	if s.TokenLiteral() != "LET" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*parser.AssignmentStatement)
	if !ok {
		t.Errorf("s not *LetStatement. got=%T", s)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}
	return true
}

func TestReturnStatements(t *testing.T) {
	tp := javic.NewTranspiler("qb/parse2.bas", true)
	p := tp.Parser

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*parser.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *returnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "RETURN" {
			t.Errorf("returnStmt.TokenLiteral not 'RETURN', got %q",
				returnStmt.TokenLiteral())
		}
	}
}

func TestLetIdentifierExpression(t *testing.T) {
	tp := javic.NewTranspiler("qb/parse3.bas", true)
	p := tp.Parser

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*parser.AssignmentStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not AssignmentStatement. got=%T", program.Statements[0])
	}

	// Check left-hand side identifier
	if stmt.Name.Value != "X" {
		t.Errorf("stmt.Name.Value not 'X'. got=%q", stmt.Name.Value)
	}

	// Check right-hand side is identifier 'foobar'
	ident, ok := stmt.Value.(*parser.Identifier)
	if !ok {
		t.Fatalf("stmt.Value is not *Identifier. got=%T", stmt.Value)
	}

	if ident.Value != "FOOBAR" {
		t.Errorf("ident.Value not 'FOOBAR'. got=%q", ident.Value)
	}

	if ident.TokenLiteral() != "FOOBAR" {
		t.Errorf("ident.TokenLiteral not 'FOOBAR'. got=%q", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	tp := javic.NewTranspiler("qb/parse4.bas", true)
	p := tp.Parser

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*parser.AssignmentStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not AssignmentStatement. got=%T", program.Statements[0])
	}

	// Check left-hand side identifier
	if stmt.Name.Value != "X" {
		t.Errorf("stmt.Name.Value not 'X'. got=%q", stmt.Name.Value)
	}

	// Check right-hand side integerliteral
	intlit, ok := stmt.Value.(*parser.IntegerLiteral)
	if !ok {
		t.Fatalf("stmt.Value is not *IntegerLiteral. got=%T", stmt.Value)
	}

	if intlit.Value != 55 {
		t.Errorf("intlit.Value not 55. got=%d", intlit.Value)
	}

	if intlit.TokenLiteral() != "55" {
		t.Errorf("intlit.TokenLiteral not '55'. got=%q", intlit.TokenLiteral())
	}
}

func testIntegerLiteral(t *testing.T, il parser.Expression, value int64) bool {
	integ, ok := il.(*parser.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}
	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}
	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value,
			integ.TokenLiteral())
		return false
	}
	return true
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"LET A = 5 + 5\n", 5, "+", 5},
		{"LET A = 5 - 5\n", 5, "-", 5},
		{"LET A = 5 * 5\n", 5, "*", 5},
		{"LET A = 5 / 5\n", 5, "/", 5},
		{"LET A = 5 ^ 5\n", 5, "^", 5},
		{"LET A = 5 > 5\n", 5, ">", 5},
		{"LET A = 5 < 5\n", 5, "<", 5},
		// {"LET A = 5 = 5\n", 5, "=", 5}, // enable if you support '=' as comparison
	}

	for _, tt := range infixTests {
		p := javic.NewTranspilerFromString(tt.input, true).Parser
		program := p.ParseProgram()
		checkParseErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statement. got=%d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*parser.AssignmentStatement)
		if !ok {
			t.Fatalf("stmt is not *AssignmentStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Value.(*parser.InfixExpression)
		if !ok {
			t.Fatalf("stmt.Value is not *InfixExpression. got=%T", stmt.Value)
		}

		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
			return
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"LET X = -A * B\n",
			"LET X = ((-A) * B)\n",
		},
		{
			"LET X = A + B + C\n",
			"LET X = ((A + B) + C)\n",
		},
		{
			"LET X = A + B - C\n",
			"LET X = ((A + B) - C)\n",
		},
		{
			"LET X = A * B * C\n",
			"LET X = ((A * B) * C)\n",
		},
		{
			"LET X = A * B / C\n",
			"LET X = ((A * B) / C)\n",
		},
		{
			"LET X = A + B / C\n",
			"LET X = (A + (B / C))\n",
		},
		{
			"LET X = A + B * C + D / E - F\n",
			"LET X = (((A + (B * C)) + (D / E)) - F)\n",
		},
		{
			"LET X = 3 = 4\nLET Y = -5 * 5\n",
			"LET X = (3 == 4)\nLET Y = ((-5) * 5)\n",
		},
	}

	for _, tt := range tests {
		// Parse using string input (simulating a QBASIC file)
		p := javic.NewTranspilerFromString(tt.input, true).Parser
		program := p.ParseProgram()
		checkParseErrors(t, p)

		// Get actual output from program's String()
		actual := program.String()

		// Compare in uppercase
		if actual != strings.ToUpper(tt.expected) {
			t.Errorf("\nInput: %q\nExpected: %q\nGot:      %q", tt.input, strings.ToUpper(tt.expected), actual)
		}
	}
}
