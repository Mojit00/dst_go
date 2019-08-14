package stack

import "testing"


func TestMyStack_IsEmpty(t *testing.T) {
	stack := NewStack(5)
	t.Log(stack.IsEmpty())
	t.Log(stack.IsFull())
}

func TestMyStack_Push(t *testing.T) {
	stack := NewStack(5)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	t.Log(stack.IsFull())
	t.Log(stack.Pop())
	t.Log(stack.Pop())

}
