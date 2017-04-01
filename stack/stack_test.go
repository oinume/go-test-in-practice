package stack

import "testing"

func TestNewIntStack(t *testing.T) {
	istack := NewIntStack()
	if istack == nil {
		t.Error("NewIntStack() returns nil")
	}
}
