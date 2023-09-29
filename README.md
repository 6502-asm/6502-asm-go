_6502-asm_ is an assembler for our school computer.

Example program that adds two numbers stored in registers `A` and `B`.
```
LDAI 5
LDBI 4
SUM
STA 0x05
HLT
```

Augmented Backus–Naur form
```
program ::= statement*
statement ::= opcode "\n"
opcode ::= IDENT arguments?
arguments ::= NUMBER | hex
hex ::= "0x" 2*hex-digit
hex-digit ::= %x30-39 | %x41-46
```
