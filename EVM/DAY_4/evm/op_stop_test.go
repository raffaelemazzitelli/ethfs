package evm

import (
	"testing"
)

func TestEVM_RunStopsOnOpcode0x00(t *testing.T) {
	// Program with a 0x00 opcode at the second instruction
	program := []byte{0x00, 0x02}

	evm := NewEVM(program, 1000, 0, nil)
	err := evm.Run()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if evm.pc != 0 {
		t.Errorf("Expected pc to be 0 after execution, got %d", evm.pc)
	}

	if !evm.stop_flag {
		t.Errorf("Expected stop_flag to be true, but it was false")
	}
}

// func TestEVM_ProgramCounterIncrement(t *testing.T) {
// 	// Program with three non-stop opcodes
// 	program := []byte{0x01, 0x02, 0x03}

// 	evm := NewEVM(program, 1000, 0, nil)
// 	err := evm.Run()

// 	if err != nil {
// 		t.Fatalf("Unexpected error: %v", err)
// 	}

// 	if evm.pc != 3 {
// 		t.Errorf("Expected pc to be 3 after execution, got %d", evm.pc)
// 	}

// 	if evm.stop_flag {
// 		t.Errorf("Expected stop_flag to be false, but it was true")
// 	}
// }
