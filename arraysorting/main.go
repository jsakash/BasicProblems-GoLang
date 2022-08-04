package main

import "fmt"

func main() {

	var value [50]int
	var size, temp int

	fmt.Print("Enter the size of array : ")
	fmt.Scanln(&size)
	fmt.Print("Enter the values of array : ")
	for i := 0; i < size; i++ {
		fmt.Scan(&value[i])
	}

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if value[i] < value[j] {
				temp = value[i]
				value[i] = value[j]
				value[j] = temp
			}

		}
	}
	fmt.Print("Sorted array = ")
	for i := 0; i < size; i++ {
		fmt.Print(" ", value[i])
	}
	fmt.Println()

}
