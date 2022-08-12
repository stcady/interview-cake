package main

import "fmt"

func getMaxProfit(stockPrices []int) (int, error) {

	// check length
	if len(stockPrices) < 2 {
		return 0, fmt.Errorf("invalid array length")
	}

	// set min price and max profit
	minPrice := stockPrices[0]
	maxProfit := stockPrices[1] - stockPrices[0]

	// iterate greedily
	for currentValue := range stockPrices {
		profit := currentValue - minPrice
		maxProfit = max(profit, maxProfit)
		minPrice = min(currentValue, minPrice)
	}
	return maxProfit, nil
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

func main() {
	stockPrices := []int{10, 7, 5, 8, 11, 9}
	fmt.Println(getMaxProfit(stockPrices))
}
