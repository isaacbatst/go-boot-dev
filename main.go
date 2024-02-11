package main

import "fmt"

func fizzbuzz() {
	for i := range 100 {
		n := i + 1
		fizzbuzz := getString(n)
		fmt.Printf("%v: %v\n", n, fizzbuzz)
	}
}

func getString(n int) string {
	if n%5 == 0 && n%3 == 0 {
		return "fizzbuzz"
	}

	if n%5 == 0 {
		return "buzz"
	}

	if n%3 == 0 {
		return "fizz"
	}
	return ""
}

func main() {
	fizzbuzz()
}
