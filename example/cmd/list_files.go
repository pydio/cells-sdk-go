package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"

	"github.com/pydio/cells-sdk-go/models"
	"github.com/pydio/cells-sdk-go/config"
	"github.com/pydio/cells-sdk-go/client/meta_service"

)

var listFiles = &cobra.Command{
	Use:   "list-files",
	Short: "lf",
	Long:  `List files on pydio cells`,
	Run: func(cmd *cobra.Command, args []string) {

		//connects to the pydio api via the sdkConfig
		httpClient := config.GetHttpClient(config.DefaultConfig)
		apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
		if err != nil {
			log.Println("USER : ",config.DefaultConfig.User)
			time.Sleep(100 * time.Millisecond)
			log.Fatalln("CANNOT CONNECT WITH THIS USER : ", err)
		}

		/*GetBulkMetaParams contains all the parameters to send to the API endpoint
		for the get bulk meta operation typically these are written to a http.Request
		*/
		params := &meta_service.GetBulkMetaParams{
			Body: &models.RestGetBulkMetaRequest{NodePaths: []string{
				//the workspaces from whom the files are listed
				"/*",
				"personal-files/*",
			}},
			Context:    ctx,
			HTTPClient: httpClient,
		}

		//assigns the files data retrieved above in the results variable
		result, err := apiClient.MetaService.GetBulkMeta(params)
		if err != nil {
			fmt.Printf("could not list meta: %s\n", err.Error())
			log.Fatal(err)
		}

		//prints the path therefore the name of the files listed
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
