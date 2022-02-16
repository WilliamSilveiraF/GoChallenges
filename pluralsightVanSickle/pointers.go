package main

import (
	 "fmt"
)

func main() {
	 var a int = 42
		var b *int = &a
		fmt.Println(a, b) // b equal memory value
		fmt.Println(&a, b) // &a equal memory value
		fmt.Println(a, *b) // * get the memory value
		
}
