package cmd

import (
	"fmt"
	"os"

	"github.com/aliharis/chartmogul-playground/utils/chartmogul"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var datasourcesListCmd = &cobra.Command{
	Use:   "datasources:list",
	Short: "List all the datasources on your account.",
	Run: func(cmd *cobra.Command, args []string) {
		var api = chartmogul.GetAPI()

		// Get the datasources from ChartMogul API
		datasources, err := api.ListDataSources()

		// Check for errors
		if err != nil {
			fmt.Println(err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "UUID"})

		// Iterate over the datasources
		for _, datasource := range datasources.DataSources {
			table.Append([]string{
				datasource.Name,
				datasource.UUID,
			})
		}

		// Render the table
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(datasourcesListCmd)
}
