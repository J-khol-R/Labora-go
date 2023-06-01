package repositories

import (
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
		return false, err
	}

	if count != 2 {
		return false, err
	}

	var balance float64
	query = "SELECT balance FROM wallet WHERE id_persona = $1"
	err = tx.QueryRow(query, transaction.SenderId).Scan(&balance)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	if balance < transaction.Amount {
		return false, err
	}

	// query = "update wallet set balance = balance - $1 where id_persona = $2"
	// _, err = tx.Exec(query, transaction.Amount, transaction.SenderId)
	// if err != nil {
	// 	tx.Rollback()
	// 	// log.Fatal(err)
	// 	fmt.Print("update sender")
	// 	return false, err
	// }

	// query = "update wallet set balance = balance + $1 where id_persona = $2"
	// _, err = tx.Exec(query, transaction.Amount, transaction.ReceiverId)
	// if err != nil {
	// 	tx.Rollback()
	// 	// log.Fatal(err)
	// 	fmt.Print("update receiber")
	// 	return false, err
	// }

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return false, err
	}

	return true, nil
}

func (p *PostgresTransaction) SaveTransaction(transaction models.Transaction) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}

	query := "update wallet set balance = balance - $1 where id_persona = $2"
	_, err = tx.Exec(query, transaction.Amount, transaction.SenderId)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = "update wallet set balance = balance + $1 where id_persona = $2"
	_, err = tx.Exec(query, transaction.Amount, transaction.ReceiverId)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = `INSERT INTO transactions (sender_id, receiver_id, amount, time_transaction)
	VALUES ($1, $2, $3, NOW())`
	_, err = tx.Exec(query, transaction.SenderId, transaction.ReceiverId, transaction.Amount)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (p *PostgresTransaction) GetWalletTransactions(id string) (models.WalletDetails, error) {
	var wallet models.WalletDetails

	tx, err := db.Conn.Begin()
	if err != nil {
		return models.WalletDetails{}, err
	}

	query := "select id_persona, amount from wallet where id_persona = $1"
	err = tx.QueryRow(query, id).Scan(&wallet.Id_persona, &wallet.Balance)
	if err != nil {
		tx.Rollback()
		return models.WalletDetails{}, err
	}

	query = `select * from 
	(SELECT 'receiver_id' AS role_string , receiver_id as wallet_id, time_transaction AS hora, amount FROM transactions
	UNION ALL
	SELECT 'sender_id' AS role_string, sender_id as wallet_id, time_transaction AS hora, amount FROM transactions
	order by hora) as trans where wallet_id = $1;`

	rows, err := tx.Query(query, id)
	if err != nil {
		return models.WalletDetails{}, err
	}
	defer rows.Close()

	var transactions []models.TransactionDetails
	var role string
	for rows.Next() {
		var transaction models.TransactionDetails
		var id int
		err := rows.Scan(&role, &id, &transaction.Amount, &transaction.Time)
		if err != nil {
			return models.WalletDetails{}, err
		}
		transaction.MovementType(role)
	}
	err = rows.Err()
	if err != nil {
		return models.WalletDetails{}, err
	}

	wallet.WalletTransactions = transactions

	return wallet, nil
}
