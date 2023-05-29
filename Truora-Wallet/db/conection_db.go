package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func OpenDbConection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("Error cargando el archivo .env: %v", err)
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	dbName := os.Getenv("DB_NAME")
	rolName := os.Getenv("ROL_NAME")
	rolPassword := os.Getenv("ROL_PASSWORD")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, rolName, rolPassword, dbName)
	dbConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Errorf("Error conectando a la base de datos: %v", err)
	}
	return dbConn
}

var Conn *sql.DB

func init() {
	Conn = OpenDbConection()
}
