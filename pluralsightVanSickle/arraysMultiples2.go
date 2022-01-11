package main

import (
		"fmt"
)

func main() {
		var identifyMatrix [3][3]int
		identifyMatrix[0] = [3]int{ 1, 0, 0 }
		identifyMatrix[1] = [3]int{ 0, 1, 0 }
		identifyMatrix[2] = [3]int{ 0, 0, 1 }

		fmt.Println(identifyMatrix)
}
