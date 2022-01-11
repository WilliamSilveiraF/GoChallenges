package main

import (
	 "fmt"
)

func main() {
		a := []string{}
		fmt.Println(a)
		fmt.Printf("Length: %v\n", len(a))
		fmt.Printf("Capacity: %v\n", cap(a))

		a = append(a, "first")
		a = append(a, "second", "third", "forth")

		fmt.Println(a)
		fmt.Printf("Length: %v\n", len(a))
		fmt.Printf("Capacity: %v\n", cap(a))

		b := []int{}
		fmt.Println(b)
		fmt.Printf("Length: %v\n", len(b))
		fmt.Printf("Capacity: %v\n", cap(b))

		b = append(b, 1)
		b = append(b, 2, 3, 4, 5) // Note that capacity is added by 2-2

		fmt.Println(b)
		fmt.Printf("Length: %v\n", len(b))
		fmt.Printf("Capacity: %v\n", cap(b))
}
