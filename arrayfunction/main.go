package main

import "fmt"

var array [50]int
var size int

func main() {

	getArray()
	displayArray()
}

func getArray() {

	fmt.Print("Enter the size of Array ")
	fmt.Scanln(&size)
	fmt.Println("Enter the values of array")
	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
	}
}

func displayArray() {

	fmt.Println("Array values ")
	for i := 0; i < size; i++ {
		fmt.Print(array[i], " ")
	}
	fmt.Println("")
}
