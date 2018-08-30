package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/pydio/cells-sdk-go/client/user_service"
	"github.com/pydio/cells-sdk-go/config"
)

var listUsers = &cobra.Command{
	Use:   "list-users",
	Short: "lu",
	Long:  `List users on the pydio cells`,
	Run: func(cmd *cobra.Command, args []string) {

		httpClient := config.GetHttpClient(config.DefaultConfig)
		apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
		if err != nil {
			log.Fatal(err)
		}

		// query := api.RestSearchUserRequest{}
		params := &user_service.SearchUsersParams{
			Context:    ctx,
			HTTPClient: httpClient,
		}

		//assigns the users data retrieved above in the results variable
		result, err := apiClient.UserService.SearchUsers(params)
		if err != nil {
			fmt.Printf("could not list users: %s\n", err.Error())
			log.Fatal(err)
		}

		//prints the login of the users retrieved previously
		if len(result.Payload.Users) > 0 {
			fmt.Printf("Found %d users\n", len(result.Payload.Users))
			for _, u := range result.Payload.Users {
				fmt.Println("  - " + u.Login)
			}
		}

	},
}

func init() {
	ExampleCmd.AddCommand(listUsers)
}
