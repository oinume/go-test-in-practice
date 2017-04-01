package stack

import "testing"

func TestNewIntStack(t *testing.T) {
	istack := NewIntStack()
	if istack == nil {
		t.Error("NewIntStack() returns nil")
	}
}

func TestIntStack_Pop(t *testing.T) {
	istack := NewIntStack()
	istack.Push(1)
	istack.Push(2)

	v, err := istack.Pop()
	if err != nil {
		t.Errorf("err must not be returned: %v", err)
	}
	if want, got := 2, v; want != got {
		t.Errorf("want=%v, got=%v", want, got)
	}
}
