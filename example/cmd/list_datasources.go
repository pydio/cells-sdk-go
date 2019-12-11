package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	
	"github.com/pydio/cells-sdk-go/client/config_service"
)

var listDatasources = &cobra.Command{
	Use:   "list-datasources",
	Short: "ld",
	Long:  `List all the datasources`,
	Run: func(cmd *cobra.Command, args []string) {

		// Connect to the Pydio API via the sdkConfig
		ctx, apiClient, err := GetApiClient(DefaultConfig)
		if err != nil {
			log.Fatal(err.Error())
		}

		/*ListDataSourcesParams contains all the parameters to send to the API endpoint
		for the list data sources operation typically these are written to a http.Request */
		params := &config_service.ListDataSourcesParams{Context: ctx}

		// Assign the datasources data retrieved above in the results variable
		result, err := apiClient.ConfigService.ListDataSources(params)
		if err != nil {
			fmt.Printf("could not list workspaces: %s\n", err.Error())
			log.Fatal(err.Error())
		}

		// Print the name of the datasources retrieved previously
		if len(result.Payload.DataSources) > 0 {
			fmt.Printf("* %d datasources	\n", len(result.Payload.DataSources))
			for _, u := range result.Payload.DataSources {
				fmt.Println("  - " + u.Name)
			}
		}

	},
}

func init() {
	ExampleCmd.AddCommand(listDatasources)
}
