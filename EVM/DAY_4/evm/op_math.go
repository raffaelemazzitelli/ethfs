package evm

import (
	"fmt"

	"github.com/holiman/uint256"
)

func (evm *EVM) opAdd() (int, error) {
	//unsigned operation
	// https://github.com/ethereum/go-ethereum/blob/ea3b5095f439d63e35c8c37941836f4815fb380a/core/vm/instructions.go#L30
	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}
	b, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	evm.stack.PushInt(a.Add(a, b))
	return 0, nil
}

func (evm *EVM) opMul() (int, error) {
	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	b, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	evm.stack.PushInt(a.Mul(a, b))
	return 0, nil
}

func (evm *EVM) opSub() (int, error) {
	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	b, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	evm.stack.PushInt(a.Sub(a, b))
	return 0, nil
}

func (evm *EVM) opDiv() (int, error) {
	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	b, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	// Division by 0 are not directly handled by the EVM. It returns 0 in that case.
	if b.Sign() == 0 {
		fmt.Printf("\nSECURITY: attempt to divide by 0\n")
		evm.stack.PushInt(new(uint256.Int))
	} else {
		evm.stack.PushInt(a.Div(a, b))
	}

	return 0, nil
}

func (evm *EVM) opSDiv() (int, error) {
	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	b, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	// Division by 0 are not directly handled by the EVM. It returns 0 in that case.
	if b.Sign() == 0 {
		fmt.Printf("\nSECURITY: attempt to divide by 0\n")
		evm.stack.PushInt(new(uint256.Int))
	} else {
		evm.stack.PushInt(a.SDiv(a, b))
	}

	return 0, nil
}

func (evm *EVM) opMod() (int, error) {
	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	b, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	evm.stack.PushInt(a.Mod(a, b))

	return 0, nil
}

func (evm *EVM) opSMod() (int, error) {
	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	b, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	evm.stack.PushInt(a.SMod(a, b))

	return 0, nil
}

func (evm *EVM) opAddMod() (int, error) {
	c, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	b, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	evm.stack.PushInt(a.AddMod(a, b, c))

	return 0, nil
}

func (evm *EVM) opMulMod() (int, error) {
	c, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	b, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	evm.stack.PushInt(a.MulMod(a, b, c))

	return 0, nil
}

func (evm *EVM) opExp() (int, error) {
	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	b, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	evm.stack.PushInt(a.Exp(a, b))

	return 50 * b.ByteLen(), nil
}

func (evm *EVM) opSignExtend() (int, error) {
	b, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	evm.stack.PushInt(a.ExtendSign(a, b))

	return 0, nil
}
