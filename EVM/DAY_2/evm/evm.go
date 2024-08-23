package evm

import (
	"errors"
	"fmt"
)

type EVM struct {
	stop_flag bool
	pc        int
	program   []byte
	gas       int
	value     int
	calldata  []byte
	stack     Stack
}

func NewEVM(program []byte, gas int, value int, calldata []byte) *EVM {
	return &EVM{
		program:  program,
		gas:      gas,
		value:    value,
		calldata: calldata,

		pc: 0,

		stop_flag: false,

		stack: *NewStack(),
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
			return evm.defaultHandler(op)
		} else if err := handler.Handler(evm); err != nil {
			return err
		}

		evm.DecrementGas(handler.GasCost)

		if !evm.stop_flag {
			evm.pc++
		}

	}
	return nil
}

func (evm *EVM) DecrementGas(gas int) error {
	if evm.gas-gas < 0 {
		return errors.New("negative offset is invalid")
	}
	evm.gas -= gas
	return nil
}

func (evm *EVM) defaultHandler(op byte) error {
	return fmt.Errorf("opcode 0x%x (%s) is not implemented", op, "UNKNOWN")
}
