package database

import (
	"IManager-Src/utils"
	"database/sql"
	"fmt"
)

type Buyer struct {
	BuyerID     int64
	FirstName   string
	LastName    string
	PhoneNumber string
	Telegram    string
	ItemID      int64
}

func CreateBuyersTable(db *sql.DB) error {
	const query = `
		CREATE TABLE IF NOT EXISTS buyers (
			buyer_id INTEGER PRIMARY KEY AUTOINCREMENT,
			first_name TEXT NOT NULL DEFAULT '',
			last_name TEXT NOT NULL DEFAULT '',
			phone_number TEXT NOT NULL DEFAULT '',
			telegram TEXT NOT NULL DEFAULT ''
			item_id INTEGER,
			FOREIGN KEY (item_id) REFERENCES inventory(inventory_id) 
		);`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("create buyers table: %W", err)
	}
	return nil
}

func AddBuyer(db *sql.DB, buyer Buyer) (string, error) {
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
