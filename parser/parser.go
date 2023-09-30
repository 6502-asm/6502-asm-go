package parser

import (
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
	stmt := &ast.Opcode{Token: p.curToken}

	// We want to scan all operands. It will be compiler's task to decide
	// if given opcode has the right amount of operands.
	for p.peekIs(token.NUMBER) {
		// TODO: Parse numbers
	}

	if !p.expectPeek(token.LINE) {
		return nil
	}

	return stmt
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
