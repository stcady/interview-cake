package main

func reverse(list []int) []int {
	left := 0
	right := len(list) - 1
	for left < right {
		list[left], list[right] = list[right], list[left]
		left += 1
		right -= 1
	}
	return list
}
