package main

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	for {
		actualCostInPennies *= costMultiplier
		maxMessagesToSend++
	}
	if actualCostInPennies > float64(maxCostInPennies) {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}
