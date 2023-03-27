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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {
		err := database.AccessToDatabase()
		if err != nil {
			panic(err)
		}
		fmt.Println()
		userio.DisplayResult(database.Search("*"))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
