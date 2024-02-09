package main

import "fmt"

func maxMessages(thresh float64) int {

}

func test(tresh float64) {
	fmt.Printf("Treshold: %.2f\n", tresh)
	max := maxMessages(tresh)
	fmt.Printf("Max messages: %d\n", max)
	fmt.Println("=======================")
}

func main() {
	test(10.0)
	test(20.0)
	test(30.0)
	test(5.1)
	test(40.0)
	test(50.0)
}
