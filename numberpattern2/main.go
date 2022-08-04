package main

import "fmt"

func main() {

	var count int

	for i := 0; i < 4; i++ {
		for j := 0; j <= i; j++ {
			count++
			fmt.Print(count, " ")
		}
		fmt.Println()
	}
}
