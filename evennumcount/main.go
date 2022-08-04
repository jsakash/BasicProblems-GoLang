package main

import "fmt"

func main() {

	var arr [50]int
	var i, size, even_count int

	fmt.Print("Enter the size of array : ")
	fmt.Scanln(&size)
	fmt.Print("Enter the values of first array : ")
	for i = 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 0; i < size; i++ {
		if arr[i]%2 == 0 {
			even_count++
		}
	}
	fmt.Println("Number of even numbers in the given array is ", even_count)
}
