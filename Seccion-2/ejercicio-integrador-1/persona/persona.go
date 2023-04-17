package persona

import (
	"fmt"
	"sort"
)

type Estudiante struct {
	Nombre string
	Nota   float64
	Codigo string
}

func OrdenarEstudiantes(ascendente bool, criterio string, estudiantes []Estudiante) {
	var comparar func(a, b Estudiante) bool
	switch criterio {
	case "1":
		comparar = func(a, b Estudiante) bool {
			if ascendente {
				return a.Nombre < b.Nombre
			}
			return a.Nombre > b.Nombre
		}
	case "2":
		comparar = func(a, b Estudiante) bool {
			if ascendente {
				return a.Nota < b.Nota
			}
			return a.Nota > b.Nota
		}
	case "3":
		comparar = func(a, b Estudiante) bool {
			if ascendente {
				return a.Codigo < b.Codigo
			}
			return a.Codigo > b.Codigo
		}
	default:
		fmt.Println("Criterio de ordenamiento no válido.")
		return
	}

	if ascendente {
		sort.Slice(estudiantes, func(i, j int) bool {
			return comparar(estudiantes[i], estudiantes[j])
		})
	} else {
		sort.Slice(estudiantes, func(i, j int) bool {
			return comparar(estudiantes[j], estudiantes[i])
		})
	}
}

func CrearEstudiante(estudiantes *[]Estudiante) {
	fmt.Println("======= Crear Estudiante =========")
	var estudiante Estudiante
	fmt.Println("Nombre: ")
	fmt.Scan(&estudiante.Nombre)

	fmt.Println("Nota: ")
	fmt.Scan(&estudiante.Nota)

	fmt.Println("Código: ")
	fmt.Scan(&estudiante.Codigo)

	*estudiantes = append(*estudiantes, estudiante)
}

func BuscarPorCodigo(estudiantes []Estudiante) {
	var cod string
	fmt.Print("Ingresa el código del estudiante: ")
	fmt.Scan(&cod)
	for _, estudiante := range estudiantes {
		if estudiante.Codigo == cod {
			fmt.Printf("\nCódigo: %s - Nombre: %s, Nota: %.2f\n", estudiante.Codigo, estudiante.Nombre, estudiante.Nota)
			break
		}
	}
}

func ImprimirEstudiantes(estudiantes []Estudiante) {
	for i, estudiante := range estudiantes {
		fmt.Printf("===== Estudiante %d ======\n", i+1)
		fmt.Println("Nombre:", estudiante.Nombre)
		fmt.Println("Nota:", estudiante.Nota)
		fmt.Println("Código:", estudiante.Codigo)
	}
}
