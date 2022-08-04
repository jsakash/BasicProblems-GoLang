package main

import "fmt"

func main() {

	var mark float32
	fmt.Print("Enter Your Mark : ")
	fmt.Scanln(&mark)

	if mark >= 50 {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

}
