/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"mi0772/pm/database"
	"mi0772/pm/models"
	"mi0772/pm/userio"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new password into database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := database.AccessToDatabase()
		if err != nil {
			panic(err)
		}

		var label, account, password string
		fmt.Println()
		if label, err = userio.ReadInput("label"); err != nil {
			fmt.Println(err)
		}

		if account, err = userio.ReadInput("user id"); err != nil {
			fmt.Println(err)
		}

		if password, err = userio.ReadInput("password"); err != nil {
			fmt.Println(err)
		}
		var sc = &models.StoreCommand{Label: &label, Account: &account, Password: &password}

		fmt.Printf("memorized new password for label:%s, account:%s, password:%s\n", *sc.Label, *sc.Account, *sc.Password)
		database.Memorize(*sc.Label, *sc.Account, *sc.Password)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
