package service

import (
	"database/sql"
	"fmt"

	conection "github.com/J-khol-R/Labora-go/Truora-Wallet/db"
	"github.com/J-khol-R/Labora-go/Truora-Wallet/models"
)

type PostgresDBHandler struct {
	db *sql.DB
}

func (p *PostgresDBHandler) CreateWallet(wallet models.Wallet) error {
	p.db = conection.DbConnection()
	defer p.db.Close()

	query := `INSERT INTO wallet (id_persona, dni, pais_id, creacion)
	VALUES ($1, $2, $3, $4)`
	_, err := p.db.Exec(query, wallet.Id_persona, wallet.Dni, wallet.Country_id, wallet.Fecha_creacion)

	return err
}

func (p *PostgresDBHandler) UpdateWallet(wallet models.Wallet) error {
	p.db = conection.DbConnection()
	defer p.db.Close()

	query := `UPDATE wallet 
	SET id_persona = $1, dni = $2, pais_id = $3, creacion = $4,
	WHERE id = $5`
	_, err := p.db.Exec(query, wallet.Id_persona, wallet.Dni, wallet.Country_id, wallet.Fecha_creacion, wallet.Id_persona)

	return err
}

func (p *PostgresDBHandler) DeleteWallet(id int) error {
	p.db = conection.DbConnection()
	defer p.db.Close()

	query := `DELETE FROM wallet WHERE id = $1`
	_, err := p.db.Exec(query, id)

	return err
}

func (p *PostgresDBHandler) WalletStatus(id int) (models.Wallet, error) {
	p.db = conection.DbConnection()
	defer p.db.Close()

	query := fmt.Sprintf("SELECT * FROM wallet WHERE id =%d", id)
	var wallet models.Wallet
	err := p.db.QueryRow(query).Scan(&wallet.Id_persona, &wallet.Dni, &wallet.Country_id, &wallet.Fecha_creacion)

	return wallet, err
}

func (p *PostgresDBHandler) CreateLog(log models.Log) error {
	p.db = conection.DbConnection()
	defer p.db.Close()

	query := `INSERT INTO solicitud (id_persona, dni_solicitud, fecha_solicitud, pais, estado, codigo)
	VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := p.db.Exec(query, log.Id_persona, log.Dni_solicitud, log.Fecha_solicitud, log.Pais, log.Estado, log.Codigo)

	return err
}

func (p *PostgresDBHandler) UpdateLog(log models.Log) error {
	p.db = conection.DbConnection()
	defer p.db.Close()

	query := `UPDATE wallsolicitudet 
	SET id_persona = $1, dni_solicitud = $2, fecha_solicitud = $3, pais = $4, estado = $5, codigo = $6
	WHERE id = $7`
	_, err := p.db.Exec(query, log.Id_persona, log.Dni_solicitud, log.Fecha_solicitud, log.Pais, log.Estado, log.Codigo, log.Id_persona)

	return err
}

func (p *PostgresDBHandler) DeleteLog(id int) error {
	p.db = conection.DbConnection()
	defer p.db.Close()

	query := `DELETE FROM solicitud WHERE id = $1`
	_, err := p.db.Exec(query, id)

	return err
}

func (p *PostgresDBHandler) GetLog(id int) (models.Log, error) {
	p.db = conection.DbConnection()
	defer p.db.Close()

	query := fmt.Sprintf("SELECT * FROM solicitud WHERE id =%d", id)
	var log models.Log
	err := p.db.QueryRow(query).Scan(&log.Id_persona, &log.Dni_solicitud, &log.Fecha_solicitud, &log.Pais, &log.Estado, &log.Codigo)

	return log, err
}
