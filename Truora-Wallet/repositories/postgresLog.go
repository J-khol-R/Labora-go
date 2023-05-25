package repositories

import (
	"fmt"

	"github.com/J-khol-R/Labora-go/Truora-Wallet/db"
	"github.com/J-khol-R/Labora-go/Truora-Wallet/models"
)

type PostgresLog struct {
}

func (p *PostgresLog) Create(log models.Log) error {

	query := `INSERT INTO solicitud (id_persona, dni_solicitud, fecha_solicitud, pais, estado, codigo)
	VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Conn.Exec(query, log.Id_persona, log.Dni_solicitud, log.Fecha_solicitud, log.Pais, log.Estado, log.Codigo)

	return err
}

func (p *PostgresLog) Update(log models.Log) error {

	query := `UPDATE wallsolicitudet 
	SET id_persona = $1, dni_solicitud = $2, fecha_solicitud = $3, pais = $4, estado = $5, codigo = $6
	WHERE id = $7`
	_, err := db.Conn.Exec(query, log.Id_persona, log.Dni_solicitud, log.Fecha_solicitud, log.Pais, log.Estado, log.Codigo, log.Id_persona)

	return err
}

func (p *PostgresLog) Delete(id int) error {

	query := `DELETE FROM solicitud WHERE id = $1`
	_, err := db.Conn.Exec(query, id)

	return err
}

func (p *PostgresLog) Get(id int) (models.Log, error) {

	query := fmt.Sprintf("SELECT * FROM solicitud WHERE id =%d", id)
	var log models.Log
	err := db.Conn.QueryRow(query).Scan(&log.Id_persona, &log.Dni_solicitud, &log.Fecha_solicitud, &log.Pais, &log.Estado, &log.Codigo)

	return log, err
}
