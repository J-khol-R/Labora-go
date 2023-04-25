package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Crear el enrutador y definir las rutas en la función definirRutas
	enrutador := mux.NewRouter()

	// aquí definir todas las rutas...
	//enrutador.HandleFunc...
	definirRutas(enrutador)

	// Dirección del servidor. En este caso solo indicamos el puerto
	// pero podría ser algo como "127.0.0.1:8000"
	direccion := ":8000"

	servidor := &http.Server{
		Handler: enrutador,
		Addr:    direccion,
		// Timeouts para evitar que el servidor se quede "colgado" por siempre
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Escuchando en %s. Presiona CTRL + C para salir", direccion)
	log.Fatal(servidor.ListenAndServe())
}

func definirRutas(enrutador *mux.Router) {
	enrutador.HandleFunc("/hola", func(respuesta http.ResponseWriter, peticion *http.Request) {
		respuesta.Write([]byte("¡Hola, mundo!"))
	}).Methods("GET")

	enrutador.HandleFunc("/adios", func(respuesta http.ResponseWriter, peticion *http.Request) {
		respuesta.Write([]byte("adios mundo :c..."))
	}).Methods("GET")

	enrutador.HandleFunc("/gato", func(respuesta http.ResponseWriter, peticion *http.Request) {
		respuesta.Write([]byte("miauuuuu!!!!"))
	}).Methods("GET")

	enrutador.HandleFunc("/perro", func(respuesta http.ResponseWriter, peticion *http.Request) {
		respuesta.Write([]byte("guau guau..."))
	}).Methods("GET")

	enrutador.HandleFunc("/usuarios", obtenerUsuarios).Methods("GET")

	enrutador.HandleFunc("/ventas/{tipo}/{anio}", func(respuesta http.ResponseWriter, peticion *http.Request) {

		variablesDePeticion := mux.Vars(peticion)

		tipo := variablesDePeticion["tipo"]

		anio := variablesDePeticion["anio"]

		respuesta.Write([]byte("El tipo de venta es " + tipo + " y el año es " + anio))

	}).Methods("GET")

	enrutador.HandleFunc("/usuario/{id}", obtenerUsuarioPorId).Methods("GET")

	enrutador.HandleFunc("/pedidos/{anio:[0-9]{4}}", func(respuesta http.ResponseWriter, peticion *http.Request) {

		variablesDePeticion := mux.Vars(peticion)

		anio := variablesDePeticion["anio"]

		respuesta.Write([]byte("Pedidos del año: " + anio))

	}).Methods("GET")

	enrutador.HandleFunc("/libros/{orden:(?:ascendente|descendente)}", func(respuesta http.ResponseWriter, peticion *http.Request) {

		variablesDePeticion := mux.Vars(peticion)

		orden := variablesDePeticion["orden"]

		respuesta.Write([]byte("Libros ordenados: " + orden))

	}).Methods("GET")

	enrutador.HandleFunc("/cursos", func(respuesta http.ResponseWriter, peticion *http.Request) {

		variablesGet := peticion.URL.Query()

		// Cada variable es un arreglo

		orden := variablesGet["orden"]

		// Si mide más que 0 entonces el orden sí está definido

		if len(orden) > 0 {

			fmt.Fprintf(respuesta, "El orden: %s.", orden[0]) // Acceder al primer elemento

		}

		fmt.Fprintf(respuesta, "Parámetros de consulta: %v", variablesGet)

	}).Methods("GET")
}

func obtenerUsuarios(respuesta http.ResponseWriter, peticion *http.Request) {
	// Nota: estos usuarios podrían venir de una base de datos que podrían obtenerse dentro
	// de cada ruta
	type Usuario struct {
		Id     int    `json:"id"`
		Correo string `json:"correo"`
	}

	var usuarios []Usuario = []Usuario{
		Usuario{
			Id:     1,
			Correo: "contacto@parzibyte.me",
		},
		Usuario{
			Id:     2,
			Correo: "john.galt@atlas.com",
		},
	}
	// También podrías codificar otro tipo de datos como un arreglo plano
	// o una simple variable, todo lo soportado por JSON:
	// https://parzibyte.me/blog/2019/05/16/codificar-decodificar-json-go-golang/
	json.NewEncoder(respuesta).Encode(usuarios)
}

func obtenerUsuarioPorId(respuesta http.ResponseWriter, peticion *http.Request) {

	variablesDePeticion := mux.Vars(peticion)

	// El id viene como cadena, hay que convertirlo a entero de 32 bits

	// Aquí "id" es la variable que indicamos en la ruta

	idUsuarioBuscado, err := strconv.Atoi(variablesDePeticion["id"])

	if err != nil {
		http.Error(respuesta, "ID de usuario inválido", http.StatusBadRequest)
		return
	}

	// Imprimir el valor de idUsuarioBuscado en la respuesta HTTP
	respuesta.Write([]byte(fmt.Sprintf("El ID de usuario buscado es %d", idUsuarioBuscado)))

}
