package main

import (
	"fmt"
	"math"
)

func printPrimes(max int) {
	for n := 2; n <= max; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}

		if n%2 == 0 {
			continue
		}

		multipliable := false

		for i := 3; float64(i) <= math.Sqrt(float64(n))+1; i++ {
			if n%i == 0 {
				multipliable = true
				break
			}
		}

		if multipliable == false {
			fmt.Println(n)
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
