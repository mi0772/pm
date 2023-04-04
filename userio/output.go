package userio

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"mi0772/pm/models"
	"os"
	"sort"
)

func DisplayResult(result []models.Entry) {
	fmt.Println()
	tbl := table.NewWriter()
	tbl.SetStyle(table.StyleLight)
	tbl.SetOutputMirror(os.Stdout)

	sort.Slice(result, func(i, j int) bool {
		return result[i].Label < result[j].Label
	})

	tbl.AppendHeader(table.Row{"Label", "Account", "Password", "Id"})

	for _, widget := range result {
		if widget.ModifiedAt.IsZero() {
			//tbl.AppendRow([]interface{}{widget.Label, widget.Account, color.YellowString(widget.Password), widget.Id, widget.CreatedAt.Format("2006-01-02 15:04:05"), "-"})
			tbl.AppendRow([]interface{}{widget.Label, widget.Account, color.YellowString(widget.Password), widget.Id})
		} else {
			//tbl.AppendRow([]interface{}{widget.Label, widget.Account, color.YellowString(widget.Password), widget.Id, widget.CreatedAt.Format("2006-01-02 15:04:05"), widget.ModifiedAt.Format("2006-01-02 15:04:05")})
			tbl.AppendRow([]interface{}{widget.Label, widget.Account, color.YellowString(widget.Password), widget.Id})

		}
		tbl.AppendSeparator()
	}

	tbl.Render()
	if len(result) == 0 {
		fmt.Println("no records found")
	}
}
