package main

type MaxStack struct {
	stack Stack
	max   Stack
}

func (s *MaxStack) Push(item int) {
	s.stack.Push(item)
	if s.max.Peek() != 0 && item > s.max.Peek() {
		s.max.Push(item)
	}
}

func (s *MaxStack) Pop() int {
	item := s.stack.Pop()
	if item == s.max.Peek() {
		s.max.Pop()
	}
	return item
}

func (s *MaxStack) Peek() int {
	return s.stack.Peek()
}

func (s *MaxStack) PeekMax() int {
	return s.max.Peek()
}

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if len(s.items) < 1 {
		return 0
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() int {
	if len(s.items) < 1 {
		return 0
	}
	return s.items[len(s.items)-1]
}
