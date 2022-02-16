package main

import (
	 "fmt"
)

func main() {
	 myResult := sum(1, 2, 3, 4, 5)
		fmt.Println("The sum is", *myResult)
}

func sum(myArr ...int) *int {
		fmt.Println(myArr)
		result := 0
		for _, v := range myArr {
			 result += v
		}
		return &result
}
