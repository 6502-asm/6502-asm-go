package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"6502-asm/lexer"
	"6502-asm/token"
)

func main() {
	// Read the source file
	filename := flag.String("filename", "main.asm", "Name of the source file")
	flag.Parse()

	source, err := os.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}

	// Run lexer
	l := lexer.New(string(source))

	var tokens []token.Token
	for {
		nextToken := l.NextToken()
		tokens = append(tokens, nextToken)

		if nextToken.Type == token.EOF {
			break
		}
	}

	fmt.Printf("%+v", tokens)
}
