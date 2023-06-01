package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/J-khol-R/Labora-go/Truora-Wallet/models"
	repo "github.com/J-khol-R/Labora-go/Truora-Wallet/repositories"
	service "github.com/J-khol-R/Labora-go/Truora-Wallet/services"
	"github.com/gorilla/mux"
)

var walletService service.WalletService

// var logService service.LogService
var txService service.TxService
var transactionService service.TransactionService

var mutex sync.Mutex

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

	transactionService = service.TransactionService{
		Repository: &repo.PostgresTransaction{},
	}
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	variable := mux.Vars(r)
	id := (variable["id"])

	num, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error al convertir el id: "+err.Error(), http.StatusInternalServerError)
		return
	}

	wallet, err := walletService.GetStatus(num)
	if err != nil {
		http.Error(w, "Error al obtener datos: "+err.Error(), http.StatusInternalServerError)
		return
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

func GetWallet(w http.ResponseWriter, r *http.Request) {
	variable := mux.Vars(r)
	id := (variable["id"])

	wallet, err := transactionService.GetWalletTransactions(id)
	if err != nil {
		http.Error(w, "Error al traer datos: "+err.Error(), http.StatusInternalServerError)
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

	log, wallet, verificar, err := service.InstanciarStructs(datos.Dni, datos.Country)
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
	var wallet models.Wallet
	err := json.NewDecoder(r.Body).Decode(&wallet)
	if err != nil {
		http.Error(w, "Error al ingresar los datos: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = walletService.Update(wallet)
	if err != nil {
		http.Error(w, "Error al actualizar en la base de datos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("wallet actualizado exitosamente"))
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

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaccion models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaccion)
	if err != nil {
		http.Error(w, "Error al decodificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	mutex.Lock()
	exit, err := transactionService.Create(transaccion)
	if err != nil {
		http.Error(w, "Error al procesar la transaccion: "+err.Error(), http.StatusInternalServerError)
		return
	}
	mutex.Unlock()

	if !exit {
		http.Error(w, "transaccion invalida: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = transactionService.SaveTransaction(transaccion)
	if err != nil {
		http.Error(w, "Error al almacenar en la base de datos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var mensaje models.Message
	mensaje.Ok()

	err = json.NewEncoder(w).Encode(mensaje)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

}
