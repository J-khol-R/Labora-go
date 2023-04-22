package main

import (
	"fmt"
)

func sumar(rango1, rango2 int, resultado chan<- int) {
	total := 0
	for i := rango1; i <= rango2; i++ {
		total += i
	}
	resultado <- total
}

func main() {
	rango1, rango2 := 1, 100
	dividir := 2
	resultado := make(chan int, dividir)

	for i := 0; i < dividir; i++ {
		parte1 := rango1 + (i * (rango2 - rango1 + 1) / dividir)
		parte2 := rango1 + ((i + 1) * (rango2 - rango1 + 1) / dividir) - 1
		go sumar(parte1, parte2, resultado)
	}

	total := 0
	for i := 0; i < dividir; i++ {
		total += <-resultado
	}

	fmt.Printf("La suma de los números de %d a %d es %d.\n", rango1, rango2, total)
}
