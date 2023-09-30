package lexer

import (
	"testing"

	"6502-asm/token"
)

func TestLexer_NextToken(t *testing.T) {
	input := `LDAI 5 ; comment
LDBI 4
SUM
STA 0x05
HLT
`

	l := New(input)

	tests := []token.Token{
		// LDAI 5;
		{token.IDENT, "LDAI"},
		{token.NUMBER, "5"},
		{token.LINE, "\n"},
		// LDBI 4;
		{token.IDENT, "LDBI"},
		{token.NUMBER, "4"},
		{token.LINE, "\n"},
		// SUM;
		{token.IDENT, "SUM"},
		{token.LINE, "\n"},
		// STA 0x05;
		{token.IDENT, "STA"},
		{token.NUMBER, "0x05"},
		{token.LINE, "\n"},
		// HLD;
		{token.IDENT, "HLT"},
		{token.LINE, "\n"},
		{token.EOF, ""},
	}

	for i, testToken := range tests {
		outToken := l.Next()

		if outToken.Type != testToken.Type {
			t.Fatalf("tests[%d]: invalid Type: expected=%q, got=%q\n%+v\n%+v", i, testToken.Type, outToken.Type, testToken, outToken)
		}

		if outToken.Literal != testToken.Literal {
			t.Fatalf("tests[%d]: invalid Literal: expected=%q, got %q\n%+v\n%+v", i, testToken.Literal, outToken.Literal, testToken, outToken)
		}
	}
}
