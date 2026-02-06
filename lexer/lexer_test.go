package lexer

import (
	"custom-interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){};,"
	tests := []struct {
		ExpectedType    token.TokenType
		ExpectedLiteral string
	}{
		{ExpectedType: token.ASSIGN, ExpectedLiteral: "="},
		{ExpectedType: token.PLUS, ExpectedLiteral: "+"},
		{ExpectedType: token.LPAREN, ExpectedLiteral: "("},
		{ExpectedType: token.RPAREN, ExpectedLiteral: ")"},
		{ExpectedType: token.LBRACE, ExpectedLiteral: "{"},
		{ExpectedType: token.RBRACE, ExpectedLiteral: "}"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},
		{ExpectedType: token.COMMA, ExpectedLiteral: ","},
	}
	l := New(input)
	for i, test := range tests {
		tok := l.NextToken()
		if tok.Type != test.ExpectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, test.ExpectedType, tok.Type)
		}
		if tok.Literal != test.ExpectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, test.ExpectedLiteral, tok.Literal)
		}
	}
}

func TestSourceTokens(t *testing.T) {
	inp := `let five = 5;
let ten = 10;`
	tests := []struct {
		ExpectedType    token.TokenType
		ExpectedLiteral string
	}{
		{ExpectedType: token.LET, ExpectedLiteral: "let"},
		{ExpectedType: token.IDENT, ExpectedLiteral: "five"},
		{ExpectedType: token.ASSIGN, ExpectedLiteral: "="},
		{ExpectedType: token.INT, ExpectedLiteral: "5"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},
		{ExpectedType: token.LET, ExpectedLiteral: "let"},
		{ExpectedType: token.IDENT, ExpectedLiteral: "ten"},
		{ExpectedType: token.ASSIGN, ExpectedLiteral: "="},
		{ExpectedType: token.INT, ExpectedLiteral: "10"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},
	}
	l := New(inp)
	for i, test := range tests {
		tok := l.NextToken()
		if tok.Type != test.ExpectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, test.ExpectedType, tok.Type)
		}
		if tok.Literal != test.ExpectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, test.ExpectedLiteral, tok.Literal)
		}
	}
}

func TestSpecialChars(t *testing.T) {
	input := `!-/*5;
5 < 10 > 5;
let sum = fn(x, y) {
  x + y;
};`
	tests := []struct {
		ExpectedType    token.TokenType
		ExpectedLiteral string
	}{
		{ExpectedType: token.BANG, ExpectedLiteral: "!"},
		{ExpectedType: token.MINUS, ExpectedLiteral: "-"},
		{ExpectedType: token.SLASH, ExpectedLiteral: "/"},
		{ExpectedType: token.ASTERISK, ExpectedLiteral: "*"},
		{ExpectedType: token.INT, ExpectedLiteral: "5"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},
		{ExpectedType: token.INT, ExpectedLiteral: "5"},
		{ExpectedType: token.LT, ExpectedLiteral: "<"},
		{ExpectedType: token.INT, ExpectedLiteral: "10"},
		{ExpectedType: token.GT, ExpectedLiteral: ">"},
		{ExpectedType: token.INT, ExpectedLiteral: "5"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},
		{ExpectedType: token.LET, ExpectedLiteral: "let"},
		{ExpectedType: token.IDENT, ExpectedLiteral: "sum"},
		{ExpectedType: token.ASSIGN, ExpectedLiteral: "="},
		{ExpectedType: token.FUNCTION, ExpectedLiteral: "fn"},
		{ExpectedType: token.LPAREN, ExpectedLiteral: "("},
		{ExpectedType: token.IDENT, ExpectedLiteral: "x"},
		{ExpectedType: token.COMMA, ExpectedLiteral: ","},
		{ExpectedType: token.IDENT, ExpectedLiteral: "y"},
		{ExpectedType: token.RPAREN, ExpectedLiteral: ")"},
		{ExpectedType: token.LBRACE, ExpectedLiteral: "{"},
		{ExpectedType: token.IDENT, ExpectedLiteral: "x"},
		{ExpectedType: token.PLUS, ExpectedLiteral: "+"},
		{ExpectedType: token.IDENT, ExpectedLiteral: "y"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},
		{ExpectedType: token.RBRACE, ExpectedLiteral: "}"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},
	}
	l := New(input)
	for i, test := range tests {
		tok := l.NextToken()
		if tok.Type != test.ExpectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, test.ExpectedType, tok.Type)
		}
		if tok.Literal != test.ExpectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, test.ExpectedLiteral, tok.Literal)
		}
	}
}

func TestConditionals(t *testing.T) {
	input := `if (5 < 10) {
  return true;
} else {
  return false;
}`
	tests := []struct {
		ExpectedType    token.TokenType
		ExpectedLiteral string
	}{
		{ExpectedType: token.IF, ExpectedLiteral: "if"},
		{ExpectedType: token.LPAREN, ExpectedLiteral: "("},
		{ExpectedType: token.INT, ExpectedLiteral: "5"},
		{ExpectedType: token.LT, ExpectedLiteral: "<"},
		{ExpectedType: token.INT, ExpectedLiteral: "10"},
		{ExpectedType: token.RPAREN, ExpectedLiteral: ")"},
		{ExpectedType: token.LBRACE, ExpectedLiteral: "{"},
		{ExpectedType: token.RETURN, ExpectedLiteral: "return"},
		{ExpectedType: token.TRUE, ExpectedLiteral: "true"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},
		{ExpectedType: token.RBRACE, ExpectedLiteral: "}"},
		{ExpectedType: token.ELSE, ExpectedLiteral: "else"},
		{ExpectedType: token.LBRACE, ExpectedLiteral: "{"},
		{ExpectedType: token.RETURN, ExpectedLiteral: "return"},
		{ExpectedType: token.FALSE, ExpectedLiteral: "false"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},
		{ExpectedType: token.RBRACE, ExpectedLiteral: "}"},
	}
	l := New(input)
	for i, test := range tests {
		tok := l.NextToken()
		if tok.Type != test.ExpectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, test.ExpectedType, tok.Type)
		}
		if tok.Literal != test.ExpectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, test.ExpectedLiteral, tok.Literal)
		}
	}
}

func TestTwoCharTokens(t *testing.T) {
	input := `10 == 10;
10 != 9;`
	tests := []struct {
		ExpectedType    token.TokenType
		ExpectedLiteral string
	}{
		{ExpectedType: token.INT, ExpectedLiteral: "10"},
		{ExpectedType: token.EQL, ExpectedLiteral: "=="},
		{ExpectedType: token.INT, ExpectedLiteral: "10"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},
		{ExpectedType: token.INT, ExpectedLiteral: "10"},
		{ExpectedType: token.NEQ, ExpectedLiteral: "!="},
		{ExpectedType: token.INT, ExpectedLiteral: "9"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},
	}
	l := New(input)
	for i, test := range tests {
		tok := l.NextToken()
		if tok.Type != test.ExpectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, test.ExpectedType, tok.Type)
		}
		if tok.Literal != test.ExpectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, test.ExpectedLiteral, tok.Literal)
		}
	}
}
