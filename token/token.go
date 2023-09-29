package token

const (
	SEMICOLON = ";"
	COMMA     = ","
	HEX       = "x"
	NUMBER    = "NUMBER"
	LINE      = "LINE"

	OP = "OP"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

// Token represents a single token produced by lexer.Lexer.
type Token struct {
	Type    string
	Literal string
}

// FromByte creates a new token with given token type and char literal.
func FromByte(tokenType string, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}
