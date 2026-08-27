package database

import (
	"IManager-Src/utils"
	"database/sql"

	_ "modernc.org/sqlite"
)

func Open() (*sql.DB, error) {
	path, exists := utils.DatabasePath()
	if !exists {
		utils.InitalizeDatabase()
		path, exists = utils.DatabasePath()
	}
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
