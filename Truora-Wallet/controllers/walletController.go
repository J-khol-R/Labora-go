package controllers

import (
	"net/http"

	repo "github.com/J-khol-R/Labora-go/Truora-Wallet/repositories"
	service "github.com/J-khol-R/Labora-go/Truora-Wallet/services"
)

var walletService service.WalletService
var logService service.LogService

func init() {
	walletService = service.WalletService{
		Repository: &repo.PostgresWallet{},
	}

	logService = service.LogService{
		Repository: &repo.PostgresLog{},
	}
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	// walletService.GetStatus()
	// logService.CreateLog()
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
