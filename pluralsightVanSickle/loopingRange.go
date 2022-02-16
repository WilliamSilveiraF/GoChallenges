package main

import (
	 "fmt"
)

func main() {
		s := []int{1, 2, 3}
		for myIndex, myValue := range s {
			 fmt.Println(myIndex, myValue)
		}
}