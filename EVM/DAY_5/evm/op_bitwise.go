package evm

func (evm *EVM) opAnd() (int, error) {
	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}
	b, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	evm.stack.PushInt(a.And(a, b))
	return 0, nil
}

func (evm *EVM) opOr() (int, error) {
	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}
	b, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	evm.stack.PushInt(a.Or(a, b))
	return 0, nil
}

func (evm *EVM) opXor() (int, error) {
	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}
	b, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	evm.stack.PushInt(a.Xor(a, b))
	return 0, nil
}

func (evm *EVM) opNot() (int, error) {
	a, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	evm.stack.PushInt(a.Not(a))
	return 0, nil
}

func (evm *EVM) opByte() (int, error) {
	offset, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}
	value, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	result := value.Byte(offset)
	// fmt.Printf("value:%x,offset:%x,result:%x", value, offset, result)
	evm.stack.PushInt(result)
	return 0, nil
}
