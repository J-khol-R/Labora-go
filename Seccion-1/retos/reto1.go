package main

import "fmt"

func main() {
	mostrarConsola()
}

var maximo = [4]int{50, 50, 600, 800}
var suma [5]int

func mostrarConsola() {
	var numero int
	fmt.Print("ingrese un numero: ")
	fmt.Scan(&numero)
	fmt.Printf("cuando X vale %d su suma es:\n", numero)

	asignarValores(numero, maximo[:])

	for i := 0; i < len(suma); i++ {
		fmt.Printf("s%d = %d\n", i+1, suma[i])
	}
}

func asignarValores(numero int, arr []int) {
	for i := 0; i < len(arr); i++ {

		if numero >= arr[i] {
			suma[i] = arr[i]
			numero -= arr[i]
		} else {
			suma[i] = numero
			break
		}

		if i == (len(arr) - 1) {
			suma[i+1] = numero
		}

	}
}
