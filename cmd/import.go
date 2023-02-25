package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/aliharis/chartmogul-playground/utils/chartmogul"
	cm "github.com/chartmogul/chartmogul-go/v3"
	"github.com/spf13/cobra"
)

// importCustomersCmd represents the importCustomers command
var importCustomersCmd = &cobra.Command{
	Use:   "import:customers",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if the required flag is set
		if !cmd.Flag("file").Changed {
			fmt.Println("The file flag is required.")
			return
		}

		// Check If the datasource flag is set
		if !cmd.Flag("datasource").Changed {
			fmt.Println("The datasource flag is required.")
			return
		}

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

		// map the header columns to struct fields
		fields := make(map[string]int)
		nameIndex := -1
		externalIdIndex := -1

		for i, name := range header {
			if name == "Name" {
				nameIndex = i
			} else if name == "External ID" {
				externalIdIndex = i
			}

			fields[name] = i
		}

		// CHeck If the header has the required fields
		if nameIndex == -1 || externalIdIndex == -1 {
			fmt.Println("Error: CSV file header is missing 'Name' or 'External ID' fields")
			return
		}

		// Get the ChartMogul API instance
		var api = chartmogul.GetAPI()

		// iterate over the remaining rows in the CSV file
		for {
			row, err := reader.Read()
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				fmt.Println("Error reading CSV file:", err)
				return
			}

			// create a new customer
			api.CreateCustomer(&cm.NewCustomer{
				DataSourceUUID: cmd.Flag("datasource").Value.String(),
				Name:           row[fields["Name"]],
				ExternalID:     row[fields["External ID"]],
				Email:          row[fields["Email"]],
				Company:        row[fields["Company"]],
				Country:        row[fields["Country"]],
				State:          row[fields["State"]],
				City:           row[fields["City"]],
				Zip:            row[fields["Zip"]],
			})

			fmt.Println("Customer imported:", row[fields["Name"]])
		}
	},
}

var importPlansCmd = &cobra.Command{
	Use:   "import:plans",
	Short: "Import plans from a CSV file",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if the required flag is set
		if !cmd.Flag("file").Changed {
			fmt.Println("The file flag is required.")
			return
		}

		// Check If the datasource flag is set
		if !cmd.Flag("datasource").Changed {
			fmt.Println("The datasource flag is required.")
			return
		}

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

		// Get the ChartMogul API instance
		var api = chartmogul.GetAPI()

		// iterate over the remaining rows in the CSV file
		for {
			row, err := reader.Read()
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				fmt.Println("Error reading CSV file:", err)
				return
			}

			// Convert the interval count to an integer
			intervalCount, err := strconv.ParseUint(row[fields["IntervalCount"]], 10, 32)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			// Create a new plan
			api.CreatePlan(&cm.Plan{
				DataSourceUUID: cmd.Flag("datasource").Value.String(),
				Name:           row[fields["Name"]],
				IntervalCount:  uint32(intervalCount),
				IntervalUnit:   row[fields["IntervalUnit"]],
				ExternalID:     row[fields["External ID"]],
			})

			fmt.Println("Plan imported:", row[fields["Name"]])
		}
	},
}

func init() {
	// Import Customers
	rootCmd.AddCommand(importCustomersCmd)
	importCustomersCmd.PersistentFlags().String("file", "", "Path to the file to import.")
	importCustomersCmd.PersistentFlags().String("datasource", "", "UUID of the data source to import the data to.")

	// Commands
	rootCmd.AddCommand(importPlansCmd)
	importPlansCmd.PersistentFlags().String("file", "", "Path to the file to import.")
	importPlansCmd.PersistentFlags().String("datasource", "", "UUID of the data source to import the data to.")
}
