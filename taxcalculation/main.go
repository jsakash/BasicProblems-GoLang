package main

import "fmt"

func main() {

	var income float32

	fmt.Print("Enter the annual income = ")
	fmt.Scan(&income)

	if income < 250000 {
		fmt.Println("No Tax")
	} else if income >= 250000 {
		fmt.Println("Income Tax amount = ", income*5/100)
	} else if income >= 500000 {
		fmt.Println("Income Tax amount = ", income*20/100)
	} else if income >= 1000000 {
		fmt.Println("Income Tax amount = ", income*30/100)
	}

}
