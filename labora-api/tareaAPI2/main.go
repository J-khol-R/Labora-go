package main

import (
	"fmt"
	"log"

	"github.com/J-khol-R/Labora-go/labora-api/tareaAPI2/config"
)

func main() {
	servidor := config.SetupServer()

	fmt.Printf("Escuchando en %s. Presiona CTRL + C para salir", servidor.Addr)
	log.Fatal(servidor.ListenAndServe())
}
