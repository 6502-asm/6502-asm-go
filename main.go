package main

import (
	"6502-asm/ast"
	"6502-asm/parser"
	"flag"
	"fmt"
	"log"
	"os"

	"6502-asm/lexer"
)

func main() {
	// Read the source file
	filename := flag.String("filename", "program.asm", "Name of the source file")
	flag.Parse()

	source, err := os.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}

	// Parse
	l := lexer.New(string(source))
	p := parser.New(l)

	program := p.ParseProgram()
	fmt.Printf("%+v", program.Statements[0].(*ast.Opcode).Operands[0])
}
