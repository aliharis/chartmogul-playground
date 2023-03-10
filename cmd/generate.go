package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aliharis/chartmogul-playground/utils/chartmogul"
	"github.com/aliharis/chartmogul-playground/utils/helpers"
	cm "github.com/chartmogul/chartmogul-go/v3"
	"github.com/spf13/cobra"
)

// Generate random invoices for all the customers
var generateInvoices = &cobra.Command{
	Use:   "generate:invoices",
	Short: "Generate dummy invoices for all the customers",
	Run: func(cmd *cobra.Command, args []string) {
		// Check If the required datasource flag is set
		if !cmd.Flag("datasource").Changed {
			fmt.Println("The datasource flag is required.")
			return
		}

		// Get the ChartMogul API instance
		var api = chartmogul.GetAPI()

		// Get all the customers from ChartMogul API
		customers, err := api.ListCustomers(&cm.ListCustomersParams{})
		monthlyPlans := chartmogul.GetMonthlyPlans()

		// Check for errors
		if err != nil {
			fmt.Println(err)
			return
		}

		// Set the invoice sequence to 0
		sequence := 0

		// Iterate over the customers
		for _, customer := range customers.Entries {
			// Select a random plan for the customer
			rand.Seed(time.Now().UnixNano())
			plan := monthlyPlans[rand.Intn(len(monthlyPlans))]
			price := 5000
			if plan.Name == "Gold plan" {
				price = 10000
			}

			// Run the generator for the provided date range (interval of 1 month)
			// TODO: Parameterize the date range
			startDate := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
			endDate := time.Date(2019, 4, 1, 0, 0, 0, 0, time.UTC)
			date := startDate

			// Generate one subscription ID per customer to be used in the invoice
			subscriptionId := helpers.GenerateSubscriptionId()

			// Do something with the date
			fmt.Println("Generating the invoices for:", customer.Name)

			// Iterate over the customers
			for date.Before(endDate) || date.Equal(endDate) {
				// Increment the date by 1 month
				dueDate := date.AddDate(0, 0, 7)
				servicePeriodEndDate := date.AddDate(0, 1, 0)
				sequence = sequence + 1

				// Create a new invoice for the customer
				_, err = api.CreateInvoices([]*cm.Invoice{
					&cm.Invoice{
						ExternalID:         fmt.Sprintf("INV-%03d", sequence),
						Date:               date.Format("2006-01-02 00:00:00"),
						Currency:           "USD",
						DueDate:            dueDate.Format("2006-01-02 00:00:00"),
						CustomerExternalID: customer.ExternalID,
						DataSourceUUID:     cmd.Flag("datasource").Value.String(),
						LineItems: []*cm.LineItem{
							&cm.LineItem{
								Type:                      "subscription",
								SubscriptionExternalID:    fmt.Sprintf("sub_%s", subscriptionId),
								SubscriptionSetExternalID: fmt.Sprintf("set_%s", subscriptionId),
								PlanUUID:                  plan.UUID,
								ServicePeriodStart:        date.Format("2006-01-02 00:00:00"),
								ServicePeriodEnd:          servicePeriodEndDate.Format("2006-01-02 00:00:00"),
								AmountInCents:             price,
								Quantity:                  1,
								TransactionFeesCurrency:   "USD",
							},
						},
						Transactions: []*cm.Transaction{
							&cm.Transaction{
								Date:          date.Format("2006-01-02 00:00:00"),
								Type:          "payment",
								AmountInCents: &price,
								Result:        "successful",
							},
						},
					}}, customer.UUID)

				// Check for errors
				if err != nil {
					fmt.Println(err)
					return
				}

				fmt.Println("Generated invoices for period:", date.Format("2006-01-02 00:00:00"), "to", servicePeriodEndDate.Format("2006-01-02 00:00:00"))

				// Increment the date by 1 month
				date = date.AddDate(0, 1, 0)
			}

			fmt.Println("")
		}
	},
}

func init() {
	rootCmd.AddCommand(generateInvoices)

	// Add the required flags
	generateInvoices.Flags().StringP("datasource", "d", "", "The ChartMogul data source UUID")
}
