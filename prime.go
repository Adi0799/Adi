package main

import (
	"fmt"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	sum := 0

	for i := 2; i <= 10; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Println("Sum of prime numbers between 1 and 10 is:", sum)
}
