package evm

func (evm *EVM) opPush(n int) func(*EVM) error {
	return func(evm *EVM) error {
		value := make([]byte, n)
		for i := 0; i < n; i++ {
			position := evm.pc + 1 + i
			value[i] = evm.program[position]
		}
		var result [32]byte
		copy(result[32-len(value):], value)
		evm.stack.Push(result)
		evm.pc += n
		return nil
	}
}

func (evm *EVM) opDup(position int) func(*EVM) error {
	return func(evm *EVM) error {
		item, err := evm.stack.Peek(position - 1)
		if err != nil {
			return err
		}
		evm.stack.Push(item)
		return nil
	}
}

func (evm *EVM) opSwap(position int) func(*EVM) error {
	return func(evm *EVM) error {
		evm.stack.Swap(0, position)
		return nil
	}
}

func (evm *EVM) opPop() error {
	_, err := evm.stack.Pop()
	return err
}
