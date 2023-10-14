package compiler

import (
	"6502-asm/ast"
)

type Error struct {
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

type Compiler struct {
	ast *ast.Program
}

func New(ast *ast.Program) *Compiler {
	return &Compiler{
		ast: ast,
	}
}

func Make(operand byte, operands ...byte) ([]byte, error) {
	arity, ok := arities[operand]
	if !ok {
		return nil, &Error{
			Message: "unknown operand",
		}
	}

	if len(operands) != arity {
		return nil, &Error{
			Message: "invalid operands count",
		}
	}

	return append([]byte{operand}, operands...), nil
}
