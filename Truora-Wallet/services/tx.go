package service

import (
	"github.com/J-khol-R/Labora-go/Truora-Wallet/db"
	"github.com/J-khol-R/Labora-go/Truora-Wallet/models"
	repo "github.com/J-khol-R/Labora-go/Truora-Wallet/repositories"
)

type TxService struct {
	RepoLog    repo.Log
	RepoWallet repo.Wallet
}

var walletService WalletService
var logService LogService

func init() {
	walletService = WalletService{
		Repository: &repo.PostgresWallet{},
	}

	logService = LogService{
		Repository: &repo.PostgresLog{},
	}
}

func (s *TxService) CrearLogAndWalletTx(log models.Log, wallet models.Wallet, verificar bool) error {

	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}
	err = logService.CreateLog(log, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	if verificar {
		err = walletService.Create(wallet, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *TxService) DeleteLogAndWalletTx(dni string) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}
	err = walletService.Delete(dni, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = logService.DeleteLog(dni, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
