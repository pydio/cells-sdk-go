package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/pydio/cells-sdk-go/client/workspace_service"
	"github.com/pydio/cells-sdk-go/config"
	"github.com/pydio/cells-sdk-go/models"
)

var listWorkspaces = &cobra.Command{
	Use:   "list-workspaces",
	Short: "lw",
	Long:  `List all the workspaces`,
	Run: func(cmd *cobra.Command, args []string) {

		httpClient := config.GetHttpClient(config.DefaultConfig)
		apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
		if err != nil {
			log.Fatal(err)
		}

		//retrieves the users using the searchWorkspacesParams function
		params := &workspace_service.SearchWorkspacesParams{
			Body:       &models.RestSearchWorkspaceRequest{CountOnly: true},
			Context:    ctx,
			HTTPClient: httpClient,
		}

		//assigns the workspaces data retrieved above in the results variable
		result, err := apiClient.WorkspaceService.SearchWorkspaces(params)
		if err != nil {
			fmt.Printf("could not list workspaces: %s\n", err.Error())
			log.Fatal(err)
		}

		//prints the workspace label
		if len(result.Payload.Workspaces) > 0 {
			fmt.Printf("* %d workspace found\n", len(result.Payload.Workspaces))
			for _, u := range result.Payload.Workspaces {
				fmt.Println("  - " + u.Slug)
			}
		}

	},
}

func init() {
	ExampleCmd.AddCommand(listWorkspaces)
}
