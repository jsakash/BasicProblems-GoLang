package main

import "fmt"

var arr1 [50][50]int
var arr2 [50][50]int
var sum [50][50]int
var size int

func main() {

	getArray()
	addArray()
	displayArray()
}

func getArray() {

	fmt.Print("Enter the size of array : ")
	fmt.Scanln(&size)

	fmt.Println("Enter the values of first  array : ")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Scan(&arr1[i][j])
		}
	}
	fmt.Println("Enter the values of second array : ")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Scan(&arr2[i][j])
		}
	}
}
func addArray() {

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum[i][j] = arr1[i][j] + arr2[i][j]
		}	
	}
}
func displayArray() {

	fmt.Println("Sum of array 1 and array2 ")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum[i][j] = arr1[i][j] + arr2[i][j]
			fmt.Print(sum[i][j], " ")
		}
		fmt.Println()
	}

}
