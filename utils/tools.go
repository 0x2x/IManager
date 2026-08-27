package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func GenerateID() string {
	id := uuid.New()
	return id.String()
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	if err == nil {
		return info.IsDir()
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func fileExist(path string) bool { // I know its just running dirExists().
	return dirExists(path)
}

func InitalizeDatabase() bool { // May need to make OS specific?
	cacheDIR, err := os.UserCacheDir()
	path := filepath.Join(cacheDIR, "IManager")
	dbPath := filepath.Join(path, "imanager.db")
	if err != nil {
		fmt.Println("Issue occurred while creating folder: ", err)
		return false
	}
	if dirExists(path) {
		if fileExist(dbPath) {
			return true // Database .db exists in folder
		} else {
			db_file, db_err := os.Create(dbPath) // creates database file
			if db_err != nil {
				fmt.Println("Error creating file: ", db_err)
				return false
			}
			defer db_file.Close()
			return true
		}
	} else {
		folder_err := os.Mkdir(path, 0755)
		if folder_err != nil { // Issue creating folder
			fmt.Println("Issue occurred while creating folder: ", folder_err)
			return false
		}
		// Folder has been created
		db_file, db_err := os.Create(dbPath) // creates database file
		if db_err != nil {
			fmt.Println("Error creating file: ", db_err)
			return true
		}
		defer db_file.Close()
	}
	return false
}

func InitalizeConfig() bool {
	cacheDIR, err := os.UserCacheDir()
	path := filepath.Join(cacheDIR, "IManager")
	configPath := filepath.Join(path, "imanager.db")

	if err != nil {
		fmt.Println("Error creating file:", err)
		return false
	}

	if dirExists(path) {
		if fileExist(configPath) {
			// Config exists
			return true
		} else {
			config_file, db_err := os.Create(configPath)
			if db_err != nil {
				fmt.Println("Error creating file:", db_err)
				return false
			}
			defer config_file.Close()
			return true
		}
	} else {
		folder_err := os.Mkdir(path, 0755)
		if folder_err != nil {
			fmt.Println("Issue Occurred while creating folder:", folder_err)
			return false
		}
		config_file, db_err := os.Create(configPath)
		if db_err != nil {
			fmt.Println("Error creating file:", db_err)
			return false
		}
		defer config_file.Close()
		return true
	}
}
