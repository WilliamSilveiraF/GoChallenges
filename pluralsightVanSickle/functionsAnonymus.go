package main

import (
	 "fmt"
)

func main() {
	 for i := 0; i < 5; i++ {
				func(i int) { //best practice
					 fmt.Println(i)
				}(i)
		}
}