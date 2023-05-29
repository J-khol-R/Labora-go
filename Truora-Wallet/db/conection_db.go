package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Conn *sql.DB

type EnvConfig struct {
	Host        string
	Port        string
	DbName      string
	RolName     string
	RolPassword string
}

func GetEnvConfig() EnvConfig {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	envConfig := EnvConfig{
		Host:        os.Getenv("HOST"),
		Port:        os.Getenv("PORT"),
		DbName:      os.Getenv("DB_NAME"),
		RolName:     os.Getenv("ROL_NAME"),
		RolPassword: os.Getenv("ROL_PASSWORD"),
	}

	return envConfig
}

func OpenDbConnection() *sql.DB {
	envConfig := GetEnvConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		envConfig.Host, envConfig.Port, envConfig.RolName, envConfig.RolPassword, envConfig.DbName)

	dbConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return dbConn
}

func init() {
	Conn = OpenDbConnection()
}
