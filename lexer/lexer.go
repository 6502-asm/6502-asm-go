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
	line    int
}

// New creates a new instance of Lexer.
func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1}
	l.advance()
	return l
}

// Next generates the next token from input source.
func (l *Lexer) Next() token.Token {
	var t token.Token

	l.skipWhitespace()
	l.skipComment()

	switch l.c {
	case '\n':
		t = token.FromByte(token.NEWLINE, l.c, l.line)
		l.line++
		break
	case ',':
		t = token.FromByte(token.COMMA, l.c, l.line)
		break
	case ':':
		t = token.FromByte(token.COLON, l.c, l.line)
		break
	case 0:
		t = token.FromByte(token.EOF, 0, l.line)
		break
	default:
		if isLetter(l.c) {
			t = token.Token{
				Type:    token.IDENTIFIER,
				Literal: l.readOpcode(),
				Line:    l.line,
			}
			return t
		} else if isDigit(l.c) {
			t = token.Token{
				Type:    token.NUMBER,
				Literal: l.readNumber(),
				Line:    l.line,
			}
			return t
		} else {
			t = token.FromByte(token.ILLEGAL, l.c, l.line)
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
	for l.c == ' ' || l.c == '\t' || l.c == '\r' {
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

func (l *Lexer) readOpcode() string {
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
	return 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z'
}

// TODO: We cannot do that because if isDigit will be used together with isLetter things can break. We might want to create isHexDigit function that would be used after initial check with isDigit.
func isDigit(c byte) bool {
	return '0' <= c && c <= '9' || 'A' <= c && c <= 'F' || c == 'x'
}
