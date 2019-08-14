package stack

import "fmt"

type MyStack struct {
	Size  int
	Top   int
	Stack []int
}

func NewStack(size int) *MyStack {
	stack := &MyStack{
		Size:  size,
		Top:   -1,
		Stack: make([]int, size),
	}
	return stack
}

func (s *MyStack) IsEmpty() bool {
	return s.Top == -1
}

func (s *MyStack) IsFull() bool {
	return s.Top == s.Size-1
}

func (s *MyStack) Push(value int) {
	if s.IsFull() {
		fmt.Print("stack is full")
		return
	}
	s.Top++
	s.Stack[s.Top] = value
}

func (s *MyStack) Pop() int {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	res := s.Stack[s.Top]
	s.Top = s.Top - 1
	return res
}

func (s *MyStack) Peek() int {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.Stack[s.Top]
}

func (s *MyStack) Items() {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
	}

	for i := s.Top; i >= 0; i-- {
		fmt.Println(s.Stack[i],"index",i)
	}
}
