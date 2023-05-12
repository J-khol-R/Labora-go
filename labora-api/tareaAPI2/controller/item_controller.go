package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/J-khol-R/Labora-go/labora-api/tareaAPI2/model"
	"github.com/J-khol-R/Labora-go/labora-api/tareaAPI2/service"
	"github.com/gorilla/mux"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
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

	db := service.DbConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, customer_name, order_date, product, quantity, price FROM items")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var items []model.Items
	for rows.Next() {
		var item model.Items
		err := rows.Scan(&item.Id, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price)
		if err != nil {
			log.Fatal(err)
		}
		item.TotalPrice = item.CalcularPrecio()
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

func GetItem(w http.ResponseWriter, r *http.Request) {
	variable := mux.Vars(r)
	id := variable["id"]

	db := service.DbConnection()
	defer db.Close()

	query := fmt.Sprintf("SELECT id, customer_name, order_date, product, quantity, price FROM items WHERE id =%s", id)
	var item model.Items
	err := db.QueryRow(query).Scan(&item.Id, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price)
	if err != nil {
		log.Fatal(err)
	}

	item.TotalPrice = item.CalcularPrecio()

	encontrado := false
	if item != (model.Items{}) {
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

func GetItemDetails(w http.ResponseWriter, r *http.Request) {

	db := service.DbConnection()
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

	detailsChannel := make(chan model.Itemsdetails, count)

	var wg sync.WaitGroup
	for _, item := range items {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			GetDetails(id, detailsChannel)
		}(item)

	}
	wg.Wait()

	var detailedItems []model.Itemsdetails
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

func GetDetails(id int, c chan model.Itemsdetails) {
	time.Sleep(100 * time.Millisecond)

	db := service.DbConnection()
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM items WHERE id =%d", id)
	var item model.Itemsdetails
	err := db.QueryRow(query).Scan(&item.Id, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price, &item.Details)
	if err != nil {
		log.Fatal(err)
	}
	item.TotalPrice = item.CalcularPrecio()

	c <- item
}

func GetItemByName(w http.ResponseWriter, r *http.Request) {
	variable := mux.Vars(r)
	name := variable["name"]

	db := service.DbConnection()
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("SELECT id, customer_name, order_date, product, quantity, price FROM items WHERE customer_name = '%s'", name))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var items []model.Items
	for rows.Next() {
		var item model.Items
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

	var names []model.Items
	var encontrado bool
	for _, valor := range items {
		if strings.EqualFold(valor.CustomerName, name) {
			names = append(names, valor)
			encontrado = true
			break
		}
	}

	if !encontrado {
		w.WriteHeader(http.StatusNotFound) //devuelve el error 404 que es cuando no encuantra algo
		return
	}
	err = json.NewEncoder(w).Encode(names)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item model.Items
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Error al decodificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	db := service.DbConnection()
	defer db.Close()

	query := `INSERT INTO items (customer_name, order_date, product, quantity, price)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = db.QueryRow(query, item.CustomerName, item.OrderDate, item.Product, item.Quantity, item.Price).Scan(&item.Id)
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	variable := mux.Vars(r)
	id := variable["id"]

	var datos model.Items
	err := json.NewDecoder(r.Body).Decode(&datos)
	if err != nil {
		http.Error(w, "Error al decodificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	db := service.DbConnection()
	defer db.Close()

	var existe bool
	query := `SELECT COUNT(*) > 0 AS exists FROM items WHERE id = $1`
	err = db.QueryRow(query, id).Scan(&existe)
	if err != nil {
		log.Fatal(err)
	}

	if !existe {
		w.Write([]byte("El item que esta intentando actualizar no existe :("))
		return
	}

	query = `UPDATE items 
	SET customer_name = $1, order_date = $2, product = $3, quantity = $4, price = $5
	WHERE id = $6`
	_, err = db.Exec(query, datos.CustomerName, datos.OrderDate, datos.Product, datos.Quantity, datos.Price, id)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("El item se actualizo correctamente :)"))

}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	variable := mux.Vars(r)
	id := variable["id"]

	db := service.DbConnection()
	defer db.Close()

	var existe bool
	query := `SELECT COUNT(*) > 0 AS exists FROM items WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&existe)
	if err != nil {
		log.Fatal(err)
	}

	if !existe {
		w.Write([]byte("El item que esta intentando borrar no existe :("))
		return
	}

	query = `DELETE FROM items WHERE id = $1`
	_, err = db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("el item se elimino correctamente :)"))

}
