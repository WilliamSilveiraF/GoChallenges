package main

import (
	 "fmt"
)

func main() {
	 myResult, err := divide(5.0, 0.0)
		if err != nil {
			 fmt.Println(err)
				return
		}
		fmt.Println(myResult)
}

func divide(dividend, divisor float64) (float64, error) { // 2 result types
		if divisor == 0.0 {
		  return 0.0, fmt.Errorf("Cannot divide by zero")
		}
		return dividend / divisor, nil
}
