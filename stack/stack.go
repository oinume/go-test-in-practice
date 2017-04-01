package stack

import "fmt"

type intElement struct {
	value int
	next  *intElement
}

type IntStack struct {
	top  *intElement
	size int
}

func NewIntStack() *IntStack {
	return &IntStack{}
}

func (s *IntStack) Push(value int) {
	t := s.top
	s.top = &intElement{
		value: value,
		next:  t,
	}
}

func (s *IntStack) Pop() (int, error) {
	if s.Size() == 0 {
		return 0, fmt.Errorf("Stack is empty.")
	}
	v := s.top.value
	s.top = s.top.next
	return v, nil
}

func (s *IntStack) Size() int {
	return s.size
}
