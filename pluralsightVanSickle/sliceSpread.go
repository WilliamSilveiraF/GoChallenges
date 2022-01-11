package main

import (
	 "fmt"
)

func main() {
		myArr := []int{1, 3, 5, 7, 9}
		fmt.Println(myArr)

		myArr = append(myArr, []int{2, 4, 6, 8}...)
		fmt.Println(myArr)
}