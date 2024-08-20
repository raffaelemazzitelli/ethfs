package evm

import "fmt"

type EVM struct {
	stop_flag bool
	pc        int
	program   []byte
	gas       int
	value     int
	calldata  []byte
}

type OpcodeHandler struct {
	Mnemonic string
	Handler  func(*EVM) error
}

func NewEVM(program []byte, gas int, value int, calldata []byte) *EVM {
	return &EVM{
		program:  program,
		gas:      gas,
		value:    value,
		calldata: calldata,

		pc: 0,

		stop_flag: false,
	}
}

func (evm *EVM) ShallExecuteNextOp() bool {
	if evm.stop_flag {
		return false
	}
	if evm.pc > len(evm.program)-1 {
		return false
	}

	return true
}
func (evm *EVM) Run() error {
	opcodeTable := evm.initOpcodeTable()

	for evm.ShallExecuteNextOp() {
		op := evm.program[evm.pc]
		handler, exists := opcodeTable[op]
		if !exists {
			evm.defaultHandler(op)
		} else if err := handler.Handler(evm); err != nil {
			return err
		}
	}
	return nil
}

func (evm *EVM) initOpcodeTable() map[byte]OpcodeHandler {
	return map[byte]OpcodeHandler{
		0x00: {"STOP", (*EVM).opStop},
	}
}

func (evm *EVM) defaultHandler(op byte) error {
	fmt.Printf("opcode 0x%x (%s) is not implemented\n", op, "UNKNOWN")
	evm.pc++
	return nil
}
