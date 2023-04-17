package main

import (
	"fmt"
	i "funciones/imprimir"
	p "funciones/persona"
)

func encontrarMaxMin(estudiante []p.Estudiante) {
	var max, min float64
	var nombreMax, nombreMin string
	max = estudiante[0].Nota
	nombreMax = estudiante[0].Nombre
	min = estudiante[0].Nota
	nombreMin = estudiante[0].Nombre

	for i := 1; i < len(estudiante); i++ {
		if estudiante[i].Nota > max {
			max = estudiante[i].Nota
			nombreMax = estudiante[i].Nombre
		}
		if estudiante[i].Nota < min {
			min = estudiante[i].Nota
			nombreMin = estudiante[i].Nombre
		}
	}

	fmt.Printf("El estudiante con la nota mas alta es: %s: %f\n", nombreMax, max)
	fmt.Printf("El estudiante con la nota mas baja es: %s: %f\n", nombreMin, min)
}

func promedioNotas(lista []p.Estudiante) {
	var suma, promedio float64

	for _, estudiante := range lista {
		suma += estudiante.Nota
	}

	promedio = suma / float64(len(lista))
	fmt.Printf("la nota promedio de los estudiantes es: %f\n", promedio)
}

func menuOrdenar(opcion string, estudiantes []p.Estudiante) {
	opMenu := []string{"ORDENAR", "ordenar por nombre", "ordenar por nota", "ordenar por codigo"}
	switch opcion {
	case "1":
		i.ImprimirMenu(opMenu)
		fmt.Scan(&opcion)
		p.OrdenarEstudiantes(true, opcion, estudiantes)
	case "2":
		i.ImprimirMenu(opMenu)
		fmt.Scan(&opcion)
		p.OrdenarEstudiantes(false, opcion, estudiantes)
	}
}

func main() {

	var estudiantes []p.Estudiante
	var opcion string
	opcionesPrincipal := []string{"GESTION ESTUDIANTES", "Crear estudiantes", "Ordenar estudiantes", "Buscar estudiantes", "listar estudiantes", "Salir"}

	defer func() {
		fmt.Println("----- INFORMACION -----")
		fmt.Println("Estos fueron los estudiantes que tuvieron la nota max y minima: ")
		encontrarMaxMin(estudiantes)
		fmt.Println("Esate fue el promedio de las notas de todos los estudiantes: ")
		promedioNotas(estudiantes)
	}()

	for {
		i.ImprimirMenu(opcionesPrincipal)
		fmt.Scan(&opcion)

		switch opcion {
		case "1":
			p.CrearEstudiante(&estudiantes)
		case "2":
			opOrdenar := []string{"ORDENAR", "Ordenar en forma ascendente", "Ordenar en forma descendente"}
			i.ImprimirMenu(opOrdenar)
			fmt.Scan(&opcion)
			menuOrdenar(opcion, estudiantes)
			fmt.Println("lista ordenada")
			p.ImprimirEstudiantes(estudiantes)
		case "3":
			p.BuscarPorCodigo(estudiantes)
		case "4":
			p.ImprimirEstudiantes(estudiantes)
		case "5":
			return
		default:
			fmt.Print("la opcion ingresada no es valida")
		}
	}

}
