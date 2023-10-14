package parser

import (
	"6502-asm/lexer"
	"testing"
)

func TestParser_ParseProgram(t *testing.T) {
	input := `LDAI 0x05 ; comment
LDBI 4
SUM
STA 0x05
HLT
`

	l := lexer.New(input)
	p := New(l)
}
