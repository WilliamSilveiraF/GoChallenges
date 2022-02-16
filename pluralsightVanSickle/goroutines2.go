package main

import (
	 "fmt"
		"time"
)

func main() {
	 var msg = "Hello"
		
		go func() {
			 fmt.Println(msg) // this is reassigned because it was acessing the msg variable
		}()

		go func(msg string) {
			fmt.Println(msg)
		}(msg) // this is not, because i pass on the function parameter scope

		msg = "Goodbye"

		time.Sleep(100 * time.Millisecond)
}
