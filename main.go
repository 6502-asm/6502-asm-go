package main

import (
	"6502-asm/parser"
	"flag"
	"github.com/sanity-io/litter"
	"log"
	"os"

	"6502-asm/lexer"
)

func main() {
	lgr := log.New(os.Stderr, "", 0)

	filename := flag.String("filename", "program.asm", "Name of the source file")
	flag.Parse()

	source, err := os.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}

	l := lexer.New(string(source))
	p := parser.New(l)

	ast, errs := p.ParseProgram()

	if errs != nil {
		for _, err := range errs {
			lgr.Println(err)
		}
	} else {
		litter.Config.DisablePointerReplacement = true
		litter.Dump(ast)
	}
}
