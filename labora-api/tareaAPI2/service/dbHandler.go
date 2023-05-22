package service

import "github.com/J-khol-R/Labora-go/labora-api/tareaAPI2/model"

type DBHandler interface {
	CreateItem(item model.Items) error
	GetItem(id int) (model.Items, error)
	UpdateItem(item model.Items, id int) error
	DeleteItem(id int) error
}
