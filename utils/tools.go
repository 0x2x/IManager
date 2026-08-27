package utils

import (
	"errors"
	"fmt"
	"os"
	"runtime"

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

func InitalizeDatabase() {
	switch runtime.GOOS {
	case "windows":
		path := "C:\\Users\\name\\AppData\\Local\\Temp\\IManager"
		if dirExists(path) {
			if fileExist(path + "/imanager.db") {
				// Database exists
			} else {
				db_file, db_err := os.Create(path + "imanager.db")
				if db_err != nil {
					fmt.Println("Error creating file:", db_err)
					return
				}
				defer db_file.Close()
			}
		} else {
			folder_err := os.Mkdir(path, 0755)
			if folder_err != nil {
				fmt.Println("Issue Occurred while creating folder:", folder_err)
				return
			}
			db_file, db_err := os.Create(path + "imanager.db")
			if db_err != nil {
				fmt.Println("Error creating file:", db_err)
				return
			}
			defer db_file.Close()
		}
	case "darwin":
		path := "~/Library/Caches/IManager"
		if dirExists(path) {
			if fileExist(path + "/imanager.db") {
				// passes
			} else {
				db_file, db_err := os.Create(path + "/imanager.db") // creates database file
				if db_err != nil {
					fmt.Println("Error creating file: ", db_err)
					return
				}
				defer db_file.Close()
			}
		} else {
			folder_err := os.Mkdir(path, 0755)
			if folder_err != nil { // Issue creating folder
				fmt.Println("Issue occurred while creating folder: ", folder_err)
				return
			}
			// Folder has been created
			db_file, db_err := os.Create(path + "/imanager.db") // creates database file
			if db_err != nil {
				fmt.Println("Error creating file: ", db_err)
				return
			}
			defer db_file.Close()
		}
	case "linux":
		path := "~/.cache/IManager"
		if dirExists(path) {
			if fileExist(path + "/imanager.db") {
				// passes
			} else {
				db_file, db_err := os.Create(path + "/imanager.db") // creates database file
				if db_err != nil {
					fmt.Println("Error creating file: ", db_err)
					return
				}
				defer db_file.Close()
			}
		} else {
			folder_err := os.Mkdir(path, 0755)
			if folder_err != nil { // Issue creating folder
				fmt.Println("Issue occurred while creating folder: ", folder_err)
				return
			}
			// Folder has been created
			db_file, db_err := os.Create(path + "/imanager.db") // creates database file
			if db_err != nil {
				fmt.Println("Error creating file: ", db_err)
				return
			}
			defer db_file.Close()
		}
	default:
		fmt.Printf("Unknown Operating Type. Please create an Issue inside the github repo. \n\tOS: %s", runtime.GOOS)
	}
}

func InitalizeConfig() {
	switch runtime.GOOS {
	case "windows":
		path := "C:\\Users\\name\\AppData\\Local\\Temp\\IManager"
		if dirExists(path) {
			if fileExist(path + "\\app.config") {
				// Database exists
			} else {
				db_file, db_err := os.Create(path + "\\app.config")
				if db_err != nil {
					fmt.Println("Error creating file:", db_err)
					return
				}
				defer db_file.Close()
			}
		} else {
			folder_err := os.Mkdir(path, 0755)
			if folder_err != nil {
				fmt.Println("Issue Occurred while creating folder:", folder_err)
				return
			}
			db_file, db_err := os.Create(path + "imanager.db")
			if db_err != nil {
				fmt.Println("Error creating file:", db_err)
				return
			}
			defer db_file.Close()
		}
	case "darwin":

	case "linux":
	}
}
