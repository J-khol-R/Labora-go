package main

import "fmt"

func main() {
	a := 10
	b := 20

	var ptrA *int = &a
	*ptrA, b = b, *ptrA

	fmt.Print(a, b)

}
