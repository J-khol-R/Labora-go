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

// var logService service.LogService
var txService service.TxService

func init() {
	walletService = service.WalletService{
		Repository: &repo.PostgresWallet{},
	}
	// logService = service.LogService{
	// 	Repository: &repo.PostgresLog{},
	// }
	txService = service.TxService{
		RepoLog:    &repo.PostgresLog{},
		RepoWallet: &repo.PostgresWallet{},
	}
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
	var datos models.Datos
	err := json.NewDecoder(r.Body).Decode(&datos)
	if err != nil {
		http.Error(w, "Error al decodificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var wallet models.Wallet
	var log models.Log
	var verificar bool
	log, wallet, verificar, err = service.VerificarEstado(datos.Dni, datos.Country)
	if err != nil {
		http.Error(w, "Error al hacer la peticion: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = txService.CrearLogAndWalletTx(log, wallet, verificar)
	if err != nil {
		http.Error(w, "Error al almacenar en la base de datos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if verificar {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Wallet creada exitosamente, id: " + datos.Dni))
		return
	}

	w.Write([]byte("No podemos crear su Wallet, score menor a 1 "))

}

func UpdateWallet(w http.ResponseWriter, r *http.Request) {
	//implementar aqui
}

func DeleteWallet(w http.ResponseWriter, r *http.Request) {
	var datos models.Datos
	err := json.NewDecoder(r.Body).Decode(&datos)
	if err != nil {
		http.Error(w, "Error al decodificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = txService.DeleteLogAndWalletTx(datos.Dni)
	if err != nil {
		http.Error(w, "Error al borrar el wallet: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("wallet borrado exitosamente, id: " + datos.Dni))
}
