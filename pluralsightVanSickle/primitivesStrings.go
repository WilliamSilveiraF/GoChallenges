package main

import (
	 "fmt"
)

func main() {
		str :=  "this is a string"
		fmt.Printf("%v, %T\n", str, str)
		fmt.Printf("%v, %T\n", string(str[2]), str[2])

		str2  := "this is also a string"
		fmt.Printf("%v, %T\n", str + str2, str + str2)

		forByte := []byte(str)
		fmt.Printf("%v, %T\n", forByte, forByte)
}
