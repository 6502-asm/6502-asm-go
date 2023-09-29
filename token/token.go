package token

const (
	NUMBER    = "NUMBER"
	IMMEDIATE = "IMMEDIATE"
	ADDRES    = "ADDRESS"
	OP        = "OP"
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
)

// Token represents a single token produced by lexer.Lexer.
type Token struct {
	Type    string
	Literal string
}
