package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/pydio/cells-sdk-go/transport/rest"

	"github.com/spf13/cobra"

	"github.com/pydio/cells-sdk-go/client/meta_service"
	"github.com/pydio/cells-sdk-go/models"
)

var (
	lsPath string
)

var listFiles = &cobra.Command{
	Use:  "ls",
	Long: `List files on pydio cells`,
	Run: func(cmd *cobra.Command, args []string) {

		//connects to the pydio api via the sdkConfig
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

		//assigns the files data retrieved above in the results variable
		result, err := apiClient.MetaService.GetBulkMeta(params)
		if err != nil {
			fmt.Printf("Could not list meta: %s\n", err.Error())
			log.Fatal(err)
		}

		//prints the path therefore the name of the files listed
		if len(result.Payload.Nodes) > 0 {
			fmt.Printf("Info for %s, got %d results\n", lsPath, len(result.Payload.Nodes))
			for _, u := range result.Payload.Nodes {
				fType := "F"
				if u.Type == models.TreeNodeTypeCOLLECTION {
					fType = "D"
				}
				fmt.Printf("  - [%s]\t%s\t%s\n", fType, u.Path, u.Size)
			}
		}

	},
}

func init() {
	listFiles.Flags().StringVarP(&lsPath, "filepath", "f", "/*", "File or folder path, use /* to list children")
	ExampleCmd.AddCommand(listFiles)
}
