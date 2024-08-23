package evm

func (evm *EVM) opStop() error {
	evm.stop_flag = true
	return nil
}
