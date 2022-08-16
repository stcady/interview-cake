package main

func fib(n int) int {
	if n < 0 {
		return 0
	} else if n < 2 {
		return n
	}

	prevPrev := 0
	prev := 1
	i := n - 1
	current := 0
	for i > 0 {
		current = prev + prevPrev
		prevPrev = prev
		prev = current
	}
	return current
}
