package parser

import (
	"fmt"
	"strconv"
	"strings"

	"6502-asm/ast"
	"6502-asm/lexer"
	"6502-asm/token"
)

// Error represents error reported by the parser in case of invalid token stream.
type Error struct {
	Token   token.Token
	Message string
	Cause   error
}

func (pe *Error) Error() string {
	return fmt.Sprintf("[line %v] Error at '%v':\n\t%v", pe.Token.Line, pe.Token.Literal, pe.Message)
}

func (pe *Error) Unwrap() error {
	return pe.Cause
}

// Parser generates AST from tokens produced by given lexer.
type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	peekToken token.Token
}

// New creates a new parser.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

// nextToken advances the parser tokens.
func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.Next()
}

// ParseProgram parses the tokens producing an ast.
func (p *Parser) ParseProgram() (*ast.Program, []error) {
	var errors []error

	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	p.ignoreNewLines()

	for p.currToken.Type != token.EOF {
		stmt, err := p.parseStatement()

		if err != nil {
			errors = append(errors, err)
			p.nextToken()
		}

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		if !p.currIs(token.EOF) {
			p.match(token.NEWLINE)
			p.ignoreNewLines()
		}
	}

	return program, errors
}

func (p *Parser) parseStatement() (ast.Statement, error) {
	if p.peekIs(token.COLON) {
		return p.parseLabel()
	}

	return p.parseOpcode()
}

func (p *Parser) parseLabel() (*ast.Label, error) {
	ident, err := p.consume(token.IDENTIFIER, "expected label identifier")
	if err != nil {
		return nil, err
	}

	p.nextToken()

	return &ast.Label{Token: ident}, nil
}

// parseOpcode parses the token as an opcode.
func (p *Parser) parseOpcode() (*ast.Opcode, error) {
	ident, err := p.consume(token.IDENTIFIER, "expected opcode identifier")
	if err != nil {
		return nil, err
	}

	opcode := &ast.Opcode{
		Token:    ident,
		Operands: []ast.Expression{},
	}

	if !p.currIs(token.NEWLINE) {
		operands, err := p.parseOperands()
		if err != nil {
			return nil, err
		}
		opcode.Operands = operands
	}

	return opcode, nil
}

// parseOperands is a helper for parsing a list of operands.
func (p *Parser) parseOperands() ([]ast.Expression, error) {
	var operands []ast.Expression

	for !p.currIs(token.EOF) && !p.currIs(token.NEWLINE) {
		var operand ast.Expression

		if p.currIs(token.NUMBER) {
			var err error
			operand, err = p.parseNumber()
			if err != nil {
				return nil, err
			}
		} else if p.currIs(token.IDENTIFIER) {
			operand = &ast.LabelRef{Token: p.currToken}
		}

		if operand == nil {
			return nil, &Error{
				Token:   p.currToken,
				Message: "expected label reference of number literal",
			}
		}

		operands = append(operands, operand)
		p.nextToken()
	}

	return operands, nil
}

// parseNumber parses the token as a number.
func (p *Parser) parseNumber() (*ast.NumberLiteral, error) {
	var value int64
	var err error

	if strings.HasPrefix(p.currToken.Literal, "0x") {
		value, err = strconv.ParseInt(strings.TrimPrefix(p.currToken.Literal, "0x"), 16, 8)
	} else {
		value, err = strconv.ParseInt(p.currToken.Literal, 10, 8)
	}

	if err != nil {
		return nil, &Error{
			Token:   p.currToken,
			Message: "expected a number expression",
			Cause:   err,
		}
	}

	return &ast.NumberLiteral{
		Token: p.currToken,
		Value: byte(value),
	}, nil
}

// ignoreNewLines consumes unused new line tokens.
func (p *Parser) ignoreNewLines() {
	for p.match(token.NEWLINE) {
	}
}

// consume checks if current token is of the given type. If so returns the token and advances.
// Otherwise, returns an Error with given message.
func (p *Parser) consume(t token.Type, message string) (token.Token, error) {
	if p.currIs(t) {
		previous := p.currToken
		p.nextToken()
		return previous, nil
	}

	return token.Token{}, &Error{
		Token:   p.currToken,
		Message: message,
	}
}

// match advances parser only if the current token is of the given type.
func (p *Parser) match(t token.Type) bool {
	if p.currToken.Type == t {
		p.nextToken()
		return true
	}

	return false
}

// currIs checks if current token has the given token type.
func (p *Parser) currIs(t token.Type) bool {
	return p.currToken.Type == t
}

// peekIs checks if peek token has the given token type.
func (p *Parser) peekIs(t token.Type) bool {
	return p.peekToken.Type == t
}
