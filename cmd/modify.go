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

// modifyCmd represents the modify command
var modifyCmd = &cobra.Command{
	Use:   "modify",
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

		var id int
		var ok bool
		if id, ok = userio.GetParameterAsInt(args, 0); !ok {
			fmt.Println("specify record Id to update : eg: pm modify 1")
		}

		if _, ok = database.GetById(id); ok == false {
			fmt.Printf("record with Id %d does not exists\n", id)
		}

		var label, account, password string
		fmt.Println()
		if label, err = userio.ReadInputNotBlank("label"); err != nil {
			fmt.Println(err)
		}

		if account, err = userio.ReadInputNotBlank("user id"); err != nil {
			fmt.Println(err)
		}

		if password, err = userio.ReadInputNotBlank("password"); err != nil {
			fmt.Println(err)
		}

		fmt.Printf("updated entry for label:%s, account:%s, password:%s\n", label, account, password)
		database.Update(id, label, account, password)
	},
}

func init() {
	rootCmd.AddCommand(modifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
