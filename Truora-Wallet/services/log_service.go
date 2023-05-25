package service

import "github.com/J-khol-R/Labora-go/Truora-Wallet/models"

type LogService struct {
	DbHandler DBhandler
}

func (s *LogService) CreateLog(log models.Log) error {
	return s.DbHandler.CreateLog(log)
}

func (s *LogService) DeleteLog(id int) error {
	return s.DbHandler.DeleteLog(id)
}

func (s *LogService) UpdateLog(log models.Log) error {
	return s.DbHandler.UpdateLog(log)
}

func (s *LogService) GetLog(id int) (models.Log, error) {
	return s.DbHandler.GetLog(id)
}
