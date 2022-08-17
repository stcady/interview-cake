package main

func AreBracketsValid(code string) bool {
	table := make(map[rune]rune)
	table['{'] = '}'
	table['['] = ']'
	table['('] = ')'
	openers := []rune{'{', '[', '('}
	closers := []rune{'}', ']', ')'}
	stack := &Stack{
		items: []rune{},
	}
	for _, char := range code {
		if check(openers, char) {
			stack.Push(char)
		} else if check(closers, char) {
			if len(stack.items) < 1 {
				return false
			} else {
				lastOpenerChar := stack.Pop()
				if char != table[lastOpenerChar] {
					return false
				}
			}
		}

	}
	if len(stack.items) > 0 {
		return false
	}
	return true
}

func check(list []rune, value rune) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

type Stack struct {
	items []rune
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() rune {
	if len(s.items) < 1 {
		return 0
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() rune {
	if len(s.items) < 1 {
		return 0
	}
	return s.items[len(s.items)-1]
}
