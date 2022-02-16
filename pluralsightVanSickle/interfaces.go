package main

import (
	 "fmt"
)

func main() {
	 var myInterface Writer = ConsoleWriter{}
		myInterface.Write([]byte("Hello go!"))
}

type Writer interface {
		Write([]byte) (int, error)
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
		myResult, err := fmt.Printf(string(data))

		return myResult, err
}
