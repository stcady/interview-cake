package main

import "fmt"

func productOfInts(list []int) ([]int, error) {
	if len(list) < 2 {
		return []int{}, fmt.Errorf("invalid list size")
	}
	productList := make([]int, len(list))
	productSoFar := 1
	i := 0
	for i < len(productList) {
		productList[i] = productSoFar
		productSoFar *= list[i]
		i += 1
	}
	i = len(list) - 1
	for i >= 0 {
		productList[i] = productSoFar
		productSoFar *= list[i]
		i -= 1
	}
	return productList, nil
}
