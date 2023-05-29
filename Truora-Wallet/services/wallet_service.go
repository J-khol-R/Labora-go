package service

import (
	"database/sql"

	"github.com/J-khol-R/Labora-go/Truora-Wallet/models"
	repo "github.com/J-khol-R/Labora-go/Truora-Wallet/repositories"
)

type WalletService struct {
	Repository repo.Wallet
}

func (s *WalletService) Create(wallet models.Wallet, tx *sql.Tx) error {
	return s.Repository.Create(wallet, tx)
}

func (s *WalletService) Delete(id string, tx *sql.Tx) error {
	return s.Repository.Delete(id, tx)
}

func (s *WalletService) Update(wallet models.Wallet) error {
	return s.Repository.Update(wallet)
}

func (s *WalletService) GetStatus(id int) (models.Wallet, error) {
	return s.Repository.GetStatus(id)
}
