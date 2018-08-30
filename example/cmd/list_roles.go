package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/pydio/cells-sdk-go/client/role_service"
	"github.com/pydio/cells-sdk-go/config"
	"github.com/pydio/cells-sdk-go/models"
)

var listRoles = &cobra.Command{
	Use:   "list-roles",
	Short: "Lise roles",
	Long:  "List roles on pydio cells and also technical roles such as user/group",
	Run: func(cmd *cobra.Command, args []string) {

		httpClient := config.GetHttpClient(config.DefaultConfig)
		apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
		if err != nil {
			log.Fatal(err)
		}

		params := &role_service.SearchRolesParams{
			Body:       &models.RestSearchRoleRequest{},
			Context:    ctx,
			HTTPClient: httpClient,
		}

		result, err := apiClient.RoleService.SearchRoles(params)
		if err != nil {
			fmt.Printf("could not list roles: %s\n", err.Error())
			log.Fatal(err)
		}

		if len(result.Payload.Roles) > 0 {
			fmt.Printf("Found %d roles\n", len(result.Payload.Roles))
			for _, u := range result.Payload.Roles {
				//fmt.Println("  -- " + u.Label)
				if u.GroupRole == true {
					fmt.Printf(" -- %s __ GROUP ROLE \n", u.Label)
				} else if u.UserRole == true {
					fmt.Printf(" -- %s __ USER ROLE \n", u.Label)
				} else {
					fmt.Println(" -- " + u.Label)

				}
			}
		}
	},
}

func init() {
	ExampleCmd.AddCommand(listRoles)
}
