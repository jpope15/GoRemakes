package main

import (
	"fmt"
)

func main() {
	start, end := 0, 0

	fmt.Println("Enter a starting value: ")
	fmt.Scanln(&start)

	fmt.Println("Enter a ending value: ")
	fmt.Scanln(&end)

	for i := start; i <= end; i++ {
		if isPrime(i) {
			fmt.Println(fmt.Sprintf("%d is a prime number", i))
		}
	}
}

func isPrime(x int) bool {
	if x <= 1 {
		return false
	}
	for i := 2; i < 1+x/2; i++ {
		for j := i; i*j <= x; j++ {
			if i*j == x {
				return false
			}
		}
	}
	return true
}
