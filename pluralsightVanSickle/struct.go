package main

import (
	 "fmt"
)

type Doctor struct {
		number int
		actorName string
		positionMakeDifference []string
		companions []string
}

func main() {
		
	aDoctor := Doctor {
			 number: 3,
				actorName: "Jon Pertwee",
				positionMakeDifference: []string {},
				companions: []string {
					 "Liz Shaw",
						"Jo Grant",
						"Sarah Jane Smith",
				},
		}
		fmt.Println(aDoctor)
}
