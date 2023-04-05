package main

import "fmt"

func main() {
	nombres := [10]string{"Avatar", "Creed", "Mario Bros", "Ant Man", "Viaje al centro de la tierra", "Green Book", "Cenicienta", "Bella Durmiente", "Enredados", "Jurasic world"}
	imprimirArray(nombres[:])
	fmt.Printf("El segundo elemento del array es: %s\n", nombres[1])
	fmt.Printf("La longitud del array es: %d", len(nombres))

}
func imprimirArray(arr []string) {
	for i := 0; i < len(arr); i++ {
		fmt.Println((arr[i]))
	}
}
