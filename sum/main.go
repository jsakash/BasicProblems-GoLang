package main

import "fmt"

func main() {

	var num int
	var num2, sum, num1 float64
	fmt.Println("Enter first number")
	fmt.Scanln(&num)
	num1 = float64(num)
	fmt.Println("Enter second number")
	fmt.Scanln(&num2)
	sum = num1 + num2
	fmt.Println("Sum is : ", sum)
}
