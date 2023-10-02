package ast

import "6502-asm/token"

type Node interface {
	TokenLiteral() string
}

// Statement represents statement node in AST.
type Statement interface {
	Node
	statementNode()
}

// Expression represents expression node in AST.
type Expression interface {
	Node
	expressionNode()
}

// Program node
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// Opcode node
type Opcode struct {
	Token    token.Token
	Operands []Expression
}

func (o *Opcode) statementNode() {}
func (o *Opcode) TokenLiteral() string {
	return o.Token.Literal
}

type NumberLiteral struct {
	Token token.Token
	Value int8
}

func (bl *NumberLiteral) expressionNode() {}
func (bl *NumberLiteral) TokenLiteral() string {
	return bl.Token.Literal
}
