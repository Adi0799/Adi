package main

import (
	"calculator/calc" // Import the custom package
	"fmt"
)

func main() {
	num1 := 10
	num2 := 5

	result1, result2, result3 := calc.Calc(num1, num2)

	fmt.Println("Addition:", result1)
	fmt.Println("Subtraction:", result2)
	fmt.Println("Division:", result3)
}
