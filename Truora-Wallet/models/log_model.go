package models

import "time"

type Log struct {
	Dni_solicitud   string    //check id
	Fecha_solicitud time.Time // fecha solicitud
	Pais            string    //pais harcodeado
	Estado          string    // completado o rechazado
	Codigo          int       // 1, 0 o -1
	Id_persona      int       // cedula persona
}
