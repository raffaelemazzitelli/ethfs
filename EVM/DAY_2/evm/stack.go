package evm

import (
	"errors"
	"fmt"
)

type Stack struct {
	data     [1024][32]byte
	elements int
}

func NewStack() *Stack {
	return &Stack{
		elements: 0,
	}
}

func (s *Stack) Push(value [32]byte) error {
	if len(s.data) < s.elements+1 {
		return fmt.Errorf("Stack overflow: current elements %d ", s.elements)
	}
	s.data[s.elements] = value
	s.elements++
	return nil
}

func (s *Stack) Pop() ([32]byte, error) {
	if s.elements == 0 {
		return [32]byte{}, errors.New("Stack underflow")
	}
	s.elements--
	topValue := s.data[s.elements]

	return topValue, nil
}

func (s *Stack) Peek(position int) ([32]byte, error) {
	if position < 0 || position >= s.elements {
		return [32]byte{}, errors.New("Stack peek out of bounds")
	}
	return s.data[s.elements-1-position], nil
}

// Swap exchanges the elements at the top of the stack and the position specified.
func (s *Stack) Swap(pos1, pos2 int) error {
	if pos1 < 0 || pos1 >= s.elements || pos2 < 0 || pos2 >= s.elements {
		return errors.New("Stack swap out of bounds")
	}

	tmp1 := s.data[s.elements-1-pos1]
	tmp2 := s.data[s.elements-1-pos2]

	s.data[s.elements-1-pos2] = tmp1
	s.data[s.elements-1-pos1] = tmp2

	return nil
}
