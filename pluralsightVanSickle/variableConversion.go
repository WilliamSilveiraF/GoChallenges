package main

import (
		"fmt"
		"strconv"
	)

func main() {
		var i float64 = 42
		fmt.Printf("%v, %T\n", i, i)

		var j int
		j = int(i)
		fmt.Printf("%v, %T\n", j, j)

		var x string
		x = strconv.Itoa(j)
		fmt.Printf("%v, %T\n", x, x)
}