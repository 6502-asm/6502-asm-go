package lexer

import (
	"testing"

	"6502-asm/token"
)

func TestLexer_NextToken(t *testing.T) {
	input := `LDAI 0x05 ; comment
LDBI 4
SUM
STA 0x05
HLT
`

	l := New(input)

	tests := []token.Token{
		{token.OPCODE, "LDAI"},
		{token.NUMBER, "0x05"},
		{token.NEWLINE, "\n"},
		{token.OPCODE, "LDBI"},
		{token.NUMBER, "4"},
		{token.NEWLINE, "\n"},
		{token.OPCODE, "SUM"},
		{token.NEWLINE, "\n"},
		{token.OPCODE, "STA"},
		{token.NUMBER, "0x05"},
		{token.NEWLINE, "\n"},
		{token.OPCODE, "HLT"},
		{token.NEWLINE, "\n"},
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
