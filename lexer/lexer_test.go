package lexer

import (
	"testing"

	"6502-asm/token"
)

func TestLexer_NextToken(t *testing.T) {
	input := `LDAI 5 ; load 5 to A register
LDBI 4 ; load 5 to 4 register
SUM ; sum A and B registers
STA 0x05 ; do some magic shit
HLT ; end the program
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
