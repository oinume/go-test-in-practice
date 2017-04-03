package stack

import "fmt"

type intElement struct {
	value int
	next  *intElement
}

type IntStack struct {
	head *intElement
	size int
}

func NewIntStack() *IntStack {
	return &IntStack{}
}

func (s *IntStack) Push(value int) {
	t := s.head
	s.head = &intElement{
		value: value,
		next:  t,
	}
	s.size++
}

func (s *IntStack) Pop() (int, error) {
	if s.Size() == 0 {
		return 0, fmt.Errorf("Stack is empty.")
	}
	v := s.head.value
	s.head = s.head.next
	s.size--
	return v, nil
}

func (s *IntStack) Size() int {
	return s.size
}
