package database

import (
	"IManager-Src/utils"
	"database/sql"
	"fmt"
)

type InventoryData struct {
	inventory_id    int64
	ProductName     string
	ProductURL      string
	ProductPrice    float64
	ProductShipping float64
	ProductTax      float64
	quantity        int
}

func CreateInventoryTable(db *sql.DB) error {
	const query = `
		CREATE TABLE IF NOT EXISTS inventory (
			inventory_id INTEGER PRIMARY KEY AUTOINCREMENT,
			product_name TEXT NOT NULL DEFAULT ''
			product_url TEXT NOT NULL DEFAULT ''
			product_price FLOAT  NOT NULL DEFAULT 0.00
			product_shipping FLOAT NOT NULL DEFAULT 0.00
			product_tax FLOAT NOT NULL DEFAULT 0.00
			product_quantity INTEGER NOT NULL DEFAULT 0
		);`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("create buyers table: %W", err)
	}
	return nil
}

func AddItem(db *sql.DB, item InventoryData) error {
	const query = `
		INSERT INTO inventory (product_name, product_url, product_price, product_shipping, product_tax, product_quantity)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := db.Exec(
		query,
		item.ProductName,
		item.ProductURL,
		item.ProductPrice,
		item.ProductShipping,
		item.ProductShipping,
		item.ProductTax,
		item.quantity,
	)
	if err != nil {
		utils.PublicErrors("Issue occured while inserting into buyers table: %w\n\thttps://github.com/0x2x/IManager", err)
		return fmt.Errorf("inserting buyer into buyers table: %w", err)
	}
	return nil
}

func FindItem(db *sql.DB)
