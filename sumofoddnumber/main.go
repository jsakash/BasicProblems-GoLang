package main

import "fmt"

func main() {

	var limit, i, sum int
	fmt.Print("Enter the limit : ")
	fmt.Scanln(&limit)

	for i = 0; i < limit; i++ {
		if i%2 != 0 {
			sum = sum + i
		}
	}
	fmt.Println("Sum off odd numbers = ", sum)

}
