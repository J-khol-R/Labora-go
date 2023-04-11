package main

import "fmt"

var tipoSangre string

func clasificarSangre(tipoSangre string) string {
	var salida string

	switch tipoSangre {
	case "A+":
		salida = fmt.Sprintf("Grupo sanguíneo %s, factor Rh positivo", tipoSangre)
	case "A-":
		salida = fmt.Sprintf("Grupo sanguíneo %s, factor Rh negativo", tipoSangre)
	case "B+":
		salida = fmt.Sprintf("Grupo sanguíneo %s, factor Rh positivo", tipoSangre)
	case "B-":
		salida = fmt.Sprintf("Grupo sanguíneo %s, factor Rh negativo", tipoSangre)
	case "Ab+":
		salida = fmt.Sprintf("Grupo sanguíneo %s, factor Rh positivo", tipoSangre)
	case "AB-":
		salida = fmt.Sprintf("Grupo sanguíneo %s, factor Rh negativo", tipoSangre)
	case "O+":
		salida = fmt.Sprintf("Grupo sanguíneo %s, factor Rh positivo", tipoSangre)
	case "O-":
		salida = fmt.Sprintf("Grupo sanguíneo %s, factor Rh negativo", tipoSangre)
	default:
		salida = fmt.Sprintf("Debe de ingresar un grupo sanguineo valido (ej: AB+)")
	}
	return salida
}

func main() {
	var tipoSangre string
	fmt.Print("ingrese su tipo de sangre: ")
	fmt.Scan(&tipoSangre)
	resultado := clasificarSangre(tipoSangre)
	fmt.Print(resultado)
}
