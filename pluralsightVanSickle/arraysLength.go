package main

import ( 
		"fmt" 
)

func main() {
		var students [4]string
		fmt.Printf("Students: %v\n", students)

		students[0] = "Lisa"
		students[2] = "Ahmed"
		students[1] = "Arnold"

		fmt.Printf("Student: %v\n", students[3])
		fmt.Printf("Student: %v\n", students[2])
		fmt.Printf("Number of Students: %v\n", len(students))
}