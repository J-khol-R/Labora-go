package imprimir

import (
	"fmt"
)

func ImprimirMenu(opciones []string) {
	fmt.Printf("===== %s ======\n", opciones[0])
	for i := 1; i < len(opciones); i++ {
		fmt.Printf("%d. %s \n", i, opciones[i])
	}
	fmt.Print("ingrese su opcion: ")
}
