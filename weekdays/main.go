package main

import "fmt"

func main() {

	var day int

	fmt.Print("Enter Day Number : ")
	fmt.Scanln(&day)

	switch day {
	case 1:
		fmt.Println("SUNDAY")
		break
	case 2:
		fmt.Println("MONDAY")
		break
	case 3:
		fmt.Println("TUESDAY")
		break
	case 4:
		fmt.Println("WEDNESDAY")
		break
	case 5:
		fmt.Println("THURSDAY")
		break
	case 6:
		fmt.Println("FRIDAY")
		break
	case 7:
		fmt.Println("SATURDAY")
		break
	default:
		fmt.Println("INVALID ENTRY!!!!")

	}
}
