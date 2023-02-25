# chartmogul-playground

ChartMogul playground is a CLI tool to interact with ChartMogul API and primarily use for importing data to ChartMogul.

## Pre-requisite

Setup your ChartMogul API key as an environment variable

`export CHARTMOGUL_API_KEY=YOUR_API_KEY`

This is the documentation for the Go Application.
Description

This is a simple Go application that performs a specific task. The application has been developed using Go 1.16.

## Installation

To clone this application, run the following command:

`git clone https://github.com/aliharis/chartmogul-playground`

Dependencies

This application uses Go modules to manage dependencies. To install all dependencies, run the following command from the root directory of the project:

`go mod download`

This will download all the required dependencies and store them in the local cache.

### Running the application

The application requires a ChartMogul API key. Set the API key as environment variable in the system.

`export CHARTMOGUL_API_KEY=YOUR_API_KEY`

To run the application, use the following command from the root directory of the project:

`go run main.go`

This will build and run the application, and you should see the output of the application in the console.

## Usage

```
Usage:
  chartmogul-playground [command]

Available Commands:
  completion              Generate the autocompletion script for the specified shell
  customers:delete        Delete all the customers on your account.
  customers:list          Return a list of all the customers in your ChartMogul account.
  customers:set-sales-rep Read a CSV file with the sales rep data and set the sales rep for each customer.
  datasources:list        List all the datasources on your account.
  generate:invoices       Generate dummy invoices for all the customers
  help                    Help about any command
  import:customers        A brief description of your command
  import:plans            Import plans from a CSV file
  insights:mrr            MRR for the specified time period.

Flags:
  -h, --help   help for chartmogul-playground

Use "chartmogul-playground [command] --help" for more information about a command.
```

## Examples

| Action                                                  | Command                                                                                              |
| ------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| List datasources                                        | `chartmogul-playground datasources:list`                                                             |
| List all the customers                                  | `chartmogul-playground customers:list`                                                               |
| Import plans from a CSV file                            | `chartmogul-playground import:plans --file=resources/plans.csv --datasource=DATASOURCE_UUID`         |
| Import customers from a CSV file                        | `chartmogul-playground import:customers --file=resources/customers.csv --datasource=DATASOURCE_UUID` |
| Set random sales rep for customers froma CSV file       | `chartmogul-playground customers:set-sales-rep --file=resources/sales_reps.csv`                      |
| Generate dummy invoices and subscriptions for customers | `chartmogul-playground generate:invoices --datasource=DATASOURCE_UUID`                               |
| Retrieve MRR data for a period                          | `chartmogul-playground insights:mrr --from=2019-01-01 --to=2019-05-01`                               |
