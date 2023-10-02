package token

type Type string

const (
	SEMICOLON = ";"
	COMMA     = ","

	NUMBER = "NUMBER"
	OPCODE = "OPCODE"

	LINE    = "LINE"
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

// Token represents a single token produced by lexer.Lexer.
type Token struct {
	Type    Type
	Literal string
}

// FromByte creates a new token with given token type and char literal.
func FromByte(tokenType Type, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}
