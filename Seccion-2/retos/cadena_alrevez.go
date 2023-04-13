package main

import "fmt"

func main() {
	harcodeado := "valentina"
	for _, letra := range harcodeado {
		defer fmt.Print(string(letra))
	}

}
