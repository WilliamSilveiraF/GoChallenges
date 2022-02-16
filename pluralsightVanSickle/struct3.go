package main

import (
  "fmt"
)

type Animal struct {
		Name string
		Origin string
}

type Bird struct {
		Animal
	 SpeedKPH float32
		CanFly bool
}

func main() {
		beijaFlor := Bird{}
		beijaFlor.Name = "Beija Flor Verde"
		beijaFlor.Origin = "Brasil"

		beijaFlor.CanFly = true

		fmt.Println(beijaFlor.Name)

		beijaFlor2 := Bird{
			 Animal: Animal{Name: "Beija Flor Amarelo", Origin: "Argentina"},
				SpeedKPH: 48,
				CanFly: true,
		}

		fmt.Println(beijaFlor2)
}
