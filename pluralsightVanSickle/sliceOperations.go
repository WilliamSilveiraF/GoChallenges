package main

import (
	 "fmt"
)

func main() {
		a := []int{1, 2, 3, 4, 5}
		begin := a[:1]
		end := a[len(a) - 1:]

		myConcat := append(end, a[0:2]...)

		fmt.Printf("Begin: %v\nEnd: %v\nConcat: %v", begin, end, myConcat)
}
