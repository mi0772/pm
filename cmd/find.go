/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"mi0772/pm/database"
	"mi0772/pm/userio"

	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "search for label record into database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("specify query term : eg: pm find google.com")
		}
		term := args[0]

		err := database.AccessToDatabase()
		if err != nil {
			panic(err)
		}
		userio.DisplayResult(database.Search(term))

	},
}

func init() {
	rootCmd.AddCommand(findCmd)
}
