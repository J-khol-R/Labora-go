package model

import "time"

type Items struct {
	Id           int
	CustomerName string
	OrderDate    time.Time
	Product      string
	Quantity     int
	Price        int
}

type Itemsdetails struct {
	Items
	Details string
}
