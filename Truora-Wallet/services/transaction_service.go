package service

import (
	"github.com/J-khol-R/Labora-go/Truora-Wallet/models"
	repo "github.com/J-khol-R/Labora-go/Truora-Wallet/repositories"
)

type TransactionService struct { //objeto para poder utilizar las operaciones CRUD de solicitud
	Repository repo.Transaction
}

func (s *TransactionService) Create(transaction models.Transaction) (bool, error) {
	return s.Repository.Create(transaction)
}

func (s *TransactionService) SaveTransaction(transaction models.Transaction) error {
	return s.Repository.SaveTransaction(transaction)
}
