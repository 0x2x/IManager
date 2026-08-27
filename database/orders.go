package database

import (
	"IManager-Src/utils"
	"database/sql"
	"fmt"
)

type OrdersData struct {
	BuyerID     int64
	FirstName   string
	LastName    string
	PhoneNumber string
	Telegram    string
}

func CreateOrdersTable(db *sql.DB) error {
	const query = `
		CREATE TABLE IF NOT EXISTS buyers (
			buyer_id INTEGER PRIMARY KEY AUTOINCREMENT,
			first_name TEXT NOT NULL DEFAULT '',
			last_name TEXT NOT NULL DEFAULT '',
			phone_number TEXT NOT NULL DEFAULT '',
			telegram TET NOT NULL DEFAULT ''
		);`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("create buyers table: %W", err)
	}
	return nil
}

func AddOrder(db *sql.DB, buyer Buyer) (string, error) {
	const query = `
		INSERT INTO buyers (buyer_id, first_name, last_name, phone_number, telegram)
		VALUES (?, ?, ?, ?, ?)
	`

	gid := utils.GenerateID()

	_, err := db.Exec(
		query,
		gid,
		buyer.FirstName,
		buyer.LastName,
		buyer.PhoneNumber,
		buyer.Telegram,
	)
	if err != nil {
		return gid, fmt.Errorf("inserting buyer into buyers table: %w", err)
	}

	return gid, nil
}
