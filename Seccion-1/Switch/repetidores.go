package main

import "fmt"

type Planeta struct {
	nombre string
	lunas  int
}

func newPlaneta(nombre string, lunas int) Planeta {
	return Planeta{nombre: nombre, lunas: lunas}
}

func (p Planeta) String() string {
	return fmt.Sprintf("{nombre: %s, edad: %d}", p.nombre, p.lunas)
}

func main() {
	planeta1 := newPlaneta("mercurio", 0)
	planeta2 := newPlaneta("venus", 0)
	planeta3 := newPlaneta("tierra", 1)
	planeta4 := newPlaneta("marte", 2)
	planeta5 := newPlaneta("jupiter", 63)
	planeta6 := newPlaneta("saturno", 62)
	planeta7 := newPlaneta("urano", 27)
	planeta8 := newPlaneta("neptuno", 13)

	planetas := [8]Planeta{planeta1, planeta2, planeta3, planeta4, planeta5, planeta6, planeta7, planeta8}

	for i := 0; i < len(planetas); i++ {
		if planetas[i].lunas > 0 {
			fmt.Printf("nombre: %s, lunas: %d\n", planetas[i].nombre, planetas[i].lunas)
		} else {
			fmt.Printf("nombre: %s, no tiene lunas\n", planetas[i].nombre)
		}

	}
}
