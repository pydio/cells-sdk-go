package cmd

import (
	"log"
	"fmt"

	"github.com/pydio/cells-sdk-go/client/meta_service"
	"github.com/pydio/cells-sdk-go/config"
	"github.com/pydio/cells-sdk-go/models"
	"github.com/spf13/cobra"
)



var listFiles = &cobra.Command{
	Use:   "list-files",
	Short: "List files",
	Long:  `List files on pydio cells https://demo.pydio.com, only requires the secret key`,
	Run: func(cmd *cobra.Command, args []string) {

		if secret == "" {
			log.Fatal("Please provide the secret key")
		}



		httpClient := config.GetHttpClient(config.DefaultConfig)
		apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
		if err != nil {
			log.Fatal(err)
		}

		params := &meta_service.GetBulkMetaParams{
			Body: &models.RestGetBulkMetaRequest{NodePaths: []string{
				"/*",
				//"ds1/*",
				"personal-files/*",
			}},
			Context:    ctx,
			HTTPClient: httpClient,
		}

		result, err := apiClient.MetaService.GetBulkMeta(params)
		if err != nil {
			fmt.Printf("could not list meta: %s\n", err.Error())
			log.Fatal(err)
		}

		if len(result.Payload.Nodes) > 0 {
			fmt.Printf("* %d meta\n", len(result.Payload.Nodes))
			for _, u := range result.Payload.Nodes {
				fmt.Println("  -", u.Path)
			}
		}

	},
}

func init() {
	ExampleCmd.AddCommand(listFiles)
}
