package service

import "github.com/J-khol-R/Labora-go/Truora-Wallet/models"

type WalletService struct {
	DbHandler DBhandler
}

func (s *WalletService) CreateWallet(wallet models.Wallet) error {
	return s.DbHandler.CreateWallet(wallet)
}

func (s *WalletService) DeleteWallet(id int) error {
	return s.DbHandler.DeleteWallet(id)
}

func (s *WalletService) UpdateWallet(wallet models.Wallet) error {
	return s.DbHandler.UpdateWallet(wallet)
}

func (s *WalletService) GetStatus(id int) (models.Wallet, error) {
	return s.DbHandler.GetStatus(id)
}
