package evm

func (evm *EVM) opStop() (int, error) {
	evm.stop_flag = true
	return 0, nil
}
