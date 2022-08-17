package main

import "math"

type TempTracker struct {
	occurrences    []int
	maxOccurrences int
	mode           int

	n        int
	totalSum int
	mean     float32

	minimum int
	maximum int
}

func New() *TempTracker {
	return &TempTracker{
		occurrences:    make([]int, 111),
		maxOccurrences: 0,
		mode:           0,
		n:              0,
		totalSum:       0,
		mean:           0.0,
		minimum:        int(math.Inf(1)),
		maximum:        int(math.Inf(-1)),
	}
}

func (t *TempTracker) Insert(temp int) {
	t.occurrences[temp] += 1
	if t.occurrences[temp] > t.maxOccurrences {
		t.mode = temp
		t.maxOccurrences = t.occurrences[temp]
	}
	t.n += 1
	t.totalSum += temp
	t.mean = float32(t.totalSum) / float32(t.n)
	if temp > t.maximum {
		t.maximum = temp
	}
	if temp < t.minimum {
		t.minimum = temp
	}
}

func (t *TempTracker) GetMax() int {
	return t.maximum
}

func (t *TempTracker) GetMin() int {
	return t.minimum
}
func (t *TempTracker) GetMean() float32 {
	return t.mean
}
func (t *TempTracker) GetMode() int {
	return t.mode
}
