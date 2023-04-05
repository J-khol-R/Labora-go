package main

import "fmt"

func main() {
	num := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nombres := [5]string{"valen", "natalia", "juan", "sebastian", "sofia"}
	imprimirArrayNum(num[:])
	fmt.Println("-----------------------------------")
	imprimirArrayStr(nombres[:])
}

func imprimirArrayNum(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func imprimirArrayStr(arr []string) {
	for i := 0; i < len(arr); i++ {
		fmt.Println((arr[i]))
	}
}
