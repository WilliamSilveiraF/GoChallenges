package main

import (
  "fmt"
)

func main() {
	 var i interface{} = "test"
		switch i.(type) {
		case int: 
			 fmt.Println("i is an integer")
		case string:
			 fmt.Println("i is a string")
		default:
				fmt.Println("i don't know what i is")
		}
}
