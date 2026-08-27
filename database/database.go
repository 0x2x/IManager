package database

import (
	"IManager-Src/utils"
	"database/sql"

	_ "modernc.org/sqlite"
)

func Open() (*sql.DB, error) {
	path, result := utils.DatabasePath()
	// Fetch imanager.db from temp
	if result == false {
		utils.InitalizeDatabase()           // Should pass
		path, result = utils.DatabasePath() // Update with path if need to.
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
