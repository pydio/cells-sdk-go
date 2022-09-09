package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/pydio/cells-sdk-go/v4/client/meta_service"
	"github.com/pydio/cells-sdk-go/v4/models"
	"github.com/pydio/cells-sdk-go/v4/transport/rest"
)

var (
	lsPath string
)

var listFiles = &cobra.Command{
	Use: "ls",
	Long: `List files in your Pydio Cells server.
	
	Example:
	$ go run main.go -u https://pydio.example.com -l admin -p password -f common-files/* ls 
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

func init() {
	listFiles.Flags().StringVarP(&lsPath, "filepath", "f", "/*", "File or folder path, suffix by '/*' to list children")
	ExampleCmd.AddCommand(listFiles)
}
