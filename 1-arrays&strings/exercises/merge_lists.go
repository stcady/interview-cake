package main

func mergeLists(a []int, b []int) []int {

	c := make([]int, len(a)+len(b))
	currentA := 0
	currentB := 0
	currentC := 0
	for currentC < len(c) {
		exhaustedA := currentA >= len(a)
		exhaustedB := currentB >= len(b)
		if (exhaustedA != true) && (exhaustedB != true) {
			if a[currentA] <= b[currentB] {
				c[currentC] = a[currentA]
				currentA += 1
			} else if b[currentB] < a[currentA] {
				c[currentC] = b[currentB]
				currentB += 1
			}
		} else if (exhaustedA != false) && (exhaustedB != true) {
			c[currentC] = b[currentB]
			currentB += 1
		} else {
			c[currentC] = a[currentA]
			currentA += 1
		}
		currentC += 1
	}

	return c
}
