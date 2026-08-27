/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"IManager-Src/database"
	"IManager-Src/utils"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "First command that should be ran to create necessary files.",
	Long: `Run this command to create IManager folder inside of local temporary directory
	
	This will create needed config file & database file to proceed with operations`,
	Run: func(cmd *cobra.Command, args []string) {
		/*
			Steps:
				1. InitalizeDatabase - creates cache folder, create IManager.db
				2. InitalizeConfig - creates config.json

		*/
		//utils.InitalizeDatabase() | f4b3e9a2-7d8c-4b51-9e63-2a81f5c6d7b4 - inbuilt into InitalizeConfig
		Config, Result := utils.Config()
		utils.Debug("TESTING", Config, Result)
		if Result { // File is found with data
			utils.Debug("init.go - ", Config)
			if Config.Initialized == true {
				utils.Information("Config file already exists")
			}
			if Config.DeliveredPath != "" { // Path is not empty
				utils.Information("Data inside config has been located.")
			}
			if utils.FileExist(Config.DeliveredPath) {
				utils.Information("Database has been located")
			} else {
				utils.InitalizeConfig()
			}
		} else {
			utils.Information("Detecting Operating System")
			utils.Information("Locating Cache Folder for OS")
			utils.Information("Cache Folder has been located")
			utils.Information("Creating database file")
			utils.Information("Creating config file")
			utils.InitalizeConfig()
			// CREATE Database Tables
			db, err := database.Open()
			if err != nil {
				utils.Error("Issued occurred while working with the database: ", err)
			}
			defer db.Close()

			if err := database.CreateBuyersTable(db); err != nil {
				utils.Error("Issued occurred while working with the database: ", err)
			} else {
				utils.Success("Created BuyersTable inside of database")
			}
			if err := database.CreateInventoryTable(db); err != nil {
				utils.Error("Issued occurred while working with the database: ", err)
			} else {
				utils.Success("Created InventoryTable inside of database")
			}
			if err := database.CreateOrdersTable(db); err != nil {
				utils.Error("Issued occurred while working with the database: ", err)
			} else {
				utils.Success("Created OrdersTable inside of database")
			}
			if err := database.CreateReceiptsTable(db); err != nil {
				utils.Error("Issued occurred while working with the database: ", err)
			} else {
				utils.Success("Created ReceiptsTable inside of database")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
