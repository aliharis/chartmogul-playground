package cmd

import (
	"encoding/csv"
	"fmt"
	"math/rand"
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

var customersSetSalesRepCmd = &cobra.Command{
	Use:   "customers:set-sales-rep",
	Short: "Read a CSV file with the sales rep data and set the sales rep for each customer.",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the file path from the flag
		filePath := cmd.Flag("file").Value.String()

		// open the CSV file
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		// create a new CSV reader
		reader := csv.NewReader(file)

		// read the header row and make sure it has "Name" and "External ID" fields
		header, err := reader.Read()
		if err != nil {
			fmt.Println("Error reading CSV file:", err)
			return
		}

		// TODO run a check to make sure the header has the required fields

		// map the header columns to struct fields
		fields := make(map[string]int)
		for i, name := range header {
			fields[name] = i
		}

		// Read the sales rep name from CSV and put it in a slice
		var salesReps []string
		for {
			// read a row
			row, err := reader.Read()
			if err != nil {
				break
			}

			// get the sales rep name
			salesRep := row[fields["Name"]]
			salesReps = append(salesReps, salesRep)
		}

		// Get all the customers from the API
		var api = chartmogul.GetAPI()
		customers, err := api.ListCustomers(&cm.ListCustomersParams{})
		if err != nil {
			fmt.Println(err)
		}

		// Iterate over the customers
		for _, customer := range customers.Entries {
			// Get a random sales rep name
			salesRep := salesReps[rand.Intn(len(salesReps))]

			// Set the sales rep for the customer
			api.AddCustomAttributesToCustomer(customer.UUID,
				[]*cm.CustomAttribute{
					{
						Type:  "String",
						Key:   "salesRep",
						Value: salesRep},
				})

			fmt.Println(salesRep, " has been set for customer", customer.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(customersDeleteCmd)
	rootCmd.AddCommand(customersListCmd)

	// Customer Set Sales Rep
	rootCmd.AddCommand(customersSetSalesRepCmd)
	customersSetSalesRepCmd.PersistentFlags().String("file", "", "Path to the sales rep file.")
}
