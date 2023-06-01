package repositories

import (
	"database/sql"

	"github.com/J-khol-R/Labora-go/Truora-Wallet/models"
)

//definicion de las operaciones crud para la base de datos

type Wallet interface {
	Create(wallet models.Wallet, tx *sql.Tx) error
	Update(wallet models.Wallet) error
	Delete(id string, tx *sql.Tx) error
	GetStatus(id int) (models.Wallet, error)
}

type Log interface {
	Create(log models.Log, tx *sql.Tx) error
	Update(log models.Log) error
	Delete(id string, tx *sql.Tx) error
	Get(id int) (models.Log, error)
}

type Transaction interface {
	Create(transaction models.Transaction) error
	AprovalTransaction(transaction models.Transaction) (bool, error)
	GetWalletTransactions(id string) (models.WalletDetails, error)
}
