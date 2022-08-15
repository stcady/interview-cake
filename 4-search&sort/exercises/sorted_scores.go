package main

func sortedScores(scores []int, highestScore int) []int {
	scoreCounts := make([]int, highestScore+1)
	for _, score := range scores {
		scoreCounts[score] += 1
	}
	i := highestScore
	sortedScores := []int{}
	for i >= 0 {
		j := 0
		for j < scoreCounts[i] {
			sortedScores = append(sortedScores, i)
			j += 1
		}
		i -= 1
	}
	return sortedScores
}
