package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/J-khol-R/Labora-go/labora-api/tareaAPI2/model"
	_ "github.com/lib/pq"
)

type ItemService struct {
	dbHandler DBHandler
}

const (
	host   = "localhost"
	port   = "5432"
	dbName = "labora_proyect_1"

	rolName     = "postgres"
	rolPassword = "0b3j1t4,"
)

func DbConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, rolName, rolPassword, dbName)
	dbConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	return dbConn
}

func (s *ItemService) CreateItem(item model.Items) error {
	return s.dbHandler.CreateItem(item)
}

func (s *ItemService) GetItem(id int) (model.Items, error) {
	return s.dbHandler.GetItem(id)
}

func (s *ItemService) UpdateItem(item model.Items, id int) error {
	return s.dbHandler.UpdateItem(item, id)
}

func (s *ItemService) DeleteItem(id int) error {
	return s.dbHandler.DeleteItem(id)
}
