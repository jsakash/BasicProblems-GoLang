package main

import "fmt"

func main() {

	var arr [50]int
	var sum [50]int
	var size int

	fmt.Print("Enter the size of array = ")
	fmt.Scan(&size)

	fmt.Print("Enter the values of array = ")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 0; i < size; i++ {
		sum[i] = arr[i] * arr[i+1]
	}
	fmt.Print("Array after multiplication = ")
	for i := 0; i < size-1; i++ {
		fmt.Print(sum[i], " ")
	}
	fmt.Println()
}
