package cmd

import (
	"fmt"
	"os"

	"github.com/aliharis/chartmogul-playground/utils/chartmogul"
	cm "github.com/chartmogul/chartmogul-go/v3"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// customersListCmd represents the customersList command
var insightsMRRCmd = &cobra.Command{
	Use:   "insights:mrr",
	Short: "MRR for the specified time period.",
	Run: func(cmd *cobra.Command, args []string) {
		// Check If required flags are set (from, to)
		if !cmd.Flags().Changed("from") || !cmd.Flags().Changed("to") {
			fmt.Println("Please specify the required flags: --from, --to")
			return
		}

		// Get an instance of the ChartMogul API
		var api = chartmogul.GetAPI()

		// Get the MRR data from API
		mrrData, err := api.MetricsRetrieveMRR(&cm.MetricsFilter{
			StartDate: cmd.Flag("from").Value.String(),
			EndDate:   cmd.Flag("to").Value.String(),
			Interval:  cmd.Flag("interval").Value.String(),
		})

		if err != nil {
			fmt.Println(err)
			return
		}

		// Print the MRR data to a table
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Month", "MRR", "MRR Churn", "MRR Contraction", "MRR Expansion", "MRR New Business"})

		// Iterate MRR data
		for _, mrr := range mrrData.Entries {
			// Add the customer to the table
			table.Append([]string{
				mrr.Date,
				fmt.Sprintf("%.2f", mrr.MRR/100),
				fmt.Sprintf("%.2f", mrr.MRRChurn/100),
				fmt.Sprintf("%.2f", mrr.MRRContraction/100),
				fmt.Sprintf("%.2f", mrr.MRRExpansion/100),
				fmt.Sprintf("%.2f", mrr.MRRNewBusiness/100),
			})
		}

		// Render the table
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(insightsMRRCmd)
	insightsMRRCmd.PersistentFlags().String("from", "", "The start date of the required period of data. An ISO-8601 formatted date, e.g. 2015-05-12")
	insightsMRRCmd.PersistentFlags().String("to", "", "The end date of the required period of data. An ISO-8601 formatted date, e.g. 2015-05-12")
	insightsMRRCmd.PersistentFlags().String("interval", "month", "The interval of the required period of data. One of day, week, month, quarter, year")
}
