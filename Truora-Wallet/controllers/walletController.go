package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/J-khol-R/Labora-go/Truora-Wallet/models"
	repo "github.com/J-khol-R/Labora-go/Truora-Wallet/repositories"
	service "github.com/J-khol-R/Labora-go/Truora-Wallet/services"
	"github.com/gorilla/mux"
)

var walletService service.WalletService

//var logService service.LogService

func init() {
	walletService = service.WalletService{
		Repository: &repo.PostgresWallet{},
	}

	// logService = service.LogService{
	// 	Repository: &repo.PostgresLog{},
	// }
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	variable := mux.Vars(r)
	id := (variable["id"])

	num, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	wallet, err := walletService.GetStatus(num)
	if err != nil {
		log.Fatal(err)
	}

	encontrado := false
	if wallet != (models.Wallet{}) {
		encontrado = true
	}

	if !encontrado {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("En nuestra base de datos no se encuentra un wallet con el id: " + id))
		return
	}
	err = json.NewEncoder(w).Encode(wallet)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	//implementar aqui
}

func UpdateWallet(w http.ResponseWriter, r *http.Request) {
	//implementar aqui
}

func DeleteWallet(w http.ResponseWriter, r *http.Request) {
	//implementar aqui
}
