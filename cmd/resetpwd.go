/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetpwdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetpwdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
