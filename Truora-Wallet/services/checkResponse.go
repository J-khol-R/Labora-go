package service

import "github.com/J-khol-R/Labora-go/Truora-Wallet/models"

func CrearLog(dni, country string) (models.Log, error) {
	respuesta, err := SendRequestToTruora(dni, country)
	if err != nil {
		return models.Log{}, err
	}

	var log models.Log
	log.Id_persona = respuesta.Check.NationalID
	log.Dni_solicitud = respuesta.Check.CheckID
	log.Fecha_solicitud = respuesta.Check.CreationDate
	log.Pais = respuesta.Check.Country
	log.Codigo = respuesta.Check.Score
	log.Estado = log.VerificarEstado()

	return log, nil
}

func InstanciarStructs(dni, country string) (models.Log, models.Wallet, bool, error) {
	var wallet models.Wallet

	log, err := CrearLog(dni, country)
	if err != nil {
		return models.Log{}, wallet, false, err
	}

	if log.Codigo == 1 {
		wallet.Id_persona = log.Id_persona
		wallet.Dni = log.Dni_solicitud
		wallet.Fecha_creacion = log.Fecha_solicitud
		wallet.Country_id = log.Pais

		return log, wallet, true, nil
	}

	return log, wallet, false, nil
}
