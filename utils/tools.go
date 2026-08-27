package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type ApplicationData struct {
	Initialized   bool   `json:"init"`
	PendingPath   string `json:"pending"`
	InventoryPath string `json:"inventory"`
	OrdersPath    string `json:"orders"`
	DeliveredPath string `json:"delivered"`
	ReceiptsPath  string `json:"receipts"`
}

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

func TempFileCreation(fileName string, fileExt string) (string, bool) {
	combo := fileName + "." + fileExt // example.test
	cacheDIR, err := os.UserCacheDir()
	path := filepath.Join(cacheDIR, "IManager")
	dbPath := filepath.Join(path, combo)
	if err != nil {
		fmt.Println("Issue occurred while creating folder: ", err)
		return dbPath, false
	}
	if dirExists(path) {
		if fileExist(dbPath) {
			return dbPath, true // Database .db exists in folder
		} else {
			db_file, db_err := os.Create(dbPath) // creates database file
			if db_err != nil {
				fmt.Println("Error creating file: ", db_err)
				return "", false
			}
			defer db_file.Close()
			return dbPath, true
		}
	} else {
		folder_err := os.Mkdir(path, 0755)
		if folder_err != nil { // Issue creating folder
			fmt.Println("Issue occurred while creating folder: ", folder_err)
			return "", false
		}
		// Folder has been created
		db_file, db_err := os.Create(dbPath) // creates database file
		if db_err != nil {
			fmt.Println("Error creating file: ", db_err)
			return dbPath, true
		}
		defer db_file.Close()
	}
	return "", false
}

func InitalizeDatabase() (string, bool) { // May need to make OS specific?
	path, result := TempFileCreation("IManager", "db")
	return path, result
}

func InitalizeConfig() bool {
	config_file, cresult := TempFileCreation("IManager", "json")
	database_file, dresult := InitalizeDatabase()
	if cresult != true || dresult != true { // Issue occured while creating file
		return false
	}
	//FUNC Write to IManager.json
	// as of the moment initalize config for us: f4b3e9a2-7d8c-4b51-9e63-2a81f5c6d7b4
	data := ApplicationData{
		Initialized:   true,
		PendingPath:   database_file,
		InventoryPath: database_file,
		OrdersPath:    database_file,
		DeliveredPath: database_file,
		ReceiptsPath:  database_file,
	}
	fmt.Print(data)
	file, err := os.Create(config_file)
	if err != nil { // Issue occurs
		return false
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	err = encoder.Encode(data)
	if err != nil { // Issue occurs
		return false
	}
	return true
}
