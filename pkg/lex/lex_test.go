package lex

import (
	"phrog/pkg/token"
	"testing"
)

func TestFetchNextToken(t *testing.T) {
	input := `assign five 5;
	assign ten 10;
	assign add function(x, y) {
		x + y;
	};
	
	assign result add(five, ten);
	
	!-/*5;
	5 < 10 > 5;

	result != 15;
	result == 15;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "assign"},
		{token.ID, "five"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.ASSIGN, "assign"},
		{token.ID, "ten"},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.ASSIGN, "assign"},
		{token.ID, "add"},
		{token.FUNCTION, "function"},
		{token.LPAREN, "("},
		{token.ID, "x"},
		{token.COMMA, ","},
		{token.ID, "y"},
		{token.RPAREN, ")"},
		{token.LCURLEY, "{"},
		{token.ID, "x"},
		{token.PLUS, "+"},
		{token.ID, "y"},
		{token.SEMICOLON, ";"},
		{token.RCURLEY, "}"},
		{token.SEMICOLON, ";"},
		{token.ASSIGN, "assign"},
		{token.ID, "result"},
		{token.ID, "add"},
		{token.LPAREN, "("},
		{token.ID, "five"},
		{token.COMMA, ","},
		{token.ID, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.NOT, "!"},
		{token.MINUS, "-"},
		{token.DIVIDE, "/"},
		{token.MULTIPLY, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.ID, "result"},
		{token.NEQ, "!="},
		{token.INT, "15"},
		{token.SEMICOLON, ";"},
		{token.ID, "result"},
		{token.EQ, "=="},
		{token.INT, "15"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LCURLEY, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RCURLEY, "}"},
		{token.ELSE, "else"},
		{token.LCURLEY, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RCURLEY, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		token := l.FetchNextToken()

		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, token.Type)
		}

		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, token.Literal)
		}
	}
}
