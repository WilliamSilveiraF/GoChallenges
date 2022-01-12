package main
import (
	 "fmt"
)

func main() {
		defer fmt.Println("start?")
		fmt.Println("middle")
		fmt.Println("end")
}