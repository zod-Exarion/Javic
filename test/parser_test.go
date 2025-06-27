package test

import (
	"javic/cmd/javic"
	"javic/qbasic/parser"
	"testing"
)

func TestLetStatements(t *testing.T) {
	tp := javic.NewTranspiler("qb/parse1.bas", true)
	p := tp.Parser

	program := p.ParseProgram()
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
		t.Errorf("s not *ast.LetStatement. got=%T", s)
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
