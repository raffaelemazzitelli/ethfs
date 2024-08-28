package evm

import (
	"fmt"

	"github.com/holiman/uint256"
)

func (evm *EVM) opAdd() error {
	//unsigned operation
	// https://github.com/ethereum/go-ethereum/blob/ea3b5095f439d63e35c8c37941836f4815fb380a/core/vm/instructions.go#L30
	a, err := evm.stack.PopInt()
	if err != nil {
		return err
	}
	b, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	evm.stack.PushInt(a.Add(a, b))
	return nil
}

func (evm *EVM) opMul() error {
	a, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	b, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	evm.stack.PushInt(a.Mul(a, b))
	return nil
}

func (evm *EVM) opSub() error {
	a, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	b, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	evm.stack.PushInt(a.Sub(a, b))
	return nil
}

func (evm *EVM) opDiv() error {
	a, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	b, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	// Division by 0 are not directly handled by the EVM. It returns 0 in that case.
	if b.Sign() == 0 {
		fmt.Printf("\nSECURITY: attempt to divide by 0\n")
		evm.stack.PushInt(new(uint256.Int))
	} else {
		evm.stack.PushInt(a.Div(a, b))
	}

	return nil
}

func (evm *EVM) opSDiv() error {
	a, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	b, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	// Division by 0 are not directly handled by the EVM. It returns 0 in that case.
	if b.Sign() == 0 {
		fmt.Printf("\nSECURITY: attempt to divide by 0\n")
		evm.stack.PushInt(new(uint256.Int))
	} else {
		evm.stack.PushInt(a.SDiv(a, b))
	}

	return nil
}

func (evm *EVM) opMod() error {
	a, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	b, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	evm.stack.PushInt(a.Mod(a, b))

	return nil
}

func (evm *EVM) opSMod() error {
	a, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	b, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	evm.stack.PushInt(a.SMod(a, b))

	return nil
}

func (evm *EVM) opAddMod() error {
	c, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	b, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	a, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	evm.stack.PushInt(a.AddMod(a, b, c))

	return nil
}

func (evm *EVM) opMulMod() error {
	c, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	b, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	a, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	evm.stack.PushInt(a.MulMod(a, b, c))

	return nil
}

func (evm *EVM) opExp() error {
	a, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	b, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	evm.stack.PushInt(a.Exp(a, b))

	return nil
}

func (evm *EVM) opSignExtend() error {
	b, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	a, err := evm.stack.PopInt()
	if err != nil {
		return err
	}

	evm.stack.PushInt(a.ExtendSign(a, b))

	return nil
}
