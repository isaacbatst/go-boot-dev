package main

import "fmt"

func main() {
	max := getMaxMessagesToSend(2, 10)

	fmt.Printf("No máximo %v mensagens", max)
}

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	for actualCostInPennies < float64(maxCostInPennies) {
		actualCostInPennies *= costMultiplier
		maxMessagesToSend++
	}
	if actualCostInPennies > float64(maxCostInPennies) {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}
