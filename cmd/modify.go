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
	Long:  ``,
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
}
