package parser

import (
	"fmt"
	"strconv"
	"strings"

	"6502-asm/ast"
	"6502-asm/lexer"
	"6502-asm/token"
)

// Parser generates AST from tokens produced by given lexer.
type Parser struct {
	l     *lexer.Lexer
	token token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}

	p.advance()

	return p
}

func (p *Parser) advance() {
	p.token = p.l.Next()
}

// ParseProgram parses the tokens producing an ast.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	p.ignoreNewLines()

	for !p.match(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		if !p.check(token.EOF) {
			p.consume(token.NEWLINE, "Expected a new line")
			p.ignoreNewLines()
		}
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	if p.check(token.OPCODE) {
		return p.parseOpcode()
	}

	return nil
}

func (p *Parser) parseOpcode() *ast.Opcode {
	stmt := &ast.Opcode{
		Token:    p.token,
		Operands: []ast.Expression{},
	}
	p.advance()

	stmt.Operands = p.parseOperands()

	return stmt
}

func (p *Parser) parseOperands() []ast.Expression {
	operands := []ast.Expression{}

	for !p.check(token.EOF) && !p.check(token.NEWLINE) {
		operands = append(operands, p.parseNumber())
		p.advance()
	}

	return operands
}

func (p *Parser) parseNumber() *ast.NumberLiteral {
	var value int64
	var err error

	if strings.HasPrefix(p.token.Literal, "0x") {
		value, err = strconv.ParseInt(strings.TrimPrefix(p.token.Literal, "0x"), 16, 8)
	} else {
		value, err = strconv.ParseInt(p.token.Literal, 10, 8)
	}

	if err != nil {
		panic(err)
	}

	return &ast.NumberLiteral{
		Token: p.token,
		Value: int8(value),
	}
}

// ignoreNewLines consumes unused new line tokens.
func (p *Parser) ignoreNewLines() {
	for p.match(token.NEWLINE) {
	}
}

// check checks if current token is of the given type.
func (p *Parser) check(t token.Type) bool {
	return p.token.Type == t
}

// match checks if the next token is of the given type.
func (p *Parser) match(t token.Type) bool {
	if p.token.Type == t {
		p.advance()
		return true
	}
	return false
}

// consume advances parser only if the next token is of the given type.
// Returns the consumed token.
// If the next token is not of the specified type throws an error with given message.
func (p *Parser) consume(t token.Type, message string) (token.Token, error) {
	if p.token.Type == t {
		p.advance()
		return p.token, nil
	}

	return token.Token{}, fmt.Errorf(message)
}
