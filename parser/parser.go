package parser

import (
	"strconv"
	"strings"

	"6502-asm/ast"
	"6502-asm/lexer"
	"6502-asm/token"
)

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
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	p.ignoreNewLines()

	for p.currToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		if !p.currIs(token.EOF) {
			p.consume(token.NEWLINE)
			p.ignoreNewLines()
		}
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.OPCODE:
		return p.parseOpcode()
	}

	return nil
}

func (p *Parser) parseOpcode() *ast.Opcode {
	stmt := &ast.Opcode{
		Token:    p.currToken,
		Operands: []ast.Expression{},
	}

	stmt.Operands = p.parseOperands()

	return stmt
}

func (p *Parser) parseOperands() []ast.Expression {
	operands := []ast.Expression{}

	p.nextToken()
	for !p.currIs(token.EOF) && !p.currIs(token.NEWLINE) {
		operands = append(operands, p.parseNumber())
		p.nextToken()
	}

	return operands
}

func (p *Parser) parseNumber() *ast.NumberLiteral {
	var value int64
	var err error

	if strings.HasPrefix(p.currToken.Literal, "0x") {
		value, err = strconv.ParseInt(strings.TrimPrefix(p.currToken.Literal, "0x"), 16, 8)
	} else {
		value, err = strconv.ParseInt(p.currToken.Literal, 10, 8)
	}

	if err != nil {
		panic(err)
	}

	return &ast.NumberLiteral{
		Token: p.currToken,
		Value: int8(value),
	}
}

// ignoreNewLines consumes unused new line tokens.
func (p *Parser) ignoreNewLines() {
	for p.consume(token.NEWLINE) {
	}
}

// consume advances parser only if the next token is of the given type.
func (p *Parser) consume(t token.Type) bool {
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
