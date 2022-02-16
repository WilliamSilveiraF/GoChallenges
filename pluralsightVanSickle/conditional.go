package main

import (
	 "fmt"
)

func main() {
		number := 50
		guess := 30

		if returnTrue() {
			 fmt.Println("Don't need put return true here\n")
		}
		if guess < number {
				fmt.Println("Too low")
		} else if guess > number { 
				fmt.Printf("Too high")
		} else {
				fmt.Println("You got it!")
		}
		fmt.Println(number<=guess, number>=guess, number!=guess)
}

func returnTrue() bool {
		fmt.Printf("return true\n")
		return true
}