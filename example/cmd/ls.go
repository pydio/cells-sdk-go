package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/pydio/cells-sdk-go/v5/client/meta_service"
	"github.com/pydio/cells-sdk-go/v5/client/workspace_service"
	"github.com/pydio/cells-sdk-go/v5/models"
	"github.com/pydio/cells-sdk-go/v5/transport/rest"
)

var (
	lsPath   string
	lsWsSlug string
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Set of sample commands to lists objects on the server via the API",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var listFiles = &cobra.Command{
	Use: "files",
	Long: `List files in your Pydio Cells server.
	
	Example:
	$ go run main.go -u https://pydio.example.com -l admin -p password -f common-files/* ls files 
	`,
	Run: func(cmd *cobra.Command, args []string) {

		// Connect to the pydio api via the sdkConfig
		ctx, apiClient, err := rest.GetClient(DefaultConfig, false)
		if err != nil {
			log.Fatal(err)
		}

		params := &meta_service.GetBulkMetaParams{
			Body: &models.RestGetBulkMetaRequest{NodePaths: []string{
				strings.TrimSuffix(lsPath, "/"),
			}},
			Context: ctx,
		}

		// Assign the files data retrieved above to the result variable
		result, err := apiClient.MetaService.GetBulkMeta(params)
		if err != nil {
			fmt.Printf("Could not list meta at %s\n", lsPath)
			log.Fatal(err)
		}

		// Print the path and therefore the name of the found files
		if len(result.Payload.Nodes) > 0 {
			fmt.Printf("Listing %s. Found %d results\n", lsPath, len(result.Payload.Nodes))
			for _, u := range result.Payload.Nodes {
				fType := "F"
				if *u.Type == models.TreeNodeTypeCOLLECTION {
					fType = "D"
				}
				fmt.Printf("  - [%s]\t%s\t%s\n", fType, u.Path, u.Size)
			}
		}
	},
}

var listWorkspaces = &cobra.Command{
	Use: "wss",
	Long: `List workspaces the admin user can see in the Pydio Cells server.
	
	Example:
	$ go run main.go -u https://pydio.example.com -l admin -p password ls wss 
	`,
	Run: func(cmd *cobra.Command, args []string) {

		// Connect to the pydio api via the sdkConfig
		ctx, apiClient, err := rest.GetClient(DefaultConfig, false)
		if err != nil {
			log.Fatal(err)
		}

		params := workspace_service.NewSearchWorkspacesParamsWithContext(ctx)
		queries := make([]*models.IdmWorkspaceSingleQuery, 1)
		queries[0] = &models.IdmWorkspaceSingleQuery{Slug: lsWsSlug}
		params.Body = &models.RestSearchWorkspaceRequest{
			Queries:   queries,
			Operation: models.NewServiceOperationType(models.ServiceOperationTypeAND),
		}
		result, err := apiClient.WorkspaceService.SearchWorkspaces(params)
		if err != nil {
			log.Fatal("could not list WS "+lsWsSlug, err)
		}

		// Print the path and therefore the name of the found files
		if result.Payload.Total > 0 {
			fmt.Printf("Searching with %s. Found %d results\n", lsWsSlug, result.Payload.Total)
			for _, u := range result.Payload.Workspaces {
				fmt.Printf("  - [%s]\t%s\t%s\n", u.Slug, u.Label, u.UUID)
			}
		} else {
			fmt.Printf("No error but no workspace found for %s.\n", lsWsSlug)
		}
	},
}

func init() {
	listFiles.Flags().StringVarP(&lsPath, "filepath", "f", "/*", "File or folder path, suffix by '/*' to list children")
	lsCmd.AddCommand(listFiles)
	listWorkspaces.Flags().StringVarP(&lsWsSlug, "slug", "s", "", "Slug full of with '*' wild char")
	lsCmd.AddCommand(listWorkspaces)
	ExampleCmd.AddCommand(lsCmd)
}
