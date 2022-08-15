package main

func findRotationPoint(list []string) int {
	firstItem := list[0]
	left := 0
	right := len(list) - 1
	for left+1 < right {
		mid := left + ((right - left) / 2)
		if list[mid] >= firstItem {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}
