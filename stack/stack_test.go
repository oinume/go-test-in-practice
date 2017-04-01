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
	istack.Push(3)

	v, err := istack.Pop()
	if err != nil {
		t.Errorf("err must not be returned: %v", err)
	}
	if want, got := 3, v; want != got {
		t.Errorf("Pop: want=%v, got=%v", want, got)
	}

	t.Log("Check stack size")
	if want, got := 2, istack.Size(); want != got {
		t.Errorf("Size: want=%v, got=%v", want, got)
	}
}
