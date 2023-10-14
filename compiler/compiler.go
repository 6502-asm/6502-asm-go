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
	Instructions []byte
	// TODO: What is the biggest possible offset our computer can handle?
	Labels map[string]byte
}

func New() *Compiler {
	return &Compiler{
		Instructions: []byte{},
		Labels:       map[string]byte{},
	}
}

func (c *Compiler) Compile(node ast.Node) error {
	switch node := node.(type) {
	case *ast.Program:
		for _, s := range node.Statements {
			err := c.Compile(s)
			if err != nil {
				return err
			}
		}

	case *ast.Opcode:
		definition, ok := definitions[node.Token.Literal]
		if !ok {
			return &Error{
				Message: "unknown opcode",
			}
		}

		if definition.Arity != len(node.Operands) {
			return &Error{
				Message: "invalid number of arguments",
			}
		}

		var operands []byte
		for _, operand := range node.Operands {
			switch operand := operand.(type) {
			case *ast.NumberLiteral:
				operands = append(operands, operand.Value)
			case *ast.LabelRef:
				addr, ok := c.Labels[operand.Token.Literal]
				if !ok {
					return &Error{
						Message: "label does not exist",
					}
				}
				operands = append(operands, addr)
			default:
				return &Error{
					Message: "expected number literal",
				}
			}
		}

		c.Instructions = append(c.Instructions, append([]byte{definition.Opcode}, operands...)...)

	case *ast.Label:
		c.Labels[node.Token.Literal] = byte(len(c.Instructions))
	}

	return nil
}
