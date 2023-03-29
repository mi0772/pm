package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"mi0772/pm/database"
	"mi0772/pm/userio"
)

// resetpwdCmd represents the resetpwd command
var resetpwdCmd = &cobra.Command{
	Use:   "resetpwd",
	Short: "change master password",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := database.AccessToDatabase()
		if err != nil {
			panic(err)
		}
		fmt.Println()

		newPwd, err := userio.ReadNewMasterPassword("enter new master password")

		database.ChangeMasterPassword(string(newPwd))
		fmt.Println("master password sucessfully changed")

	},
}

func init() {
	rootCmd.AddCommand(resetpwdCmd)
}
