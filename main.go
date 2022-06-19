package main

import (
	"fmt"
	"intefacesGolang/users"
)

func main() {
	var juan users.Cashier = users.NewEmployee("Juan")
	var juan2 users.Admin = users.NewEmployee("Juan2")

	total := juan.CalcTotal(90, 65, 92.6)
	fmt.Println(total)

	juan2.DeactivateEmployee(juan)

	fmt.Println(juan.CalcTotal(90, 65, 92.6))
}
