package models

import "time"

type Transaction struct {
	Nro_transaction int       `json:"nroTransaction"`
	SenderId        string    `json:"sender_id"`
	ReceiverId      string    `json:"receiver_id"`
	Amount          float64   `json:"amount"`
	Movement        string    `json:"movement"`
	Time            time.Time `json:"time"`
}
