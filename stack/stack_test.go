package stack

import (
	"reflect"
	"testing"
)

func TestNewIntStack(t *testing.T) {
	istack := NewIntStack()
	if istack == nil {
		t.Error("NewIntStack() returns nil")
	}
	t.Log("TestNewIntStack finished")
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

func TestIntStack_Compare(t *testing.T) {
	istack1 := NewIntStack()
	istack2 := NewIntStack()
	for i := range []int{1, 2, 3} {
		istack1.Push(i)
		istack2.Push(i)
	}
	if !reflect.DeepEqual(istack1, istack2) {
		t.Errorf("Stacks must be the same: istack1.Size=%d, istack2.Size=%d", istack1.Size(), istack2.Size())
	}
}
