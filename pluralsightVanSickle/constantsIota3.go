package main

import ( "fmt" )

const (
		_ = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
)

func main() {
		fmt.Printf("%v\n", KB)
		fmt.Printf("%v\n", MB)
		fmt.Printf("%v\n", GB)
		fmt.Printf("%v\n", TB)
		fmt.Printf("%v\n", PB)
		fmt.Printf("%v\n", EB)
}
