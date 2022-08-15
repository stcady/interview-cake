package main

func findRepeat(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		mid := left + ((right - left) / 2)
		lowerLeft, lowerRight := left, mid
		upperLeft, upperRight := mid+1, right
		numberOfItems := 0
		numDistinctItems := (lowerRight - lowerLeft) + 1
		for _, number := range numbers {
			if (number >= lowerLeft) && (number <= lowerRight) {
				numberOfItems += 1
			}
		}
		if numberOfItems >= numDistinctItems {
			left, right = lowerLeft, lowerRight
		} else {
			left, right = upperLeft, upperRight
		}
	}
	return left
}
