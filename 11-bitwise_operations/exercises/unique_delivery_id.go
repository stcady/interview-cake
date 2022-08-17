package main

func findUniqueDeliveryId(deliveryIds []int) int {
	uniqueDeliveryId := 0
	for deliveryId := range deliveryIds {
		uniqueDeliveryId ^= deliveryId
	}
	return uniqueDeliveryId
}
