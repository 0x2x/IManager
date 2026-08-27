/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"IManager-Src/utils"
	"fmt"

	"github.com/spf13/cobra"
)

type ApplicationData struct {
	initialized    bool   `json:"init"`
	pending_path   string `json:"pending"`
	inventory_path string `json:inventory`
	orders_path    string `json:"pending"`
	delivered_path string `json:inventory`
	receipts_path  string `json:inventory`
}

type InventoryData struct {
	item_id int32 `json:"item_id"`
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		rGID := utils.GenerateID()
		fmt.Println("init called\tBase: " + rGID)
		/*
			Steps:
				1. InitalizeDatabase - creates cache folder, create IManager.db
				2. InitalizeConfig - creates config.json

		*/
		//utils.InitalizeDatabase() | f4b3e9a2-7d8c-4b51-9e63-2a81f5c6d7b4 - inbuilt into InitalizeConfig
		utils.InitalizeConfig()

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
