package main

import (
	 "fmt"
)

func main() {
		var n complex128 = 1 + 2 + 5i
		fmt.Printf("%v, %T\n", real(n), real(n))
		fmt.Printf("%v, %T\n", imag(n), imag(n))

		var createComplexN complex128 = complex(3, 12)
		fmt.Printf("%v, %T\n", createComplexN, createComplexN)
}
