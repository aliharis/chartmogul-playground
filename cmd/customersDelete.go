package cmd

import (
	"fmt"

	"github.com/aliharis/chartmogul-playground/utils/chartmogul"
	cm "github.com/chartmogul/chartmogul-go/v3"
	"github.com/spf13/cobra"
)

// customersDeleteCmd represents the customersDelete command
var customersDeleteCmd = &cobra.Command{
	Use:   "customers:delete",
	Short: "Delete all the customers on your account.",
	Run: func(cmd *cobra.Command, args []string) {
		var api = chartmogul.GetAPI()

		// Get all the customers from ChartMogul API
		customers, err := api.ListCustomers(&cm.ListCustomersParams{})

		// Check for errors
		if err != nil {
			fmt.Println(err)
		}

		// Iterate over the customers
		for _, customer := range customers.Entries {
			// Delete the customer
			err := api.DeleteCustomer(customer.UUID)

			// Check for errors
			if err != nil {
				fmt.Println(err)
			}
		}

		fmt.Println("All the customers has been deleted.")
	},
}

func init() {
	rootCmd.AddCommand(customersDeleteCmd)
}
