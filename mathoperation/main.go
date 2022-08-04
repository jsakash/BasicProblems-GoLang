package main

import "fmt"

var num1, num2, choice int

func main() {

	fmt.Print("Enter two numbers : ")
	fmt.Scan(&num1, &num2)

	fmt.Println("1 for Addition", "\n2 for Subtraction", "\n3 for Multiplication", "\n4 for Division")
	fmt.Scanln(&choice)

	if choice == 1 {
		addition(num1, num2)
	} else if choice == 2 {
		subtraction(num1, num2)
	} else if choice == 3 {
		multiplication(num1, num2)
	} else if choice == 4 {
		division(num1, num2)
	}
}
func addition(num1 int, num2 int) {
	result := num1 + num2
	fmt.Println("Result = ", result)
}
func subtraction(num1 int, num2 int) {
	result := num1 - num2
	fmt.Println("Result = ", result)
}
func multiplication(num1 int, num2 int) {
	result := num1 * num2
	fmt.Println("Result = ", result)
}
func division(num1 int, num2 int) {
	result := num1 / num2
	fmt.Println("Result = ", result)
}
