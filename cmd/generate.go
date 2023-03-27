/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/sethvargo/go-password/password"
	"log"
	"mi0772/pm/models"
	"mi0772/pm/userio"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "a password generator",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var length, numspecial int
		var ok bool
		if length, ok = userio.GetParameterAsInt(args, 0); !ok {
			fmt.Println("wrong arguments, specify length and special char num, eg: pm generate <length> <special char nums>")
		}
		if numspecial, ok = userio.GetParameterAsInt(args, 1); !ok {
			fmt.Println("wrong arguments, specify length and special char num, eg: pm generate <length> <special char nums>")
		}
		var gc = &models.GenerateCommand{Length: length, Special: numspecial}

		fmt.Printf("generate %v password length with %v special chars\n", gc.Length, gc.Special)
		res, err := password.Generate(gc.Length, gc.Length/4, gc.Special, false, false)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("This is your password:", res)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
