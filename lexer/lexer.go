package lexer

import (
	"6502-asm/token"
)

// Lexer transforms the input source into slice of tokens.
type Lexer struct {
	input   string
	pos     int  // current position in input
	readPos int  // current reading position
	ch      byte // ch is the currently processed char
}

// New creates a new instance of Lexer.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readCh()
	return l
}

// readCh reads the char located at readPosition.
// In case all the chars have been consumed '\0' is being loaded.
func (l *Lexer) readCh() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}

	l.pos = l.readPos
	l.readPos += 1
}

// NextToken generates the next token from input source.
func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.shipWhitespace()

	switch l.ch {
	case 'r':
		t = newToken(token.IMMEDIATE, l.ch)
	case '#':
		t = newToken(token.ADDRES, l.ch)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.ch) {
			t.Literal = l.readOp()
			t.Type = token.OP
			return t
		} else if isDigit(l.ch) {
			t.Type = token.NUMBER
			t.Literal = l.readNumber()
		} else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readCh()
	return t
}

// skipWhitespace ignores whitespace characters.
func (l *Lexer) shipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readCh()
	}
}

// readOp reads the opcode.
func (l *Lexer) readOp() string {
	position := l.pos
	for isLetter(l.ch) {
		l.readCh()
	}
	return l.input[position:l.pos]
}

// readNumber reads number.
func (l *Lexer) readNumber() string {
	position := l.pos
	for isDigit(l.ch) {
		l.readCh()
	}
	return l.input[position:l.pos]
}

// newToken creates a new token.
func newToken(tokenType string, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// isLetter determines if the given ch should be treated as a letter.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit determines if the given ch should be treated as a digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
