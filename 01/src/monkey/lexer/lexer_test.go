package lexer

import (
	"testing"
	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	} {
		{TOKEN.ASSIGN, '='},
		{TOKEN.PLUS, '+'},
		{TOKEN.LPAREN, '('},
		{TOKEN.RPAREN, ')'},
		{TOKEN.LBRACE, '{'},
		{TOKEN.RBRACE, '}'},
		{TOKEN.COMMA, ','},
		{TOKEN.SEMICOLON, ';'},
		{TOKEN.EOF, ''},
	} 

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}