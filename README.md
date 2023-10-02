_6502-asm_ is an assembler for our school computer's assembly.

Example program that adds two numbers stored in registers `A` and `B`.
```
LDAI 0x0F
LDBI 4
SUM
STA 0x05
HLT
```

```ts
&ast.Program{
  Statements: []ast.Statement{
    &ast.Opcode{
      Token: token.Token{
        Type: "OPCODE",
        Literal: "LDAI",
      },
      Operands: []ast.Expression{
        &ast.NumberLiteral{
          Token: token.Token{
            Type: "NUMBER",
            Literal: "0x0F",
          },
          Value: 15,
        },
      },
    },
    &ast.Opcode{
      Token: token.Token{
        Type: "OPCODE",
        Literal: "LDBI",
      },
      Operands: []ast.Expression{
        &ast.NumberLiteral{
          Token: token.Token{
            Type: "NUMBER",
            Literal: "4",
          },
          Value: 4,
        },
      },
    },
    &ast.Opcode{
      Token: token.Token{
        Type: "OPCODE",
        Literal: "SUM",
      },
      Operands: []ast.Expression{}, // p0
    },
    &ast.Opcode{
      Token: token.Token{
        Type: "OPCODE",
        Literal: "STA",
      },
      Operands: []ast.Expression{
        &ast.NumberLiteral{
          Token: token.Token{
            Type: "NUMBER",
            Literal: "0x05",
          },
          Value: 5,
        },
      },
    },
    &ast.Opcode{
      Token: token.Token{
        Type: "OPCODE",
        Literal: "HLT",
      },
      Operands: p0,
    },
  },
}
```


Augmented Backus–Naur form
```
program   ::= statement*
statement ::= opcode "\n"
opcode    ::= IDENT operands?
operands  ::= NUMBER | hex
hex       ::= "0x" 2*hex-digit
hex-digit ::= %x30-39 | %x41-46
```
