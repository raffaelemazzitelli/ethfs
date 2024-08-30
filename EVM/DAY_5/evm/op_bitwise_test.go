package evm

import (
	"testing"
)

func TestOpByte(t *testing.T) {
	program := []byte{0x60, 0xFF, 0x60, 0x1f, 0x1A}
	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0xFF}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}

func TestOpByte2(t *testing.T) {
	program := []byte{0x61, 0xFF, 0x00, 0x60, 0x1e, 0x1A}
	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0xFF}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}
