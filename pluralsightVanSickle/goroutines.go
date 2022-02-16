// go run -race src/main.go

package main

import (
	"fmt"
	"time"
)

func main() {
		go sayHello()
		time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}
