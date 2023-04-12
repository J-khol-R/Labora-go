package main

import (
	"fmt"
	"strings"
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
	fmt.Print(WordCount("I ate a donut. Then I ate another donut."))
}