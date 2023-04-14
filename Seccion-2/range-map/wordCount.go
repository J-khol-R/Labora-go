package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var cuenta map[string]int
	cuenta = make(map[string]int)

	listaPalabras := strings.Fields(s)

	for _, valor := range listaPalabras {
		cuenta[valor] += 1
	}

	return cuenta
}

func main() {
	wc.Test(WordCount)
}
