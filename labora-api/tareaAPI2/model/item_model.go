package model

import "time"

type Items struct {
	Id           int
	CustomerName string
	OrderDate    time.Time
	Product      string
	Quantity     int
	Price        int
	TotalPrice   int
	Vistas       int
}

type Itemsdetails struct {
	Items
	Details string
}

func (i Items) CalcularPrecio() int {
	return i.Quantity * i.Price
}
