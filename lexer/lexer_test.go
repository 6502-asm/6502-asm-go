package lexer

import (
	"testing"

	"6502-asm/token"
)

func TestLexer_NextToken(t *testing.T) {
	input := `LDAI 5;
LDBI 4;
SUM;
STA 0x05;
HLT;
`

	l := New(input)

	tests := []token.Token{
		// LDAI 5;
		{token.OP, "LDAI"},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.LINE, "\n"},
		// LDBI 4;
		{token.OP, "LDBI"},
		{token.NUMBER, "4"},
		{token.SEMICOLON, ";"},
		{token.LINE, "\n"},
		// SUM;
		{token.OP, "SUM"},
		{token.SEMICOLON, ";"},
		{token.LINE, "\n"},
		// STA 0x05;
		{token.OP, "STA"},
		{token.NUMBER, "0"},
		{token.HEX, "x"},
		{token.NUMBER, "05"},
		{token.SEMICOLON, ";"},
		{token.LINE, "\n"},
		// HLD;
		{token.OP, "HLT"},
		{token.SEMICOLON, ";"},
		{token.LINE, "\n"},
		{token.EOF, ""},
	}

	for i, testToken := range tests {
		outToken := l.NextToken()

		if outToken.Type != testToken.Type {
			t.Fatalf("tests[%d]: invalid Type: expected=%q, got=%q", i, testToken.Type, outToken.Type)
		}

		if outToken.Literal != testToken.Literal {
			t.Fatalf("tests[%d]: invalid Literal: expected=%q, got %q", i, testToken.Literal, outToken.Literal)
		}
	}
}
