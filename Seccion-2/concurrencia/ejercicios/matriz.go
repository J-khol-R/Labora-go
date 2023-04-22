package main

import (
	"fmt"
)

func multiplicarFilaColumna(fila []int, matriz [][]int, c chan []int, indice int) {
	var resultado []int
	var suma int
	for i := 0; i < len(matriz[i]); i++ {
		for j := 0; j < len(matriz); j++ {
			dato := fila[j] * matriz[j][i]
			suma += dato
		}
		resultado = append(resultado, suma)
		suma = 0
	}
	resultado = append(resultado, indice)
	c <- resultado
}

func asignarRutina(matriz1 [][]int, matriz2 [][]int, c chan []int) {
	for i, _ := range matriz1 {
		go multiplicarFilaColumna(matriz1[i], matriz2, c, i)
	}

}

func acomodarMatriz(matriz1 [][]int, matriz2 [][]int, c chan []int) [][]int {
	asignarRutina(matriz1, matriz2, c)

	matriz := make([][]int, len(matriz1))

	for i := 0; i < len(matriz); i++ {
		dato := <-c
		matriz[dato[len(dato)-1]] = dato[:len(dato)-1]
	}

	return matriz

}

func imprimirMatriz(matriz [][]int) {
	for i, _ := range matriz {
		for _, fila := range matriz[i] {
			fmt.Print(fila, " ")
		}
		fmt.Print("\n")
	}
}

func main() {
	c := make(chan []int)
	matriz1 := [][]int{{1, 2, 3}, {6, 7, 8}}
	matriz2 := [][]int{{5, 7}, {3, 1}, {4, 9}}
	matriz := acomodarMatriz(matriz1, matriz2, c)
	fmt.Println("MATRIZ RESULTANTE:")
	imprimirMatriz(matriz)
}
