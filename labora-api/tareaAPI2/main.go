package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ItemDetails struct {
	Item
	Details string `json:"details"`
}

const (
	host   = "localhost"
	port   = "5432"
	dbName = "labora_proyect_1"

	rolName     = "postgres"
	rolPassword = "0b3j1t4,"
)

type manzana struct {
	Id           int
	CustomerName string
	OrderDate    time.Time
	Product      string
	Quantity     int
	Price        int
	details      string
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

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, rolName, rolPassword, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, customer_name, order_date, product, quantity, price FROM items")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var items []manzana
	for rows.Next() {
		var item manzana
		err := rows.Scan(&item.Id, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	inicio := (page - 1) * itemsPerPage
	fin := inicio + itemsPerPage
	if fin > len(items) {
		fin = len(items)
	}
	elementosPagina := items[inicio:fin]

	err = json.NewEncoder(w).Encode(elementosPagina)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func getItem(w http.ResponseWriter, r *http.Request) {
	variable := mux.Vars(r)
	id := variable["id"]

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, rolName, rolPassword, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT id, customer_name, order_date, product, quantity, price FROM items WHERE id =%s", id)
	var item manzana
	err = db.QueryRow(query).Scan(&item.Id, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price)
	if err != nil {
		log.Fatal(err)
	}

	encontrado := false
	if item != (manzana{}) {
		encontrado = true
	}

	if !encontrado {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("En nuestra base de datos no se encuentra el id " + id))
		return
	}
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

}

func getItemDetails(w http.ResponseWriter, r *http.Request) {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, rolName, rolPassword, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT COUNT(*) FROM items")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var count int
	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
	}

	rows, err = db.Query("SELECT id FROM items")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var items []int
	for rows.Next() {
		var item int
		err := rows.Scan(&item)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	detailsChannel := make(chan manzana, count)

	var wg sync.WaitGroup
	for _, item := range items {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			getDetails(id, detailsChannel)
		}(item)
	}
	wg.Wait()

	var detailedItems []manzana
	for i := 0; i < count; i++ {
		itemDetail := <-detailsChannel
		detailedItems = append(detailedItems, itemDetail)
	}

	err = json.NewEncoder(w).Encode(detailedItems)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

}

func getDetails(id int, c chan manzana) {
	time.Sleep(100 * time.Millisecond)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, rolName, rolPassword, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT id, details FROM items WHERE id =%d", id)
	var item manzana
	err = db.QueryRow(query).Scan(&item.Id, &item.details)
	if err != nil {
		log.Fatal(err)
	}

	c <- item
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

	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, rolName, rolPassword, dbName)
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("Connected to the database")

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
