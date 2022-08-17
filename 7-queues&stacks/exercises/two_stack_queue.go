package main

type TwoStackQueue struct {
	inStack  Stack
	outStack Stack
}

func (tsq *TwoStackQueue) Enqueue(item int) {
	tsq.inStack.Push(item)
}

func (tsq *TwoStackQueue) Dequeue() int {
	for len(tsq.inStack.items) > 0 {
		tsq.outStack.Push(tsq.inStack.Pop())
	}
	item := tsq.outStack.Pop()
	for len(tsq.outStack.items) > 0 {
		tsq.inStack.Push(tsq.outStack.Pop())
	}
	return item
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
