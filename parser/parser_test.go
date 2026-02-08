package parser

import (
	"custom-interpreter/ast"
	"custom-interpreter/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	inp := `let x = 5;
	let y = 10;
	let foobar = 838383;`
	l := lexer.New(inp)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if stmt.TokenLiteral() != "let" {
			t.Fatalf("stmt.TokenLiteral not 'let'. got=%q", stmt.TokenLiteral())
		}
		letStmt, ok := stmt.(*ast.LetStatement)
		if !ok {
			t.Fatalf("stmt not *ast.LetStatement. got=%T", stmt)
		}
		if letStmt.Name.Value != tt.expectedIdentifier {
			t.Fatalf("letStmt.Name.Value not '%s'. got=%s", tt.expectedIdentifier, letStmt.Name.Value)
		}
		if letStmt.Name.TokenLiteral() != tt.expectedIdentifier {
			t.Fatalf("letStmt.Name.TokenLiteral() not '%s'. got=%s", tt.expectedIdentifier, letStmt.Name.TokenLiteral())
		}
	}
}
func TestLetStatementsErrors(t *testing.T) {
	inp := `let x 5;
	let = 10;
	let 838383;`
	l := lexer.New(inp)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	errors := p.Errors()
	if len(errors) != 3 {
		t.Fatalf("parser should have 3 errors. got=%d", len(errors))
	}
}
func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
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

func checkParserErrors(t *testing.T, p *Parser) {
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
func TestReturnStatements(t *testing.T) {
	inp := `return 5;
	return 10;
	return 993322;`
	l := lexer.New(inp)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Fatalf("stmt not *ast.ReturnStatement. got=%T", stmt)
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Fatalf("returnStmt.TokenLiteral not 'return'. got=%q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	inp := "test;"
	l := lexer.New(inp)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "test" {
		t.Errorf("ident.Value not 'test'. got=%s", ident.Value)
	}
	if ident.TokenLiteral() != "test" {
		t.Errorf("ident.TokenLiteral not 'test'. got=%s", ident.TokenLiteral())
	}
}

func TestIntegerExpression(t *testing.T) {
	inp := "5;"
	l := lexer.New(inp)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}
	intExp, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	if intExp.Value != 5 {
		t.Errorf("intExp.Value not 5. got=%d", intExp.Value)
	}
	if intExp.TokenLiteral() != "5" {
		t.Errorf("intExp.TokenLiteral not '5'. got=%s", intExp.TokenLiteral())
	}
}

func TestParsingPrefix(t *testing.T) {
	prefixTests := []struct {
		input      string
		operator   string
		integerVal int64
	}{
		{"-15;", "-", 15},
		{"!5;", "!", 5},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
		}
	}
}
