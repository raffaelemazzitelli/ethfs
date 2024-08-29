package evm

import (
	"testing"
)

func TestOpPush(t *testing.T) {
	program := []byte{0x60, 0x12} // 0x60 is PUSH1, 0x12 is the value to push

	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expected := [32]byte{31: 0x12}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expected {
		t.Errorf("expected %x, got %x", expected, stackTop)
	}
}

func TestOpDup(t *testing.T) {
	program := []byte{0x60, 0x12, 0x60, 0x02, 0x80} // 0x60 is PUSH1, 0x12 is the value to push, 0x80 is DUP1

	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	// Check the stack
	expectedTop := [32]byte{31: 0x02}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}
	originalTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expectedTop || stackTop != originalTop {
		t.Errorf("expected %x,%x, got %x,%x", expectedTop, expectedTop, stackTop, originalTop)
	}
}

func TestOpSwap(t *testing.T) {

	program := []byte{0x60, 0x12, 0x60, 0x02, 0x90} // 0x60 is PUSH1, 0x12 is the value to push, 0x90 is SWAP1
	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	expectedTop := [32]byte{31: 0x12}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expectedTop {
		t.Errorf("expected %x, got %x", expectedTop, stackTop)
	}
}

func TestOpPop(t *testing.T) {
	program := []byte{0x60, 0x12, 0x60, 0x02, 0x60, 0x03, 0x50} // 0x60 is PUSH1, 0x12 is the value to push, 0x50 is POP
	evm := NewEVM(program, 1000, 0, nil)
	_ = evm.Run()

	expectedTop := [32]byte{31: 0x02}
	stackTop, err := evm.stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during stack pop: %v", err)
	}

	if stackTop != expectedTop {
		t.Errorf("expected %x, got %x", expectedTop, stackTop)
	}

}
