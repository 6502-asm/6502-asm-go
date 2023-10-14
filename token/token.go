package token

type Type string

const (
	COMMA      = ","
	COLON      = ":"
	NUMBER     = "NUMBER"
	IDENTIFIER = "IDENTIFIER"
	NEWLINE    = "NEWLINE"
	ILLEGAL    = "ILLEGAL"
	EOF        = "EOF"
)

// Token represents a single token produced by lexer.Lexer.
type Token struct {
	Type    Type
	Literal string
	Line    int
}

// FromByte creates a new token with given token type and char literal.
func FromByte(tokenType Type, ch byte, line int) Token {
	return Token{Type: tokenType, Literal: string(ch), Line: line}
}
