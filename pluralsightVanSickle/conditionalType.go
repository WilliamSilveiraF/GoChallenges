package main

import ( "fmt" )

func main() {
		var i interface{} = [3]int{}
		switch i.(type) {
		case int:
				fmt.Println("i is an int")
				break
				fmt.Println("This wont print too")
		case float64:
			 fmt.Println("i is a float64")
		case string:
			 fmt.Println("i is a string")
		case [3]int:
				fmt.Println("i is a array with index 3")
		default:
			 fmt.Println("i is another type")
		}
}