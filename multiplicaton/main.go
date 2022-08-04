package main

import "fmt"

func main() {

	var number, i int

	fmt.Println("Enter a number ")
	fmt.Scanln(&number)
	for i = 1; i <= 10; i++ {
		fmt.Println(number, "x", i, "=", number*i)
	}

}
