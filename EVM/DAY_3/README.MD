# Goal
The goal for today is to implement math functionalities

# Instruction to implement
| Value | Mnemonic    | δ: stack input | α: stack output | Description                                      | Min Gas |
|-------|-------------|----------------|----------------|--------------------------------------------------|---------|
| 0x01  | ADD         | 2              | 1              | Addition operation `a + b`                       | 3       |
| 0x02  | MUL         | 2              | 1              | Multiplication operation `a * b`                 | 5       |
| 0x03  | SUB         | 2              | 1              | Subtraction operation `a - b`                    | 3       |
| 0x04  | DIV         | 2              | 1              | Integer division operation `a // b`              | 5       |
| 0x05  | SDIV        | 2              | 1              | Signed integer division (truncated) `a // b`     | 5       |
| 0x06  | MOD         | 2              | 1              | Modulo remainder operation `a % b`               | 5       |
| 0x07  | SMOD        | 2              | 1              | Signed modulo remainder operation `a % b`        | 5       |
| 0x08  | ADDMOD      | 3              | 1              | Modulo addition operation `(a + b) % N`          | 8       |
| 0x09  | MULMOD      | 3              | 1              | Modulo multiplication operation `(a * b) % N`    | 8       |
| 0x0a  | EXP         | 2              | 1              | Exponential operation `a ** exponent`            | 10      |
| 0x0b  | SIGNEXTEND  | 2              | 1              | Extend length of two’s complement signed integer | 5       |
