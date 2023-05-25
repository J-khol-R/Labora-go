package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/J-khol-R/Labora-go/Truora-Wallet/controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/WalletStatus/{id}", controllers.GetStatus).Methods("GET")
	r.HandleFunc("/CreateWallet", controllers.CreateWallet).Methods("POST")
	r.HandleFunc("/UpdateWallet/{id}", controllers.UpdateWallet).Methods("PUT")
	r.HandleFunc("/DeleteWallet/{id}", controllers.DeleteWallet).Methods("DELETE")

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"PUT", "GET", "POST", "DELETE"},
	})

	handler := corsOptions.Handler(r)

	direccion := ":8081"

	servidor := &http.Server{
		Handler:      handler,
		Addr:         direccion,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Escuchando en %s. Presiona CTRL + C para salir", servidor.Addr)
	log.Fatal(servidor.ListenAndServe())

}
