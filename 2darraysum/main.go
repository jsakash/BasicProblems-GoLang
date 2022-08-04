package main

import "fmt"

func main() {

	var array1 [50][50]int
	var array2 [50][50]int
	var sum [50][50]int
	var size int

	fmt.Print("Enter the size of array : ")
	fmt.Scanln(&size)

	fmt.Println("Enter the values of first  array : ")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Scan(&array1[i][j])
		}
	}
	fmt.Println("Enter the values of second array : ")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Scan(&array2[i][j])
		}
	}
	fmt.Println("Sum =  ")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum[i][j] = array1[i][j] + array2[i][j]
			fmt.Print(sum[i][j], " ")
		}
		fmt.Println()
	}
}
