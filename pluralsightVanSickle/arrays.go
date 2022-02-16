package main

import (
	 "fmt"
)

func main() {
		var emptyArray [3]string
		fmt.Printf("Empty Array: %v\n", emptyArray)

		students := [3]string{ "Maria", "João", "Carlos" }
		fmt.Printf("%v\n", students)
		students[0] = "Lisa"
		fmt.Printf("%v\n", students)

		grades := [...]int{ 97, 85, 93 }
		fmt.Printf("Grades: %v\n", grades)
}
