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
var customersListCmd = &cobra.Command{
	Use:   "customers:list",
	Short: "Return a list of all the customers in your ChartMogul account.",
	Run: func(cmd *cobra.Command, args []string) {
		// List the customers from ChartMogul API
		var api = chartmogul.GetAPI()

		// Get the customers from ChartMogul API
		customers, err := api.ListCustomers(&cm.ListCustomersParams{})

		// Check for errors
		if err != nil {
			fmt.Println(err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Company", "Email", "Country", "State", "City", "ZIP"})

		// Iterate over the customers
		for _, customer := range customers.Entries {
			// Add the customer to the table
			table.Append([]string{
				customer.Name,
				customer.Company,
				customer.Email,
				customer.Country,
				customer.State,
				customer.City,
				customer.Zip,
			})
		}

		// Render the table
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(customersListCmd)
}
