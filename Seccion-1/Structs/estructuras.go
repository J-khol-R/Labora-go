package main

import (
	"fmt"
)

type Persona struct {
	nombre   string
	edad     int
	ciudad   string
	telefono int
}

func NewPersona(nombre string, edad int, ciudad string, telefono int) Persona {
	return Persona{nombre: nombre, edad: edad, ciudad: ciudad, telefono: telefono}
}

func (p *Persona) cambiarCiudad(ciudad string) {
	if p.ciudad != ciudad {
		p.ciudad = ciudad
	}
}

func (p Persona) String() string {
	return fmt.Sprintf("{nombre: %s, edad: %d, ciudad: %s, telefono: %d}", p.nombre, p.edad, p.ciudad, p.telefono)
}

func main() {
	persona1 := NewPersona("valentina", 20, "Cali", 1234)
	persona2 := NewPersona("juan", 30, "Bogota", 5634)

	fmt.Printf("persona1: %v\n", persona1)
	fmt.Printf("persona2: %v\n", persona2)

	persona1.cambiarCiudad("Armenia")
	fmt.Printf("persona1 con ciudad actualizada: %v\n", persona1)

	persona2.cambiarCiudad("Bogota")
	fmt.Printf("persona2: %v\n", persona2)
}
