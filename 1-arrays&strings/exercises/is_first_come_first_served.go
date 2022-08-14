package main

func isFirstComeFirstServed(takeOut []int, dineIn []int, served []int) bool {

	indexTakeOut := 0
	indexDineIn := 0
	endTakeOut := len(takeOut)
	endDineIn := len(dineIn)

	for _, order := range served {
		if (indexTakeOut < endTakeOut) && (order == takeOut[indexTakeOut]) {
			indexTakeOut += 1
		} else if (indexDineIn < endDineIn) && (order == dineIn[indexDineIn]) {
			indexDineIn += 1
		} else {
			return false
		}
	}
	if (indexTakeOut != endTakeOut) && (indexDineIn != endDineIn) {
		return false
	}
	return true
}
