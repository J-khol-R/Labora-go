package models

import "time"

type Transaction struct {
	Nro_transaction int       `json:"nroTransaction"`
	SenderId        string    `json:"sender_id"`
	ReceiverId      string    `json:"receiver_id"`
	Amount          float64   `json:"amount"`
	Time            time.Time `json:"time"`
}

type TransactionDetails struct {
	Movement string    `json:"movement"`
	Amount   float64   `json:"amount"`
	Time     time.Time `json:"time"`
}

func (t *TransactionDetails) MovementType(role string) {
	if role == "sender_id" {
		t.Movement = "retiro"
	} else {
		t.Movement = "deposito"
	}
}
