package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/pydio/cells-sdk-go/client/role_service"
	"github.com/pydio/cells-sdk-go/config"
	"github.com/pydio/cells-sdk-go/models"
)

var deleteRoles = &cobra.Command{
	Use:   "delete-roles",
	Short: "dr",
	Long:  "delete roles on cells",
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

		roleToDelete := "User nonAdminUser"

		if len(result.Payload.Roles) > 0 {
			fmt.Printf("Found %d roles\n", len(result.Payload.Roles))

			fmt.Println("Deletes role set in roleToDelete")
			for _, u := range result.Payload.Roles {

				if u.Label == roleToDelete {

					para := &role_service.DeleteRoleParams{
						UUID:       u.UUID,
						Context:    ctx,
						HTTPClient: httpClient,
					}

					_, err := apiClient.RoleService.DeleteRole(para)
					if err != nil {
						log.Fatal("Could not delete role ", err)
					}

				}
			}
			log.Println("ROLE DELETED")
		}
	},
}

func init() {
	ExampleCmd.AddCommand(deleteRoles)
}
