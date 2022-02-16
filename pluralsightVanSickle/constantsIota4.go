package main

import (
		"fmt"
)

const (
		isAdmin = 1 << iota
		isHeadquarters
		canSeeFinancials
		
		canSeeAfrica
		canSeeAsia
		canSeeEurope
		canSeeNorthAmerica
		canSeeSouthAmerica
)

func main() {
		var roles byte = isAdmin | canSeeFinancials | canSeeEurope
		fmt.Printf("%b\n", roles) //truth value = 0010 0101 || 00${canSeeEurope}0 0${canSeeFinancials}0${isAdmin}
		fmt.Printf("IsAdmin? %v\n", isAdmin & roles == isAdmin)
		fmt.Printf("Is HQ? %v\n", isHeadquarters & roles == isAdmin)
}
