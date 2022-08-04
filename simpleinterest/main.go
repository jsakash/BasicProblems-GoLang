package main

import "fmt"

func main() {

	var principal1 int
	var interest, numofyear, simpleinterest, principal float64

	fmt.Print("Enter the Prinipal Amount :")
	fmt.Scanln(&principal1)
	principal = float64(principal1)
	fmt.Print("Enter interst rate :")
	fmt.Scanln(&interest)
	fmt.Print("Enter the number of year :")
	fmt.Scanln(&numofyear)

	simpleinterest = (principal * interest * numofyear) / 100

	fmt.Println("Simple Interest = :", simpleinterest)

}
