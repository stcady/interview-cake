package main

func check(list []int, item int) bool {
	for _, element := range list {
		if element == item {
			return true
		}
	}
	return false
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
