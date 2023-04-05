package main

import "fmt"

func main() {
	var (
		num1 int
		num2 int
	)
	fmt.Print("ingrese el primer numero: ")
	fmt.Scan(&num1)

	fmt.Print("ingrese el segundo numero: ")
	fmt.Scan(&num2)

	calcular(&num1, &num2)
}

func calcular(a, b *int) {
	fmt.Println("Suma: ", *a+*b)
	fmt.Println("Resta: ", *a-*b)
	fmt.Println("Multiplicacion: ", *a**b)
	fmt.Println("Division: ", *a / *b)
}
