package main

type Cake struct {
	weight int
	value  int
}

func maxBagValue(cakeTuples []Cake, capacity int) int {
	// We make a list to hold the maximum possible value at every
	// duffel bag weight capacity from 0 to weight_capacity
	// starting each index with value 0
	maxValuesAtCapacities := []int{}
	rangeCapacities := []int{}
	i := 0
	for i <= capacity {
		maxValuesAtCapacities[i] = 0
		rangeCapacities[i] = i
	}

	for currentCapacity := range rangeCapacities {
		// Set a variable to hold the max monetary value so far for currentCapacity
		currentMaxValue := 0
		maxValueUsingCake := 0
		for _, cake := range cakeTuples {
			cakeWeight := cake.weight
			cakeValue := cake.value
			if cakeWeight == 0 && cakeValue != 0 {
				return 1000000000000000000
			}

			// If the current cake weighs as much or less than the
			// current weight capacity it's possible taking the cake
			// would get a better value
			if cakeWeight <= currentCapacity {
				maxValueUsingCake = cakeValue + maxValuesAtCapacities[currentCapacity-cakeWeight]
			}

			// Now we see if it's worth taking the cake. how does the
			// value with the cake compare to the current_max_value?
			currentMaxValue = max(maxValueUsingCake, currentMaxValue)
		}
		// Add each capacity's max value to our list so we can use them
		// when calculating all the remaining capacities
		maxValuesAtCapacities[currentCapacity] = currentMaxValue
	}
	return maxValuesAtCapacities[capacity]
}

// Max returns the larger of x or y.
func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
