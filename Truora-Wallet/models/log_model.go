package models

import "time"

type Log struct {
	Dni_solicitud   string
	Fecha_solicitud time.Time
	Estado          string
	Id              int
}
