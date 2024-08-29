package evm

import (
	"testing"
)

func TestOpLt(t *testing.T) {
	program := []byte{0x60, 0x10, 0x60, 0x09, 0x10} // 0x09<0x10 ==> 1

	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0x01}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}
func TestOpLt2(t *testing.T) {
	program := []byte{0x60, 0x10, 0x60, 0x10, 0x10} // 0x10<0x10 ==> 0

	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0x00}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}

func TestOpGt(t *testing.T) {
	program := []byte{0x60, 0x10, 0x60, 0x09, 0x11} // 0x10>0x09 ==> 1

	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0x0}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}

func TestOpGt2(t *testing.T) {
	program := []byte{0x60, 0x10, 0x60, 0x10, 0x11} // 0x10>0x10 ==> 0

	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0x0}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}
func TestOpSlt(t *testing.T) {
	program := []byte{0x60, 0x00, 0x7f, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x12}
	// 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF < 0x09 --> 0x01

	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0x01}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}

func TestOpSlt2(t *testing.T) {
	program := []byte{0x60, 0x10, 0x60, 0x10, 0x12}
	// 0x10 < 0x10 --> 0x00

	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0x00}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}

func TestOpSgt(t *testing.T) {
	program := []byte{0x7f, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x60, 0x00, 0x13}
	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0x01}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}
func TestOpSgt2(t *testing.T) {
	program := []byte{0x60, 0x10, 0x60, 0x10, 0x13}
	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0x00}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}

func TestOpEq(t *testing.T) {
	program := []byte{0x60, 0x10, 0x60, 0x10, 0x14}

	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0x01}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}

func TestOpEq2(t *testing.T) {
	program := []byte{0x60, 0x10, 0x60, 0x11, 0x14}

	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0x00}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}
func TestOpIsZero(t *testing.T) {
	program := []byte{0x60, 0x00, 0x15}

	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0x01}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}

func TestOpIsZero2(t *testing.T) {
	program := []byte{0x60, 0x01, 0x15}

	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0x00}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}
