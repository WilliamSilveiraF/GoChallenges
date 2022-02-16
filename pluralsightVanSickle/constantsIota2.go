package main

import (
		"fmt"
)

const (
		_ = iota + 5
		catSpecialist
		dogSpecialist
		snakeSpecialist
)

func main() {
		var specialistType int
		fmt.Printf("%v\n", specialistType == catSpecialist)
		fmt.Printf("%v\n", catSpecialist)
		fmt.Printf("%v\n", dogSpecialist)
		fmt.Printf("%v\n", snakeSpecialist)
}
