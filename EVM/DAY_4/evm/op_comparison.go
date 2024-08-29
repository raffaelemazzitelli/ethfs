package evm

import "github.com/holiman/uint256"

func (evm *EVM) opLt() (int, error) {
	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}
	b, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	result := new(uint256.Int)
	if a.Lt(b) {
		result.SetOne()
	} else {
		result.Clear()
	}
	evm.stack.PushInt(result)
	return 0, nil
}

func (evm *EVM) opGt() (int, error) {
	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}
	b, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	result := new(uint256.Int)
	if a.Gt(b) {
		result.SetOne()
	} else {
		result.Clear()
	}
	evm.stack.PushInt(result)
	return 0, nil
}

func (evm *EVM) opSlt() (int, error) {
	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}
	b, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	result := new(uint256.Int)
	if a.Slt(b) {
		result.SetOne()
	} else {
		result.Clear()
	}
	evm.stack.PushInt(result)
	return 0, nil

}

func (evm *EVM) opSgt() (int, error) {
	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}
	b, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	result := new(uint256.Int)
	if a.Sgt(b) {
		result.SetOne()
	} else {
		result.Clear()
	}
	evm.stack.PushInt(result)
	return 0, nil
}

func (evm *EVM) opEq() (int, error) {
	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}
	b, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	result := new(uint256.Int)
	if a.Eq(b) {
		result.SetOne()
	} else {
		result.Clear()
	}
	evm.stack.PushInt(result)
	return 0, nil
}

func (evm *EVM) opIsZero() (int, error) {
	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}
	result := new(uint256.Int)
	if a.IsZero() {
		result.SetOne()
	} else {
		result.Clear()
	}
	evm.stack.PushInt(result)
	return 0, nil
}
