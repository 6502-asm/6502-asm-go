package lexer

import (
	"6502-asm/token"
)

// Lexer transforms the input source into slice of tokens.
type Lexer struct {
	input   string
	pos     int
	readPos int
	c       byte
}

// New creates a new instance of Lexer.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.advance()
	return l
}

// Next generates the next token from input source.
func (l *Lexer) Next() token.Token {
	var t token.Token

	l.skipWhitespace()
	l.skipComment()

	switch l.c {
	case '\n', '\r':
		t = token.FromByte(token.LINE, l.c)
		break
	case ',':
		t = token.FromByte(token.COMMA, l.c)
		break
	case 0:
		t.Literal = ""
		t.Type = token.EOF
		break
	default:
		if isLetter(l.c) {
			t.Literal = l.readIdent()
			t.Type = token.IDENT
			return t
		} else if isDigit(l.c) {
			t.Type = token.NUMBER
			t.Literal = l.readNumber()
			return t
		} else {
			t = token.FromByte(token.ILLEGAL, l.c)
		}
	}

	l.advance()
	return t
}

func (l *Lexer) advance() {
	if l.readPos >= len(l.input) {
		l.c = 0
	} else {
		l.c = l.input[l.readPos]
	}

	l.pos = l.readPos
	l.readPos += 1
}

func (l *Lexer) skipWhitespace() {
	for l.c == ' ' || l.c == '\t' {
		l.advance()
	}
}

func (l *Lexer) skipComment() {
	if l.c == ';' {
		for !l.isAtEnd() && l.c != '\n' {
			l.advance()
		}
	}
}

func (l *Lexer) readIdent() string {
	position := l.pos
	for isLetter(l.c) {
		l.advance()
	}
	return l.input[position:l.pos]
}

func (l *Lexer) readNumber() string {
	position := l.pos
	for isDigit(l.c) {
		l.advance()
	}
	return l.input[position:l.pos]
}

// peek gives a single char of lookahead
func (l *Lexer) peek() byte {
	if l.isAtEnd() {
		return 0
	}

	return l.input[l.readPos]
}

func (l *Lexer) isAtEnd() bool {
	return l.readPos >= len(l.input)
}

func isLetter(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

// TODO: We cannot do that because if isDigit will be used together with isLetter things can break. We might want to create isHexDigit function that would be used after initial check with isDigit.
func isDigit(c byte) bool {
	return '0' <= c && c <= '9' || 'A' <= c && c <= 'F' || c == 'x'
}
