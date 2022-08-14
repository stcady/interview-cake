package main

func canTwoMoviesFillFlight(movieLengths []int, flightLength int) bool {
	set := make(map[int]bool)
	for firstMovieLength := range movieLengths {
		secondMovieLength := flightLength - firstMovieLength
		if _, ok := set[secondMovieLength]; ok {
			return true
		}
		set[firstMovieLength] = true
	}
	return false
}
