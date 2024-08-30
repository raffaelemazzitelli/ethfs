package evm

func (evm *EVM) opShl() (int, error) {
	shift, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}
	value, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	if shift.LtUint64(256) {
		value.Lsh(value, uint(shift.Uint64()))
	} else {
		value.Clear()
	}

	evm.stack.PushInt(value)
	return 0, nil
}

func (evm *EVM) opShr() (int, error) {
	shift, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}
	value, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	if shift.LtUint64(256) {
		value.Rsh(value, uint(shift.Uint64()))
	} else {
		value.Clear()
	}

	evm.stack.PushInt(value)
	return 0, nil
}

func (evm *EVM) opSar() (int, error) {
	shift, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}
	value, err := evm.stack.PopInt()
	if err != nil {
		return 0, err
	}

	if shift.GtUint64(256) {
		if value.Sign() >= 0 {
			value.Clear()
		} else {
			value.SetAllOne()
		}
		return 0, nil
	}
	n := uint(shift.Uint64())
	value.SRsh(value, n)

	evm.stack.PushInt(value)
	return 0, nil
}
