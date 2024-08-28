package evm

func (evm *EVM) opPush(n int) func(*EVM) (int, error) {
	return func(evm *EVM) (int, error) {
		value := make([]byte, n)
		for i := 0; i < n; i++ {
			position := evm.pc + 1 + i
			value[i] = evm.program[position]
		}
		var result [32]byte
		copy(result[32-len(value):], value)
		evm.stack.Push(result)
		evm.pc += n
		return 0, nil
	}
}

func (evm *EVM) opDup(position int) func(*EVM) (int, error) {
	return func(evm *EVM) (int, error) {
		item, err := evm.stack.Peek(position - 1)
		if err != nil {
			return 0, err
		}
		evm.stack.Push(item)
		return 0, nil
	}
}

func (evm *EVM) opSwap(position int) func(*EVM) (int, error) {
	return func(evm *EVM) (int, error) {
		evm.stack.Swap(0, position)
		return 0, nil
	}
}

func (evm *EVM) opPop() (int, error) {
	_, err := evm.stack.Pop()
	return 0, err
}
