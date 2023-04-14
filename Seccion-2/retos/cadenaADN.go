package main

import "fmt"

//como ya tenemos los indices en donde se encuentra las letras iguales a la primera de la cadena original
// vamos a extraer la subcadena desde ese indice, la concatenamos con los sobrantes a su izq y los añadimos a "palabrasArmadas"
func armarPalabras(arrIndices []int, sentencia string) []string {
	var palabrasArmadas []string
	var palabra string
	for i := 0; i < len(arrIndices); i++ {
		palabra = sentencia[arrIndices[i]:] + sentencia[:arrIndices[i]]
		palabrasArmadas = append(palabrasArmadas, palabra)
	}
	return palabrasArmadas
}

// recorremos la cadena de ADN corrida buscando las letras iguales a la primer letra de la cadena original
// y guardamos sus indices en un slice
func recorrerCorrida(valor string, sentencia string) []int {
	var cantLetrasEncontradas []int
	for indice, letra := range sentencia {
		if valor == string(letra) {
			cantLetrasEncontradas = append(cantLetrasEncontradas, indice)
		}
	}
	return cantLetrasEncontradas
}

// comparamos las palabras que se encuentran en el slice que retorna la funcion "armarPalabras()"
// y si ninguna de esas combinaciones es igual a la original podemos decir que la cadena de ADN no es la misma
func comparar(original, corrida string) bool {
	encontrada := false
	// añadimos en una variable la primer letra de la palabra original
	letra := string(original[0])

	//le asignamos a una variable el slice resultante con los indices donde se encuentra "letra"
	arrIndices := recorrerCorrida(letra, corrida)

	//armamos las combinaciones de palabras que empiezan con la inicial "letra"
	arrCombinaciones := armarPalabras(arrIndices, corrida)

	//comparamos las palabras con la original
	for _, palabra := range arrCombinaciones {
		if palabra == original {
			encontrada = true
			// fmt.Print(palabra+"-", original)
			break
		}
	}

	return encontrada
}

func main() {
	fmt.Println(comparar("TAGHAT", "ATTAGH"))      //TRUE
	fmt.Println(comparar("ATACGAGT", "CGAGTATA"))  //TRUE
	fmt.Println(comparar("GCTAACGT", "GTGCTAAC"))  //TRUE
	fmt.Println(comparar("GGCTAACGT", "TAACTGCA")) //FALSE
	fmt.Println(comparar("ATFTCATG", "FTCATGAT"))  //TRUE
	fmt.Println(comparar("ATFTCATG", "HHHHHHHH"))  //FALSE
}
