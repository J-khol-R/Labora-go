package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items = []Item{
	{"1", "valentina"},
	{"2", "juan"},
	{"3", "maria"},
	{"4", "esteban"},
	{"5", "andrea"},
	{"6", "felipe"},
	{"7", "david"},
	{"8", "tobby"},
	{"9", "carmela"},
	{"10", "luz"},
}

func getItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	variable := mux.Vars(r)
	id := variable["id"]

	var item Item
	var encontrado bool
	for _, valor := range items {
		if valor.ID == id {
			item = valor
			encontrado = true
			break
		}
	}

	if !encontrado {
		w.WriteHeader(http.StatusNotFound) //devuelve el error 404 que es cuando no encuantra algo
		w.Write([]byte("En nuestra base de datos no se encuentra el id " + id))
		return
	}
	json.NewEncoder(w).Encode(item)

}

func createItem(w http.ResponseWriter, r *http.Request) {
	// Función para crear un nuevo elemento
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	// Función para actualizar un elemento existente
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	// Función para eliminar un elemento
}

func main() {

	enrutador := mux.NewRouter()

	enrutador.HandleFunc("/items", getItems).Methods("GET")
	enrutador.HandleFunc("/items/{id}", getItem).Methods("GET")

	direccion := ":3000"

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
