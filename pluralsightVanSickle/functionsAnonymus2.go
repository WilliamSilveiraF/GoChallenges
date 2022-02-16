package main

import (
	 "fmt"
)

func main() {
		// divide(5.0, 3.0) Cannot call the function here, because it's declared variable
		var divide func(float64, float64) (float64, error)
		divide = func(dividend, divisor float64) (float64, error) {
				if divisor == 0.0 {
					 return 0.0, fmt.Errorf("Cannot divide by zero")
				}
				return dividend / divisor, nil
		}
		myResult, err := divide(5.0, 0.0)
		if err != nil {
			 fmt.Println(err)
				return
		}
		fmt.Println(myResult)
}
