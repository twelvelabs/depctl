package cmd

import (
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/twelvelabs/depctl/internal/core"
)

func NewListCmd(app *core.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List dependencies",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			table := tablewriter.NewWriter(app.IO.Out)

			table.SetHeader([]string{
				"Group",
				"Name",
				"Description",
				"URL",
			})
			table.SetAutoFormatHeaders(false)
			table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
			table.SetHeaderLine(true)

			table.SetBorders(tablewriter.Border{
				Left:   true,
				Right:  true,
				Top:    false,
				Bottom: false,
			})
			table.SetCenterSeparator("|")
			table.SetAutoWrapText(false)

			for key, deps := range app.Config.Dependencies {
				for _, dep := range deps {
					table.Append([]string{
						key,
						dep.Name,
						dep.Description,
						app.UI.Formatter.Underline(dep.URL),
					})
				}
			}

			table.Render()
			return nil
		},
	}

	return cmd
}
