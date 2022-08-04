package main

import "fmt"

func main() {

	var i, j int
	var row int = 5

	for i = 1; i <= row; i++ {
		for j = 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
