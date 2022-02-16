package main

import (
		"fmt"
)

func main() { // Binary operation
		a := 10 // 1010
	 b := 3 // 0011
		fmt.Println(a & b) // 0010 - VHDL EXAMPLE: PORT AND
		fmt.Println(a | b) // 1011 - VHDL EXAMPLE: PORT OR
		fmt.Println(a ^ b) // 1001 - VHDL EXAMPLE: PORT XOR
		fmt.Println(a &^ b) // 0100 - VHDL EXAMPLE: PORT NAND
}