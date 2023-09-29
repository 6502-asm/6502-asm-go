package main

import (
	"6502-asm/lexer"
	"6502-asm/token"
	"fmt"
)

func main() {
	l := lexer.New("LDAI 5\n" +
		"LDBI 4\n" +
		"SUM\n" +
		"STA 5\n" +
		"HLT",
	)

	for {
		t := l.NextToken()
		fmt.Println("%v", t)

		if t.Type == token.EOF {
			break
		}
	}
}
