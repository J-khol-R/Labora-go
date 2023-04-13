package main

import (
	"fmt"
	"strings"
)

var (
	coins   = 50
	vocales = []string{"a", "e", "i", "o", "u"}
	users   = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func contarVocales(nombre string) int {
	Vocales := 0
	for _, letra := range strings.ToLower(nombre) {
		switch string(letra) {
		case "a":
			Vocales += 1
		case "e":
			Vocales += 1
		case "i":
			Vocales += 2
		case "o":
			Vocales += 3
		case "u":
			Vocales += 4
		}
	}
	return Vocales
}

func asignacionBTC(usuarios []string, monedas int, distribucion map[string]int) {
	for _, usuario := range usuarios {
		bitcoinAsignadas := 0
		vocales := contarVocales(usuario)

		if vocales > 10 {
			bitcoinAsignadas = 10
		} else {
			bitcoinAsignadas = vocales
		}

		if monedas-bitcoinAsignadas >= 0 {
			distribucion[usuario] = bitcoinAsignadas
			monedas -= bitcoinAsignadas
		}
	}
}

func imprimirMapa(mapa map[string]int) {
	for nombre, monedas := range mapa {
		fmt.Printf("A %s se le asignaron %d BTC\n", nombre, monedas)
	}
}

func main() {
	asignacionBTC(users, coins, distribution)
	imprimirMapa(distribution)
	fmt.Println("Coins left:", coins)
}
