package models

import "time"

type Wallet struct {
	Dni            string    `json:"checkId"`       // checkid
	Country_id     string    `json:"country"`       // pais
	Fecha_creacion time.Time `json:"fechaCreacion"` // creacion
	Id_persona     string    `json:"dni"`           // cedula
}
