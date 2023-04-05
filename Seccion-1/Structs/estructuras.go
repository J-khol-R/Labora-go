package main

import (
	"fmt"
)

func main() {

	persona1 := Persona{nombre: "valentina", edad: 20, ciudad: "Cali", telefono: 1234}
	persona2 := Persona{nombre: "juan", edad: 30, ciudad: "Bogota", telefono: 5634}

	fmt.Printf("persona1: %v\n", persona1)
	fmt.Printf("persona2: %v\n", persona2)

	cambiarCiudad("Armenia", &persona1)
	fmt.Printf("persona1 con ciudad actualizada: %v\n", persona1)

	cambiarCiudad("Bogota", &persona2)
	fmt.Printf("persona2: %v", persona2)

}

type Persona = struct {
	nombre   string
	edad     int
	ciudad   string
	telefono int
}

func cambiarCiudad(ciudad string, p *Persona) {
	if p.ciudad != ciudad {
		p.ciudad = ciudad
	}
}
