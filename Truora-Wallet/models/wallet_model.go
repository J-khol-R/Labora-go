package models

import "time"

type Wallet struct {
	Dni            int       // checkid
	Country_id     string    // pais
	Fecha_creacion time.Time // creacion
	Id_persona     int       // cedula
}
