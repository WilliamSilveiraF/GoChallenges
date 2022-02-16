package main

import "fmt"

func main() {
		var i int
		i = 42
		var j float32 = 27
		k := 99.

		fmt.Println(i, j, k)
		fmt.Printf("%v, %T", k, k) //Format value(%v) and type(%T)
}
