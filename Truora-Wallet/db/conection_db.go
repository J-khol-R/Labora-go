package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

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
