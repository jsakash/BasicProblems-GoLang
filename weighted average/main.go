package main

import "fmt"

func main() {

	var test, lab, assignment, grade float32

	fmt.Println("Enter the marks scored by the students")
	fmt.Print("Written test = ")
	fmt.Scan(&test)
	fmt.Print("Lab exams = ")
	fmt.Scan(&lab)
	fmt.Print("Assignments = ")
	fmt.Scan(&assignment)

	grade = (test*70)/100 + (lab*20)/100 + (assignment*10)/100

	fmt.Print("Grade of the student = ", grade)
	fmt.Println()
}
