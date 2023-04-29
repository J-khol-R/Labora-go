package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ItemDetails struct {
	Item
	Details string `json:"details"`
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

	err := json.NewEncoder(w).Encode(elementosPagina)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}
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
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("En nuestra base de datos no se encuentra el id " + id))
		return
	}
	err := json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

}

func getItemDetails(w http.ResponseWriter, r *http.Request) {
	detailsChannel := make(chan ItemDetails, len(items))

	var wg sync.WaitGroup
	for _, item := range items {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			getDetails(id, detailsChannel)
		}(item.ID)
	}
	wg.Wait()

	var detailedItems []ItemDetails
	for i := 0; i < len(items); i++ {
		itemDetail := <-detailsChannel
		detailedItems = append(detailedItems, itemDetail)
	}

	err := json.NewEncoder(w).Encode(detailedItems)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

}

func getDetails(id string, c chan ItemDetails) {
	time.Sleep(100 * time.Millisecond)
	var foundItem Item
	for _, item := range items {
		if item.ID == id {
			foundItem = item
			break
		}
	}

	c <- ItemDetails{
		Item:    foundItem,
		Details: fmt.Sprintf("Detalles para el item %s", id),
	}
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
	err := json.NewEncoder(w).Encode(names)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Error al decodificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	items = append(items, item)
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	variable := mux.Vars(r)
	id := variable["id"]

	var datos map[string]string
	err := json.NewDecoder(r.Body).Decode(&datos)
	if err != nil {
		http.Error(w, "Error al decodificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

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
	err := json.NewDecoder(r.Body).Decode(&datos)
	if err != nil {
		http.Error(w, "Error al decodificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

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
	enrutador.HandleFunc("/items/details", getItemDetails).Methods("GET")
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
