package repositories

import (
	"fmt"

	"github.com/J-khol-R/Labora-go/Truora-Wallet/db"
	"github.com/J-khol-R/Labora-go/Truora-Wallet/models"
)

type PostgresTransaction struct {
}

func (p *PostgresTransaction) Create(transaction models.Transaction) (bool, error) {
	tx, err := db.Conn.Begin()
	if err != nil {
		return false, err
	}

	var count int
	query := "SELECT COUNT(*) FROM wallet WHERE id_persona IN ($1, $2)"
	err = tx.QueryRow(query, transaction.SenderId, transaction.ReceiverId).Scan(&count)
	if err != nil {
		tx.Rollback()
		fmt.Print("selct count")
		// log.Fatal(err)
		return false, err
	}

	// query := "SELECT COUNT(*) FROM wallet WHERE id_persona IN ($1, $2)"
	// rows, err := tx.Query(query, transaction.SenderId, transaction.ReceiverId)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// var count int
	// if rows.Next() {
	// 	err = rows.Scan(&count)
	// 	if err != nil {
	// 		return false, err
	// 	}
	// }

	if count != 2 {
		return false, err //fmt.Errorf("Uno, o ambos id no tienen una wallet")
	}

	var balance float64
	query = "SELECT balance FROM wallet WHERE id_persona = $1"
	err = tx.QueryRow(query, transaction.SenderId).Scan(&balance)
	if err != nil {
		tx.Rollback()
		fmt.Print("selct balance")
		// log.Fatal(err)
		return false, err
	}

	if balance < transaction.Amount {
		return false, err //fmt.Errorf("Fondos insuficientes")
	}

	query = "update wallet set balance = balance - $1 where id_persona = $2"
	_, err = tx.Exec(query, transaction.Amount, transaction.SenderId)
	if err != nil {
		tx.Rollback()
		// log.Fatal(err)
		fmt.Print("update sender")
		return false, err
	}

	query = "update wallet set balance = balance + $1 where id_persona = $2"
	_, err = tx.Exec(query, transaction.Amount, transaction.ReceiverId)
	if err != nil {
		tx.Rollback()
		// log.Fatal(err)
		fmt.Print("update receiber")
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return false, err
	}

	return true, nil
}

func (p *PostgresTransaction) SaveTransaction(Transaction models.Transaction) error {

	query := `INSERT INTO transactions (sender_id, receiver_id, amount, time_transaction)
	VALUES ($1, $2, $3, NOW())`
	_, err := db.Conn.Exec(query, Transaction.SenderId, Transaction.ReceiverId, Transaction.Amount)
	if err != nil {
		return err
	}
	return nil
}
