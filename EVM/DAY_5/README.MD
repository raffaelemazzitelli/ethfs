# Goal
The goal for day is to add other bitwise operations, shift operations, and misc operations

# What do we need
In the previous DAY we forgot to also implement a revert mechanism that takes in consideration when there is not enough gas.
We may implement it now or later, but we need to keep it in mind.

# Instruction to implement
| Value | Mnemonic | δ: stack input | α: stack output | Description                          | Min Gas |
|-------|----------|----------------|----------------|--------------------------------------|---------|
| 0x16  | AND       | 2              | 1              | Bitwise AND operation. `a & b`           | 3       |
| 0x17  | OR        | 2              | 1              | Bitwise OR operation. `a | b`            | 3       |
| 0x18  | XOR       | 2              | 1              | Bitwise XOR operation. `a ^ b`           | 3       |
| 0x19  | NOT       | 1              | 1              | Bitwise NOT operation. `~a`              | 3       |
| 0x1A  | BYTE      | 2              | 1              | Retrieve single byte from word.          | 3       |
| 0x1B  | SHL       | 2              | 1              | Left shift operation. `value << shift`   | 3       |
| 0x1C  | SHR       | 2              | 1              | Logical right shift operation. `value >> shift` | 3   |
| 0x1D  | SAR       | 2              | 1              | Arithmetic (signed) right shift operation. `value >> shift` | 3 |
