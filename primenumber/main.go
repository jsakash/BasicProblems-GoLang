package main

import "fmt"

func main() {

	var num int
	flag := 0

	fmt.Println("Enter a number")
	fmt.Scanln(&num)

	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			flag = 1
		}
	}
	if flag == 0 {
		fmt.Println("Number is Prime")
	}
	if flag == 1 {
		fmt.Println("Number is not Prime")
	}
 
}
