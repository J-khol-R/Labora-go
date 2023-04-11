package main

import (
	"fmt"
	"sort"
)

func main() {
	persona1 := NewPersona("valentina", 21, 153, 63)
	persona2 := NewPersona("juan", 22, 167, 62)
	persona3 := NewPersona("andrea", 1, 173, 61)
	persona4 := NewPersona("lorena", 61, 183, 20)

	personas := append(personas, persona1, persona2, persona3, persona4)

	ordenarPersonas(2, personas)
	fmt.Println(personas)

}

var (
	nombre   string
	edad     int
	altura   int
	peso     int
	personas []Persona
)

type Persona struct {
	nombre string
	edad   int
	altura int
	peso   int
}

func NewPersona(nombre string, edad int, altura int, peso int) Persona {
	return Persona{nombre: nombre, edad: edad, altura: altura, peso: peso}
}

func ordenarPersonas(criterio int, personas []Persona) {
	switch criterio {
	case 1:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].nombre < personas[j].nombre
		})
	case 2:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].edad < personas[j].edad
		})
	case 3:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].altura < personas[j].altura
		})
	case 4:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].peso < personas[j].peso
		})
	default:
		fmt.Print("no se cumplio ninguno")
	}

}

func buscarPersona(persona []Persona, nombre string) Persona {
	var encontrada Persona
	for i := 0; i < len(persona); i++ {
		if persona[i].nombre == nombre {
			encontrada = persona[i]
			break
		}
	}

	return encontrada
}

func (p *Persona) IMC() int {
	resultado := p.peso / (p.altura * p.altura)
	return resultado
}
