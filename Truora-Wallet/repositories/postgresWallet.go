package repositories

import (
	"github.com/J-khol-R/Labora-go/Truora-Wallet/db"
	"github.com/J-khol-R/Labora-go/Truora-Wallet/models"
)

type PostgresWallet struct { //objeto para acceder a la base de datos
}

func (p *PostgresWallet) Create(wallet models.Wallet) error {

	query := `INSERT INTO wallet (id_persona, dni, pais_id, creacion)
	VALUES ($1, $2, $3, $4)`
	_, err := db.Conn.Exec(query, wallet.Id_persona, wallet.Dni, wallet.Country_id, wallet.Fecha_creacion)

	return err
}

func (p *PostgresWallet) Update(wallet models.Wallet) error {
	query := `UPDATE wallet 
	SET id_persona = $1, dni = $2, pais_id = $3, creacion = $4,
	WHERE id = $5`
	_, err := db.Conn.Exec(query, wallet.Id_persona, wallet.Dni, wallet.Country_id, wallet.Fecha_creacion, wallet.Id_persona)

	return err
}

func (p *PostgresWallet) Delete(id int) error {

	query := `DELETE FROM wallet WHERE id = $1`
	_, err := db.Conn.Exec(query, id)

	return err
}

func (p *PostgresWallet) GetStatus(id int) (models.Wallet, error) {

	var wallet models.Wallet
	query := "SELECT * FROM wallet WHERE id_persona = $1"
	rows, err := db.Conn.Query(query, id)
	if err != nil {
		return wallet, err
	}
	defer rows.Close()

	if !rows.Next() {
		return wallet, err
	}

	err = rows.Scan(&wallet.Id_persona, &wallet.Dni, &wallet.Country_id, &wallet.Fecha_creacion)

	return wallet, err
}
