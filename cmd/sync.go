package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "syncronize database with remote location",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sync called")
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
