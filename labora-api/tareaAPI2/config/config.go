package config

import (
	"net/http"
	"time"

	"github.com/J-khol-R/Labora-go/labora-api/tareaAPI2/controller"

	"github.com/gorilla/mux"
)

func SetupServer() *http.Server {
	enrutador := mux.NewRouter()

	enrutador.HandleFunc("/items", controller.GetItems).Methods("GET")
	enrutador.HandleFunc("/items/details", controller.GetItemDetails).Methods("GET")
	enrutador.HandleFunc("/items/{id}", controller.GetItem).Methods("GET")
	enrutador.HandleFunc("/items/nombre/{name}", controller.GetItemByName).Methods("GET")
	enrutador.HandleFunc("/items", controller.CreateItem).Methods("POST")
	enrutador.HandleFunc("/items/{id}", controller.UpdateItem).Methods("PUT")
	enrutador.HandleFunc("/items/{id}", controller.DeleteItem).Methods("DELETE")

	direccion := ":3000"

	servidor := &http.Server{
		Handler:      enrutador,
		Addr:         direccion,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return servidor
}
