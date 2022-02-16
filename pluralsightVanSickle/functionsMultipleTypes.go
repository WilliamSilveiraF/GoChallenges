package main 

import (
	 "fmt"
)

func main() {
		sayGreeting("Hello", "Stacey")

		fmt.Println("\n=====\n")

		for i := 0; i < 5; i++ {
			 sayMessage("Hello go!", i)
		}
		
}

func sayMessage(msg string, idx int) {
		fmt.Println(msg)
		fmt.Println("The value of the index is", idx)
}

func sayGreeting(greeting, name string) {
		fmt.Println(greeting, name)
}
