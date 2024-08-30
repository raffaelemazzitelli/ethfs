package evm

type OpCodeConfig struct {
	Mnemonic    string
	StackInput  int
	StackOutput int
	GasCost     int
	Handler     func(*EVM) (int, error)
}

func (evm *EVM) initOpcodeTable() map[byte]OpCodeConfig {
	return map[byte]OpCodeConfig{
		0x00: {"STOP", 0, 0, 0, (*EVM).opStop},

		// Arithmetic Operations
		0x01: {"ADD", 2, 1, 3, (*EVM).opAdd},
		0x02: {"MUL", 2, 1, 5, (*EVM).opMul},
		0x03: {"SUB", 2, 1, 3, (*EVM).opSub},
		0x04: {"DIV", 2, 1, 5, (*EVM).opDiv},
		0x05: {"SDIV", 2, 1, 5, (*EVM).opSDiv},
		0x06: {"MOD", 2, 1, 5, (*EVM).opMod},
		0x07: {"SMOD", 2, 1, 5, (*EVM).opSMod},
		0x08: {"ADDMOD", 3, 1, 8, (*EVM).opAddMod},
		0x09: {"MULMOD", 3, 1, 8, (*EVM).opMulMod},
		0x0a: {"EXP", 2, 1, 10, (*EVM).opExp},
		0x0b: {"SIGNEXTEND", 2, 1, 5, (*EVM).opSignExtend},

		// Comparison Operations
		0x10: {"LT", 2, 1, 3, (*EVM).opLt},
		0x11: {"GT", 2, 1, 3, (*EVM).opGt},
		0x12: {"SLT", 2, 1, 3, (*EVM).opSlt},
		0x13: {"SGT", 2, 1, 3, (*EVM).opSgt},
		0x14: {"EQ", 2, 1, 3, (*EVM).opEq},
		0x15: {"ISZERO", 1, 1, 3, (*EVM).opIsZero},

		// Bitwise Operations
		0x16: {"AND", 2, 1, 3, (*EVM).opAnd},
		0x17: {"OR", 2, 1, 3, (*EVM).opOr},
		0x18: {"XOR", 2, 1, 3, (*EVM).opXor},
		0x19: {"NOT", 1, 1, 3, (*EVM).opNot},
		0x1A: {"BYTE", 2, 1, 3, (*EVM).opByte},

		// Shift Operations
		0x1B: {"SHL", 2, 1, 3, (*EVM).opShl},
		0x1C: {"SHR", 2, 1, 3, (*EVM).opShr},
		0x1D: {"SAR", 2, 1, 3, (*EVM).opSar},

		// Stack operations
		0x50: {"POP", 1, 0, 2, (*EVM).opPop},

		0x5f: {"PUSH0", 0, 1, 2, evm.opPush(0)},
		0x60: {"PUSH1", 0, 1, 3, evm.opPush(1)},
		0x61: {"PUSH2", 0, 1, 3, evm.opPush(2)},
		0x62: {"PUSH3", 0, 1, 3, evm.opPush(3)},
		0x63: {"PUSH4", 0, 1, 3, evm.opPush(4)},
		0x64: {"PUSH5", 0, 1, 3, evm.opPush(5)},
		0x65: {"PUSH6", 0, 1, 3, evm.opPush(6)},
		0x66: {"PUSH7", 0, 1, 3, evm.opPush(7)},
		0x67: {"PUSH8", 0, 1, 3, evm.opPush(8)},
		0x68: {"PUSH9", 0, 1, 3, evm.opPush(9)},
		0x69: {"PUSH10", 0, 1, 3, evm.opPush(10)},
		0x6a: {"PUSH11", 0, 1, 3, evm.opPush(11)},
		0x6b: {"PUSH12", 0, 1, 3, evm.opPush(12)},
		0x6c: {"PUSH13", 0, 1, 3, evm.opPush(13)},
		0x6d: {"PUSH14", 0, 1, 3, evm.opPush(14)},
		0x6e: {"PUSH15", 0, 1, 3, evm.opPush(15)},
		0x6f: {"PUSH16", 0, 1, 3, evm.opPush(16)},
		0x70: {"PUSH17", 0, 1, 3, evm.opPush(17)},
		0x71: {"PUSH18", 0, 1, 3, evm.opPush(18)},
		0x72: {"PUSH19", 0, 1, 3, evm.opPush(19)},
		0x73: {"PUSH20", 0, 1, 3, evm.opPush(20)},
		0x74: {"PUSH21", 0, 1, 3, evm.opPush(21)},
		0x75: {"PUSH22", 0, 1, 3, evm.opPush(22)},
		0x76: {"PUSH23", 0, 1, 3, evm.opPush(23)},
		0x77: {"PUSH24", 0, 1, 3, evm.opPush(24)},
		0x78: {"PUSH25", 0, 1, 3, evm.opPush(25)},
		0x79: {"PUSH26", 0, 1, 3, evm.opPush(26)},
		0x7a: {"PUSH27", 0, 1, 3, evm.opPush(27)},
		0x7b: {"PUSH28", 0, 1, 3, evm.opPush(28)},
		0x7c: {"PUSH29", 0, 1, 3, evm.opPush(29)},
		0x7d: {"PUSH30", 0, 1, 3, evm.opPush(30)},
		0x7e: {"PUSH31", 0, 1, 3, evm.opPush(31)},
		0x7f: {"PUSH32", 0, 1, 3, evm.opPush(32)},

		0x80: {"DUP1", 1, 2, 3, evm.opDup(1)},
		0x81: {"DUP2", 2, 3, 3, evm.opDup(2)},
		0x82: {"DUP3", 3, 4, 3, evm.opDup(3)},
		0x83: {"DUP4", 4, 5, 3, evm.opDup(4)},
		0x84: {"DUP5", 5, 6, 3, evm.opDup(5)},
		0x85: {"DUP6", 6, 7, 3, evm.opDup(6)},
		0x86: {"DUP7", 7, 8, 3, evm.opDup(7)},
		0x87: {"DUP8", 8, 9, 3, evm.opDup(8)},
		0x88: {"DUP9", 9, 10, 3, evm.opDup(9)},
		0x89: {"DUP10", 10, 11, 3, evm.opDup(10)},
		0x8a: {"DUP11", 11, 12, 3, evm.opDup(11)},
		0x8b: {"DUP12", 12, 13, 3, evm.opDup(12)},
		0x8c: {"DUP13", 13, 14, 3, evm.opDup(13)},
		0x8d: {"DUP14", 14, 15, 3, evm.opDup(14)},
		0x8e: {"DUP15", 15, 16, 3, evm.opDup(15)},
		0x8f: {"DUP16", 16, 17, 3, evm.opDup(16)},

		0x90: {"SWAP1", 2, 2, 3, evm.opSwap(1)},
		0x91: {"SWAP2", 3, 3, 3, evm.opSwap(2)},
		0x92: {"SWAP3", 4, 4, 3, evm.opSwap(3)},
		0x93: {"SWAP4", 5, 5, 3, evm.opSwap(4)},
		0x94: {"SWAP5", 6, 6, 3, evm.opSwap(5)},
		0x95: {"SWAP6", 7, 7, 3, evm.opSwap(6)},
		0x96: {"SWAP7", 8, 8, 3, evm.opSwap(7)},
		0x97: {"SWAP8", 9, 9, 3, evm.opSwap(8)},
		0x98: {"SWAP9", 10, 10, 3, evm.opSwap(9)},
		0x99: {"SWAP10", 11, 11, 3, evm.opSwap(10)},
		0x9a: {"SWAP11", 12, 12, 3, evm.opSwap(11)},
		0x9b: {"SWAP12", 13, 13, 3, evm.opSwap(12)},
		0x9c: {"SWAP13", 14, 14, 3, evm.opSwap(13)},
		0x9d: {"SWAP14", 15, 15, 3, evm.opSwap(14)},
		0x9e: {"SWAP15", 16, 16, 3, evm.opSwap(15)},
		0x9f: {"SWAP16", 17, 17, 3, evm.opSwap(16)},
	}
}
