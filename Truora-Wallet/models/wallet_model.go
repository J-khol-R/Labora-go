package models

import "time"

type Wallet struct {
	Id_persona     string    `json:"dni"`           // cedula
	Dni            string    `json:"checkId"`       // checkid
	Country_id     string    `json:"country"`       // pais
	Fecha_creacion time.Time `json:"fechaCreacion"` // creacion
	Balance        float64   `json:"balance"`
	// WalletTransactions []Transaction `json:"transactions"`
}
