package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items []Item

func getItems(w http.ResponseWriter, r *http.Request) {

	recibir_page := r.URL.Query().Get("page")
	recibir_itemsPerPage := r.URL.Query().Get("itemsPerPage")

	if recibir_page == "" {
		recibir_page = "1"
	}
	if recibir_itemsPerPage == "" {
		recibir_itemsPerPage = "10"
	}

	page, _ := strconv.Atoi(recibir_page)
	itemsPerPage, _ := strconv.Atoi(recibir_itemsPerPage)

	inicio := (page - 1) * itemsPerPage
	fin := inicio + itemsPerPage
	if fin > len(items) {
		fin = len(items)
	}
	elementosPagina := items[inicio:fin]

	// Devolver los elementos correspondientes a la página
	json.NewEncoder(w).Encode(elementosPagina)
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

func getItemByName(w http.ResponseWriter, r *http.Request) {
	variable := mux.Vars(r)
	name := variable["name"]

	var names []Item
	var encontrado bool
	for _, valor := range items {
		if strings.EqualFold(valor.Name, name) {
			names = append(names, valor)
			encontrado = true
			break
		}
	}

	if !encontrado {
		w.WriteHeader(http.StatusNotFound) //devuelve el error 404 que es cuando no encuantra algo
		return
	}
	json.NewEncoder(w).Encode(names)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	variable := mux.Vars(r)
	id := variable["id"]

	var datos map[string]string
	json.NewDecoder(r.Body).Decode(&datos)

	for i, item := range items {
		if item.ID == id {
			items[i].ID = datos["id"]
			items[i].Name = datos["name"]
			break
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("El item se actualizo correctamente :)"))

}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	variable := mux.Vars(r)
	id := variable["id"]

	var datos map[string]string
	json.NewDecoder(r.Body).Decode(&datos)

	for i, item := range items {
		if item.ID == id {
			items = append(items, items[i+1:]...)
			break
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("el item se elimino correctamente :)"))

}

func main() {
	for i := 1; i <= 50; i++ {
		items = append(items, Item{ID: fmt.Sprint(i), Name: fmt.Sprintf("Item %d", i)})
	}

	enrutador := mux.NewRouter()

	enrutador.HandleFunc("/items", getItems).Methods("GET")
	enrutador.HandleFunc("/items/{id}", getItem).Methods("GET")
	enrutador.HandleFunc("/items/nombre/{name}", getItemByName).Methods("GET")
	enrutador.HandleFunc("/items", createItem).Methods("POST")
	enrutador.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	enrutador.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

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
