package models

import (
	"time"
)

type Log struct {
	Id_persona      string    // cedula persona
	Dni_solicitud   string    // check id
	Fecha_solicitud time.Time // fecha solicitud
	Pais            string    // pais harcodeado
	Codigo          float64   // -1 a 1
	Estado          string    // completado o rechazado
}

func (l Log) VerificarEstado() string {
	if l.Codigo < 1 {
		return "rechazado"
	} else {
		return "completado"
	}
}
