package models

import "time"

type Wallet struct {
	Dni            string    // checkid
	Country_id     string    // pais
	Fecha_creacion time.Time // creacion
	Id_persona     int       // cedula
}
