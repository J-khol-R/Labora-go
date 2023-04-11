package main

import (
	"fmt"
	"sort"
)

type Estudiante struct {
	nombre string
	nota   float64
	codigo string
}

func ordenarEstudiantesA(criterio string, estudiantes []Estudiante) {
	switch criterio {
	case "1":
		sort.Slice(estudiantes, func(i, j int) bool {
			return estudiantes[i].nombre < estudiantes[j].nombre
		})
	case "2":
		sort.Slice(estudiantes, func(i, j int) bool {
			return estudiantes[i].nota < estudiantes[j].nota
		})
	case "3":
		sort.Slice(estudiantes, func(i, j int) bool {
			return estudiantes[i].codigo < estudiantes[j].codigo
		})

	default:
		fmt.Print("no se cumplio ninguno")
	}

}

func ordenarEstudiantesD(criterio string, estudiantes []Estudiante) {
	switch criterio {
	case "1":
		sort.Slice(estudiantes, func(i, j int) bool {
			return estudiantes[i].nombre > estudiantes[j].nombre
		})
	case "2":
		sort.Slice(estudiantes, func(i, j int) bool {
			return estudiantes[i].nota > estudiantes[j].nota
		})
	case "3":
		sort.Slice(estudiantes, func(i, j int) bool {
			return estudiantes[i].codigo > estudiantes[j].codigo
		})
	default:
		fmt.Print("no se cumplio ninguno")
	}

}

func menuOrdenar(opcion string, estudiantes []Estudiante) {
	opMenu := []string{"ORDENAR", "ordenar por nombre", "ordenar por nota", "ordenar por codigo"}
	switch opcion {
	case "1":
		imprimirMenu(opMenu)
		fmt.Scan(&opcion)
		ordenarEstudiantesA(opcion, estudiantes)
	case "2":
		imprimirMenu(opMenu)
		fmt.Scan(&opcion)
		ordenarEstudiantesD(opcion, estudiantes)
	}
}

func crearEstudiante(estudiantes *[]Estudiante) {
	fmt.Println("======= Crear Estudiante =========")
	var estudiante Estudiante
	fmt.Println("Nombre: ")
	fmt.Scan(&estudiante.nombre)

	fmt.Println("Nota: ")
	fmt.Scan(&estudiante.nota)

	fmt.Println("Código: ")
	fmt.Scan(&estudiante.codigo)

	*estudiantes = append(*estudiantes, estudiante)
}

func imprimirMenu(opciones []string) {
	fmt.Printf("===== %s ======\n", opciones[0])
	for i := 1; i < len(opciones); i++ {
		fmt.Printf("%d. %s \n", i, opciones[i])
	}
	fmt.Print("ingrese su opcion: ")
}

func imprimirEstudiantes(estudiantes []Estudiante) {
	for i, estudiante := range estudiantes {
		fmt.Printf("===== Estudiante %d ======\n", i+1)
		fmt.Println("Nombre:", estudiante.nombre)
		fmt.Println("Nota:", estudiante.nota)
		fmt.Println("Código:", estudiante.codigo)
	}
}

func buscarPorCodigo(estudiantes []Estudiante) {
	var cod string
	fmt.Print("Ingresa el código del estudiante: ")
	fmt.Scan(&cod)
	for _, estudiante := range estudiantes {
		if estudiante.codigo == cod {
			fmt.Printf("\nCódigo: %s - Nombre: %s, Nota: %.2f\n", estudiante.codigo, estudiante.nombre, estudiante.nota)
			break
		}
	}
}

func main() {

	var estudiantes []Estudiante
	var opcion string
	opcionesPrincipal := []string{"GESTION ESTUDIANTES", "Crear estudiantes", "Ordenar estudiantes", "Buscar estudiantes", "listar estudiantes", "Salir"}

	for {
		imprimirMenu(opcionesPrincipal)
		fmt.Scan(&opcion)

		switch opcion {
		case "1":
			crearEstudiante(&estudiantes)
		case "2":
			opOrdenar := []string{"ORDENAR", "Ordenar en forma ascendente", "Ordenar en forma descendente"}
			imprimirMenu(opOrdenar)
			fmt.Scan(&opcion)
			menuOrdenar(opcion, estudiantes)
			fmt.Println("lista ordenada")
			imprimirEstudiantes(estudiantes)
		case "3":
			buscarPorCodigo(estudiantes)
		case "4":
			imprimirEstudiantes(estudiantes)
		case "5":
			return
		default:
			fmt.Print("la opcion ingresada no es valida")
		}
	}

}
