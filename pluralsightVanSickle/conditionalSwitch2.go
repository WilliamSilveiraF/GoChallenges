package main

import (
	 "fmt"
)

func main() {
		i := 10
		switch {
		case i <= 10:
		  fmt.Println("I will be executed")
		case i <= 20:
				fmt.Println("I no")
		default:
				fmt.Println("I no")
		}
		fmt.Println("\n=====\n")
		usingFallthrough := 10
		switch {
		case usingFallthrough <= 3:
			 fmt.Println("I will be executed")
				fallthrough
		case usingFallthrough <= 10:
			 fmt.Println("I will be executed")
				fallthrough
		case usingFallthrough <= 20:
			 fmt.Println("I will be executed")
				fallthrough
		default:
			 fmt.Println("I no")
		}
}