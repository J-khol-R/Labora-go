package repositories

import "github.com/J-khol-R/Labora-go/Truora-Wallet/models"

//definicion de las operaciones crud para la base de datos

type Wallet interface {
	Create(wallet models.Wallet) error
	Update(wallet models.Wallet) error
	Delete(id int) error
	GetStatus(id int) (models.Wallet, error)
}

type Log interface {
	Create(log models.Log) error
	Update(log models.Log) error
	Delete(id int) error
	Get(id int) (models.Log, error)
}
