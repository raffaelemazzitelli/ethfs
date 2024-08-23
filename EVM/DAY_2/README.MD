# Goal
The goal for day two is to be able to use the stack as all the operation will require the stack operations

# What do we need
We need to implement the stack.
We need to take in consideration the GAS

# Instruction to implement
| Value | Mnemonic   | δ: stack input | α: stack output | Description                               | Min Gas |
|-------|------------|----------------|----------------|-------------------------------------------|---------|
| 0x50  | POP        | 1              | 0              | Remove item from stack                    | 2       |
| 0x5f  | PUSH0      | 0              | 1              | Place value 0 on stack                    | 2       |
| 0x60  | PUSH1      | 0              | 1              | Place 1 byte item on stack                | 3       |
| 0x61  | PUSH2      | 0              | 1              | Place 2 byte item on stack                | 3       |
| 0x62  | PUSH3      | 0              | 1              | Place 3 byte item on stack                | 3       |
| 0x63  | PUSH4      | 0              | 1              | Place 4 byte item on stack                | 3       |
| 0x64  | PUSH5      | 0              | 1              | Place 5 byte item on stack                | 3       |
| 0x65  | PUSH6      | 0              | 1              | Place 6 byte item on stack                | 3       |
| 0x66  | PUSH7      | 0              | 1              | Place 7 byte item on stack                | 3       |
| 0x67  | PUSH8      | 0              | 1              | Place 8 byte item on stack                | 3       |
| 0x68  | PUSH9      | 0              | 1              | Place 9 byte item on stack                | 3       |
| 0x69  | PUSH10     | 0              | 1              | Place 10 byte item on stack               | 3       |
| 0x6a  | PUSH11     | 0              | 1              | Place 11 byte item on stack               | 3       |
| 0x6b  | PUSH12     | 0              | 1              | Place 12 byte item on stack               | 3       |
| 0x6c  | PUSH13     | 0              | 1              | Place 13 byte item on stack               | 3       |
| 0x6d  | PUSH14     | 0              | 1              | Place 14 byte item on stack               | 3       |
| 0x6e  | PUSH15     | 0              | 1              | Place 15 byte item on stack               | 3       |
| 0x6f  | PUSH16     | 0              | 1              | Place 16 byte item on stack               | 3       |
| 0x70  | PUSH17     | 0              | 1              | Place 17 byte item on stack               | 3       |
| 0x71  | PUSH18     | 0              | 1              | Place 18 byte item on stack               | 3       |
| 0x72  | PUSH19     | 0              | 1              | Place 19 byte item on stack               | 3       |
| 0x73  | PUSH20     | 0              | 1              | Place 20 byte item on stack               | 3       |
| 0x74  | PUSH21     | 0              | 1              | Place 21 byte item on stack               | 3       |
| 0x75  | PUSH22     | 0              | 1              | Place 22 byte item on stack               | 3       |
| 0x76  | PUSH23     | 0              | 1              | Place 23 byte item on stack               | 3       |
| 0x77  | PUSH24     | 0              | 1              | Place 24 byte item on stack               | 3       |
| 0x78  | PUSH25     | 0              | 1              | Place 25 byte item on stack               | 3       |
| 0x79  | PUSH26     | 0              | 1              | Place 26 byte item on stack               | 3       |
| 0x7a  | PUSH27     | 0              | 1              | Place 27 byte item on stack               | 3       |
| 0x7b  | PUSH28     | 0              | 1              | Place 28 byte item on stack               | 3       |
| 0x7c  | PUSH29     | 0              | 1              | Place 29 byte item on stack               | 3       |
| 0x7d  | PUSH30     | 0              | 1              | Place 30 byte item on stack               | 3       |
| 0x7e  | PUSH31     | 0              | 1              | Place 31 byte item on stack               | 3       |
| 0x7f  | PUSH32     | 0              | 1              | Place 32 byte (full word) item on stack   | 3       |
| 0x80  | DUP1       | 1              | 2              | Duplicate 1st stack item                  | 3       |
| 0x81  | DUP2       | 2              | 3              | Duplicate 2nd stack item                  | 3       |
| 0x82  | DUP3       | 3              | 4              | Duplicate 3rd stack item                  | 3       |
| 0x83  | DUP4       | 4              | 5              | Duplicate 4th stack item                  | 3       |
| 0x84  | DUP5       | 5              | 6              | Duplicate 5th stack item                  | 3       |
| 0x85  | DUP6       | 6              | 7              | Duplicate 6th stack item                  | 3       |
| 0x86  | DUP7       | 7              | 8              | Duplicate 7th stack item                  | 3       |
| 0x87  | DUP8       | 8              | 9              | Duplicate 8th stack item                  | 3       |
| 0x88  | DUP9       | 9              | 10             | Duplicate 9th stack item                  | 3       |
| 0x89  | DUP10      | 10             | 11             | Duplicate 10th stack item                 | 3       |
| 0x8a  | DUP11      | 11             | 12             | Duplicate 11th stack item                 | 3       |
| 0x8b  | DUP12      | 12             | 13             | Duplicate 12th stack item                 | 3       |
| 0x8c  | DUP13      | 13             | 14             | Duplicate 13th stack item                 | 3       |
| 0x8d  | DUP14      | 14             | 15             | Duplicate 14th stack item                 | 3       |
| 0x8e  | DUP15      | 15             | 16             | Duplicate 15th stack item                 | 3       |
| 0x8f  | DUP16      | 16             | 17             | Duplicate 16th stack item                 | 3       |
| 0x90  | SWAP1      | 2              | 2              | Exchange 1st and 2nd stack items          | 3       |
| 0x91  | SWAP2      | 3              | 3              | Exchange 1st and 3rd stack items          | 3       |
| 0x92  | SWAP3      | 4              | 4              | Exchange 1st and 4th stack items          | 3       |
| 0x93  | SWAP4      | 5              | 5              | Exchange 1st and 5th stack items          | 3       |
| 0x94  | SWAP5      | 6              | 6              | Exchange 1st and 6th stack items          | 3       |
| 0x95  | SWAP6      | 7              | 7              | Exchange 1st and 7th stack items          | 3       |
| 0x96  | SWAP7      | 8              | 8              | Exchange 1st and 8th stack items          | 3       |
| 0x97  | SWAP8      | 9              | 9              | Exchange 1st and 9th stack items          | 3       |
| 0x98  | SWAP9      | 10             | 10             | Exchange 1st and 10th stack items         | 3       |
| 0x99  | SWAP10     | 11             | 11             | Exchange 1st and 11th stack items         | 3       |
| 0x9a  | SWAP11     | 12             | 12             | Exchange 1st and 12th stack items         | 3       |
| 0x9b  | SWAP12     | 13             | 13             | Exchange 1st and 13th stack items         | 3       |
| 0x9c  | SWAP13     | 14             | 14             | Exchange 1st and 14th stack items         | 3       |
| 0x9d  | SWAP14     | 15             | 15             | Exchange 1st and 15th stack items         | 3       |
| 0x9e  | SWAP15     | 16             | 16             | Exchange 1st and 16th stack items         | 3       |
| 0x9f  | SWAP16     | 17             | 17             | Exchange 1st and 17th stack items         | 3       |
