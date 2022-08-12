package main

import "math/rand"

func shuffle(list []int) []int {
	if len(list) < 2 {
		return list
	}

	i := 0
	end := len(list) - 1
	for i < len(list) {
		j := rand.Intn(i-end) + i
		if i != j {
			list[i], list[j] = list[j], list[i]
		}
	}
	return list
}
