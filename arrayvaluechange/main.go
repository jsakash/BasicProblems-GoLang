package main

import "fmt"

func main() {

	var arr1 [50]int
	var arr2 [50]int
	var temp [50]int
	var size, i int

	fmt.Print("Enter the size of array : ")
	fmt.Scanln(&size)
	fmt.Print("Enter the values of first array : ")
	for i = 0; i < size; i++ {
		fmt.Scan(&arr1[i])
	}
	fmt.Print("Enter the values of second array : ")
	for i = 0; i < size; i++ {
		fmt.Scan(&arr2[i])
	}
	for i := 0; i < size; i++ {
		temp[i] = arr1[i]
		arr1[i] = arr2[i]
		arr2[i] = temp[i]

	}

	fmt.Print("\nFirst array after swapping = ")
	for i = 0; i < size; i++ {
		fmt.Print(" ", arr1[i])
	}
	fmt.Print("\nSecond array after swapping = ")
	for i = 0; i < size; i++ {
		fmt.Print(" ", arr2[i])
	}
	fmt.Print()
}
