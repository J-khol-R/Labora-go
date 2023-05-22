package service

import (
	"database/sql"
	"fmt"

	"github.com/J-khol-R/Labora-go/labora-api/tareaAPI2/model"
)

type PostgresDBHandler struct {
	db *sql.DB
}

func (p *PostgresDBHandler) CreateItem(item model.Items) error {
	db := DbConnection()
	defer db.Close()

	query := `INSERT INTO items (customer_name, order_date, product, quantity, price)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(query, item.CustomerName, item.OrderDate, item.Product, item.Quantity, item.Price).Scan(&item.Id)

	return err
}

func (p *PostgresDBHandler) GetItem(id int) (model.Items, error) {
	p.db = DbConnection()
	defer p.db.Close()

	query := fmt.Sprintf("SELECT id, customer_name, order_date, product, quantity, price, vistas FROM items WHERE id =%d", id)
	var item model.Items
	err := p.db.QueryRow(query).Scan(&item.Id, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price, &item.Vistas)

	return item, err
}

func (p *PostgresDBHandler) UpdateItem(item model.Items, id int) error {
	p.db = DbConnection()
	defer p.db.Close()

	query := `UPDATE items 
	SET customer_name = $1, order_date = $2, product = $3, quantity = $4, price = $5
	WHERE id = $6`
	_, err := p.db.Exec(query, item.CustomerName, item.OrderDate, item.Product, item.Quantity, item.Price, id)

	return err
}

func (p *PostgresDBHandler) DeleteItem(id int) error {
	p.db = DbConnection()
	defer p.db.Close()

	query := `DELETE FROM items WHERE id = $1`
	_, err := p.db.Exec(query, id)

	return err
}
