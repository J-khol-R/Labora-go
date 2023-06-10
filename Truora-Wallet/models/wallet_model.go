package models

import "time"

type Wallet struct {
	Id_persona     string    `json:"dni"`           // cedula
	Dni            string    `json:"checkId"`       // checkid
	Country_id     string    `json:"country"`       // pais
	Fecha_creacion time.Time `json:"fechaCreacion"` // creacion
	Balance        float64   `json:"balance"`
}

type WalletDetails struct {
	Id_persona         string               `json:"id"`
	Balance            float64              `json:"amount"`
	WalletTransactions []TransactionDetails `json:"movements"`
}

func (w WalletDetails) IsWalletEmpty() bool {
	if w.Id_persona != "" || w.Balance != 0 {
		return false
	}

	for _, transaccion := range w.WalletTransactions {
		if (transaccion.Movement != "" || transaccion.Amount != 0 || transaccion.Time != time.Time{}) {
			return false
		}
	}

	return true
}
