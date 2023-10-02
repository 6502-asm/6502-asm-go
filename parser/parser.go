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
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

// New creates a new parser.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.Next()
}

// ParseProgram parses the tokens producing *ast.Program.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	// TODO: We might in fact want to rename IDENT into OPCODE.
	case token.IDENT:
		return p.parseOpcode()
	}
	return nil
}

func (p *Parser) parseOpcode() *ast.Opcode {
	stmt := &ast.Opcode{Token: p.curToken, Operands: []ast.Expression{}}
	if !p.peekIs(token.LINE) {
		stmt.Operands = p.parseOperands()
	}
	p.consume(token.LINE)
	return stmt
}

func (p *Parser) parseOperands() []ast.Expression {
	operands := []ast.Expression{}

	p.nextToken()
	for p.curToken.Type != token.LINE {
		operands = append(operands, p.parseNumber())
		p.nextToken()
	}

	return operands
}

func (p *Parser) parseNumber() *ast.NumberLiteral {
	var value int64
	var err error

	if strings.HasPrefix(p.curToken.Literal, "0X") {
		value, err = strconv.ParseInt(strings.TrimPrefix(p.curToken.Literal, "0x"), 16, 8)
		if err != nil {
			panic(err)
		}
	} else {
		value, err = strconv.ParseInt(p.curToken.Literal, 10, 8)
		if err != nil {
			panic(err)
		}
	}

	return &ast.NumberLiteral{
		Token: p.curToken,
		Value: int8(value),
	}
}

func (p *Parser) consume(t token.Type) {
	if p.check(t) {
		p.nextToken()
		return
	}

	panic(fmt.Sprint("Expected token type: ", t))
}

func (p *Parser) match(t token.Type) bool {
	if p.curToken.Type == t {
		p.nextToken()
		return true
	}

	return false
}

func (p *Parser) check(t token.Type) bool {
	if p.curToken.Type == token.EOF {
		return false
	}

	return true
}

func (p *Parser) peekIs(t token.Type) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	}

	return false
}
