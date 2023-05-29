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

	r.HandleFunc("/walletStatus/{id}", controllers.GetStatus).Methods(http.MethodGet)
	r.HandleFunc("/createWallet", controllers.CreateWallet).Methods(http.MethodPost)
	r.HandleFunc("/updateWallet", controllers.UpdateWallet).Methods(http.MethodPut)
	r.HandleFunc("/deleteWallet", controllers.DeleteWallet).Methods(http.MethodDelete)

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
