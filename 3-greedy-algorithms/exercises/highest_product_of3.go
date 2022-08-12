package main

import "fmt"

func highestProductOf3(list []int) (int, error) {

	if len(list) < 3 {
		return 0, fmt.Errorf("list is less than 3")
	}

	highest := max(list[0], list[1])
	lowest := min(list[0], list[1])
	highestProductOf2 := list[0] * list[1]
	lowestProductOf2 := list[0] * list[1]
	highestProductOf3 := list[0] * list[1] * list[2]

	temp := []int{}
	i := 2
	for i < len(list) {
		temp = append(temp, i)
		i += 1
	}

	for j := range temp {
		current := list[j]
		highestProductOf3 = max(highestProductOf3, max(current*highestProductOf2, current*lowestProductOf2))
		highestProductOf2 = max(highestProductOf2, max(current*highest, current*lowest))
		lowestProductOf2 = min(lowestProductOf2, min(current*highest, current*lowest))
		highest = max(highest, current)
		lowest = min(lowest, current)
	}
	return highestProductOf3, nil
}

// Max returns the larger of x or y.
func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
