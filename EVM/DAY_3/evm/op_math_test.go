package evm

import (
	"testing"
)

func TestOpAdd(t *testing.T) {
	program := []byte{0x7f, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x60, 0x01, 0x01} // PUSH32 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF, PUSH1 0x01, ADD ==> expeted top stack 0x00 and overflow

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

func TestOpMul(t *testing.T) {
	program := []byte{0x7f, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x60, 0x02, 0x02} // PUSH32 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF, PUSH1 0x02, MUL ==> expeted top stack 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE and overflow

	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}

func TestOpSub(t *testing.T) {
	program := []byte{0x60, 0x01, 0x60, 0x00, 0x03} // PUSH1 0x0, PUSH1 0x01, SUB ==> expeted top stack 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF and underflow

	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}

func TestOpDIV(t *testing.T) {
	program := []byte{0x60, 0x02, 0x60, 0x01, 0x04} // PUSH1 0x02, PUSH1 0x01, DIV ==> expected 0

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

func TestOpSDIV(t *testing.T) {
	program := []byte{0x60, 0x02, 0x60, 0x01, 0x05} // PUSH1 0x02, PUSH1 0x01, SDIV ==> expected 0

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

func TestOpMod(t *testing.T) {
	program := []byte{0x60, 0x05, 0x60, 0x11, 0x06} // PUSH1 0x05, PUSH1 0x11 (17), MOD ==> expected 2

	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0x02}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}

func TestOpSMod(t *testing.T) {
	program := []byte{0x60, 0x03, 0x60, 0x0a, 0x07} // PUSH1 0x03, PUSH1 0x0a (10), MOD ==> expected 1

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

func TestOpAddMod(t *testing.T) {
	program := []byte{0x60, 0x0a, 0x60, 0x0a, 0x60, 0x08, 0x08} // (10+10)%8=4

	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0x04}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}

func TestOpMulMod(t *testing.T) {
	program := []byte{0x60, 0x0a, 0x60, 0x0a, 0x60, 0x08, 0x09} // (10+10)%8=4

	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0x04}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}

func TestOpExp(t *testing.T) {
	program := []byte{0x60, 0x02, 0x60, 0x0a, 0x0a} //

	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0x64}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}

func TestOpExtendSign(t *testing.T) {
	program := []byte{0x60, 0x7F, 0x60, 0x0, 0x0b} //

	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0x7F}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}
