package service

import "github.com/J-khol-R/Labora-go/Truora-Wallet/models"

type DBhandler interface {
	CreateWallet(wallet models.Wallet) error
	UpdateWallet(wallet models.Wallet) error
	DeleteWallet(id int) error
	GetStatus(id int) (models.Wallet, error)

	CreateLog(lod models.Log) error
	UpdateLog(log models.Log) error
	DeleteLog(id int) error
	GetLog(id int) (models.Log, error)
}
