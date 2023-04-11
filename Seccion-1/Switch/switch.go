package main

import "fmt"

func main() {
	var num int
	fmt.Print("ingrese un numero: ")
	fmt.Scan(&num)

	switch num {
	case 1:
		fmt.Print("Lunes")
	case 2:
		fmt.Print("Martes")
	case 3:
		fmt.Print("Miercoles")
	case 4:
		fmt.Print("Jueves")
	case 5:
		fmt.Print("Viernes")
	case 6:
		fmt.Print("Sabado")
	case 7:
		fmt.Print("Domingo")
	default:
		fmt.Print("Ingrese un numero valido del 1 al 10")
	}
}
