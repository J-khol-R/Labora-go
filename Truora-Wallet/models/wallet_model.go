package models

import "time"

type Wallet struct {
	Dni            int
	Country_id     string
	Fecha_creacion time.Time
	Id             int
}
