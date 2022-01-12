// revisar

package main

import "fmt"
import "log"

func main() {
		fmt.Println("start")
		defer func() {
				if err := recover(); err != nil {
						log.Println("Error:", err)
				}
		}()
		panic("something bad happened")
		fmt.Println("end")
}
