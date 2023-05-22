package config

import (
	"net/http"
	"time"

	"github.com/J-khol-R/Labora-go/labora-api/tareaAPI2/controller"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func SetupServer() *http.Server {
	r := mux.NewRouter()

	r.HandleFunc("/items", controller.GetItems).Methods("GET")
	r.HandleFunc("/items/details", controller.GetItemDetails).Methods("GET")
	r.HandleFunc("/items/{id}", controller.GetItem).Methods("GET")
	r.HandleFunc("/items/nombre/{name}", controller.GetItemByName).Methods("GET")
	r.HandleFunc("/items", controller.CreateItem).Methods("POST")
	r.HandleFunc("/items/{id}", controller.UpdateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", controller.DeleteItem).Methods("DELETE")

	// corsOptions := handles.CORS(
	// 	handlers.AllowedOrigins([]string{"*"}),
	// 	handles.AllowedMethods([]string{"PUT"}),
	// )

	// handler := corsOptions(r)

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"PUT"},
	})

	handler := corsOptions.Handler(r)

	direccion := ":3000"

	servidor := &http.Server{
		Handler:      handler,
		Addr:         direccion,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return servidor
}
