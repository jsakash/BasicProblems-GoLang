package main

import "fmt"

func main() {

	var totalmark float32

	fmt.Print("Enter your mark percentage : ")
	fmt.Scanln(&totalmark)

	if totalmark >= 90 {
		fmt.Println("A Grade")
	} else if totalmark >= 80 {
		fmt.Println("B Grade")
	} else if totalmark >= 70 {
		fmt.Println("C Grade")
	} else if totalmark >= 60 {
		fmt.Println("D Grade")
	} else if totalmark >= 50 {
		fmt.Println("E Grade")
	} else if totalmark < 50 {
		fmt.Println("FAILED")
	}
}
