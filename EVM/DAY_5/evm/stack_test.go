package evm

import (
	"strings"
	"testing"
)

func TestStack_PushPop(t *testing.T) {
	stack := NewStack()
	value := [32]byte{0x01}

	err := stack.Push(value)
	if err != nil {
		t.Fatalf("unexpected error during push %v", err)
	}

	poppedValue, err := stack.Pop()
	if err != nil {
		t.Fatalf("unexpected error during pop: %v", err)
	}

	if poppedValue != value {
		t.Fatalf("expected %x, got %x", value, poppedValue)
	}
}

func TestStack_Underflow(t *testing.T) {
	stack := NewStack()
	_, err := stack.Pop()
	if err == nil {
		t.Errorf("expected stack underflow error, got nil")
	}
}

func TestStack_Overflow(t *testing.T) {
	stack := NewStack()

	for i := 0; i < 1024; i++ {
		value := [32]byte{byte(i)}
		err := stack.Push(value)
		if err != nil {
			t.Fatalf("unexpected error during push %d: %v", i, err)
		}
	}

	err := stack.Push([32]byte{0xff})
	if err == nil || !strings.HasPrefix(err.Error(), "Stack overflow") {
		t.Errorf("expected stack overflow error, got %v", err)
	}
}

func TestStack_Order(t *testing.T) {
	stack := NewStack()
	values := [][32]byte{
		{0x01},
		{0x02},
		{0x03},
	}

	for _, value := range values {
		err := stack.Push(value)
		if err != nil {
			t.Fatalf("unexpected error during push: %v", err)
		}
	}

	for i := len(values) - 1; i >= 0; i-- {
		poppedValue, err := stack.Pop()
		if err != nil {
			t.Fatalf("unexpected error during pop: %v", err)
		}
		if poppedValue != values[i] {
			t.Errorf("expected %x,got %x", values[i], poppedValue)
		}
	}
}

func TestStack_Peek(t *testing.T) {
	stack := NewStack()
	values := [][32]byte{
		{0x01},
		{0x02},
		{0x03},
	}

	for _, value := range values {
		err := stack.Push(value)
		if err != nil {
			t.Fatalf("unexpected error during push: %v", err)
		}
	}

	peekValue, err := stack.Peek(1)
	if err != nil {
		t.Fatalf("unexpected error during peek: %v", err)
	}
	expected := [32]byte{0x02}
	if peekValue != expected {
		t.Errorf("expected %x, got %x", expected, peekValue)
	}
}

func TestStack_Swap(t *testing.T) {
	stack := NewStack()
	values := [][32]byte{
		{0x01},
		{0x02},
		{0x03},
	}

	for _, value := range values {
		err := stack.Push(value)
		if err != nil {
			t.Fatalf("unexpected error during push: %v", err)
		}
	}

	// Perform the swap: swap the top element with the 3rd element (1-indexed)
	err := stack.Swap(0, 2)
	if err != nil {
		t.Fatalf("unexpected error during swap: %v", err)
	}

	// Verify the stack order after swap
	expectedOrder := [][32]byte{
		{0x01},
		{0x02},
		{0x03},
	}

	for i := 0; i < len(expectedOrder); i++ {
		poppedValue, err := stack.Pop()
		if err != nil {
			t.Fatalf("unexpected error during pop: %v", err)
		}
		if poppedValue != expectedOrder[i] {
			t.Errorf("expected %x, got %x", expectedOrder[i], poppedValue)
		}
	}

	for _, value := range values {
		err := stack.Push(value)
		if err != nil {
			t.Fatalf("unexpected error during push: %v", err)
		}
	}

	// Perform the swap: swap the top element with the 3rd element (1-indexed)
	err = stack.Swap(0, 1)
	if err != nil {
		t.Fatalf("unexpected error during swap: %v", err)
	}

	// Verify the stack order after swap
	expectedOrder = [][32]byte{
		{0x02},
		{0x03},
		{0x01},
	}

	for i := 0; i < len(expectedOrder); i++ {
		poppedValue, err := stack.Pop()
		if err != nil {
			t.Fatalf("unexpected error during pop: %v", err)
		}
		if poppedValue != expectedOrder[i] {
			t.Errorf("expected %x, got %x", expectedOrder[i], poppedValue)
		}
	}
}
