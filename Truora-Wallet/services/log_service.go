package service

import (
	"database/sql"

	"github.com/J-khol-R/Labora-go/Truora-Wallet/models"
	repo "github.com/J-khol-R/Labora-go/Truora-Wallet/repositories"
)

type LogService struct { //objeto para poder utilizar las operaciones CRUD de solicitud
	Repository repo.Log
}

func (s *LogService) CreateLog(log models.Log, tx *sql.Tx) error {
	return s.Repository.Create(log, tx)
}

func (s *LogService) DeleteLog(id string, tx *sql.Tx) error {
	return s.Repository.Delete(id, tx)
}

func (s *LogService) UpdateLog(log models.Log) error {
	return s.Repository.Update(log)
}

func (s *LogService) GetLog(id int) (models.Log, error) {
	return s.Repository.Get(id)
}
