_asm-go_ is an assembler for our school computer's assembly.

Augmented Backus–Naur form
```
program   ::= statement*
statement ::= (opcode | label) "\n"

label     ::= IDENT ":"

opcode    ::= IDENT operands?
operands  ::= NUMBER | hex | IDENT

hex       ::= "0x" 2*hex-digit
hex-digit ::= %x30-39 | %x41-46
```
