package lexer

import (
	"testing"

	"asm/token"
)

func TestLexer_Next(t *testing.T) {
	cases := []struct {
		Tag            string
		Source         string
		ExpectedTokens []token.Token
	}{
		{
			Tag: "operands, arguments and labels",
			Source: `LDXI 0x10
ZEROY
LOOP:
MOVAY
MOVBX
ADD
DECX
MOVYA
TESTX
JMPNZ LOOP
HLT
`,
			ExpectedTokens: []token.Token{
				{token.IDENTIFIER, "LDXI", 1},
				{token.NUMBER, "0x10", 1},
				{token.NEWLINE, "\n", 1},
				{token.IDENTIFIER, "ZEROY", 2},
				{token.NEWLINE, "\n", 2},
				{token.IDENTIFIER, "LOOP", 3},
				{token.COLON, ":", 3},
				{token.NEWLINE, "\n", 3},
				{token.IDENTIFIER, "MOVAY", 4},
				{token.NEWLINE, "\n", 4},
				{token.IDENTIFIER, "MOVBX", 5},
				{token.NEWLINE, "\n", 5},
				{token.IDENTIFIER, "ADD", 6},
				{token.NEWLINE, "\n", 6},
				{token.IDENTIFIER, "DECX", 7},
				{token.NEWLINE, "\n", 7},
				{token.IDENTIFIER, "MOVYA", 8},
				{token.NEWLINE, "\n", 8},
				{token.IDENTIFIER, "TESTX", 9},
				{token.NEWLINE, "\n", 9},
				{token.IDENTIFIER, "JMPNZ", 10},
				{token.IDENTIFIER, "LOOP", 10},
				{token.NEWLINE, "\n", 10},
				{token.IDENTIFIER, "HLT", 11},
				{token.NEWLINE, "\n", 11},
				{token.EOF, string(byte(0)), 12},
			},
		},
		{
			Tag: "operands, arguments and comments",
			Source: `LDAI 0x05 ; comment
LDBI 4
SUM
STA 0x05
HLT
`,
			ExpectedTokens: []token.Token{
				{token.IDENTIFIER, "LDAI", 1},
				{token.NUMBER, "0x05", 1},
				{token.NEWLINE, "\n", 1},
				{token.IDENTIFIER, "LDBI", 2},
				{token.NUMBER, "4", 2},
				{token.NEWLINE, "\n", 2},
				{token.IDENTIFIER, "SUM", 3},
				{token.NEWLINE, "\n", 3},
				{token.IDENTIFIER, "STA", 4},
				{token.NUMBER, "0x05", 4},
				{token.NEWLINE, "\n", 4},
				{token.IDENTIFIER, "HLT", 5},
				{token.NEWLINE, "\n", 5},
				{token.EOF, string(byte(0)), 6},
			},
		},
	}

	for _, test := range cases {
		l := New(test.Source)

		for i, testToken := range test.ExpectedTokens {
			outToken := l.Next()

			if outToken.Type != testToken.Type {
				t.Fatalf("%v: tests[%d]: invalid Type: expected=%q, got=%q\n%+v\n%+v", test.Tag, i, testToken.Type, outToken.Type, testToken, outToken)
			}

			if outToken.Literal != testToken.Literal {
				t.Fatalf("%v: tests[%d]: invalid Literal: expected=%q, got %q\n%+v\n%+v", test.Tag, i, testToken.Literal, outToken.Literal, testToken, outToken)
			}

			if outToken.Line != testToken.Line {
				t.Fatalf("%v: tests[%d]: invalid Line: expected=%q, got=%q\n%+v\n%+v", test.Tag, i, testToken.Line, outToken.Line, testToken, outToken)
			}
		}
	}

}
