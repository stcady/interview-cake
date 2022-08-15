package main

import "sort"

type Times struct {
	Start int
	End   int
}

func mergeRanges(list []Times) {

	sort.SliceStable(list, func(i, j int) bool {
		return list[i].End < list[j].End
	})
	mergedMeetings := []Times{list[0]}
	for _, current := range list {
		last := mergedMeetings[len(list)-1]
		if current.Start <= last.Start {
			mergedMeetings[len(list)-1] = Times{
				Start: last.Start,
				End:   max(last.End, current.End),
			}
		} else {
			mergedMeetings = append(mergedMeetings, Times{
				Start: current.Start,
				End:   current.End,
			})
		}
	}
}

// Max returns the larger of x or y.
func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
