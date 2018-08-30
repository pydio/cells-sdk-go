package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/pydio/cells-sdk-go/client/config_service"
	"github.com/pydio/cells-sdk-go/config"
)

var listDatasources = &cobra.Command{
	Use:   "list-datasources",
	Short: "ld",
	Long:  `List all the datasources`,
	Run: func(cmd *cobra.Command, args []string) {

		//connects to the pydio api via the sdkConfig
		httpClient := config.GetHttpClient(config.DefaultConfig)
		apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
		if err != nil {
			log.Fatal(err.Error())
		}

		/*ListDataSourcesParams contains all the parameters to send to the API endpoint
		for the list data sources operation typically these are written to a http.Request */
		params := &config_service.ListDataSourcesParams{
			Context:    ctx,
			HTTPClient: httpClient,
		}

		//assigns the datasources data retrieved above in the results variable
		result, err := apiClient.ConfigService.ListDataSources(params)
		if err != nil {
			fmt.Printf("could not list workspaces: %s\n", err.Error())
			log.Fatal(err.Error())
		}

		//prints the name of the datasources retrieved previously
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
