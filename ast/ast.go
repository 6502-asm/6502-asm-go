package ast

import "6502-asm/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type Opcode struct {
	Token    token.Token
	Operands []Expression
}

func (o *Opcode) TokenLiteral() string {
	return o.Token.Literal
}

func (o *Opcode) statementNode() {}
