# Goal
The goal for day two is to be able to execute comparison instructions

# What do we need
In the previous DAY we forgot to also implement a revert mechanism that takes in consideration when there is not enough gas.
We may implement it now or later, but we need to keep it in mind.

# Instruction to implement
| Value | Mnemonic | δ: stack input | α: stack output | Description                          | Min Gas |
|-------|----------|----------------|----------------|--------------------------------------|---------|
| 0x10  | LT       | 2              | 1              | Less-than comparison. `a < b`        | 3       |
| 0x11  | GT       | 2              | 1              | Greater-than comparison. `a > b`     | 3       |
| 0x12  | SLT      | 2              | 1              | Signed less-than comparison. `a < b` | 3       |
| 0x13  | SGT      | 2              | 1              | Signed greater-than comparison. `a > b` | 3    |
| 0x14  | EQ       | 2              | 1              | Equality comparison. `a == b`        | 3       |
| 0x15  | ISZERO   | 1              | 1              | Is-zero comparison. `a == 0`         | 3       |
