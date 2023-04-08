package main

import "fmt"

func main() {
	var numeroSegundos int
	fmt.Print("ingrese el numero en segundos a convertir: ")
	fmt.Scan(&numeroSegundos)

	horas, minutos, segundos := convertir(numeroSegundos)
	fmt.Printf("Conversion:\n %d horas %d minutos %d segundos", horas, minutos, segundos)

}

func convertir(num int) (int, int, int) {
	horas := num / 3600
	minutos := (num % 3600) / 60
	segundos := num % 60

	return horas, minutos, segundos
}
