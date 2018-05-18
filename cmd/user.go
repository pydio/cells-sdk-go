package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/pydio/cells-sdk-go/api"
	"github.com/pydio/cells-sdk-go/config"
)

var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "User-related commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var ListUserCmd = &cobra.Command{
	Use:   "ls",
	Short: "List users",
	Run: func(cmd *cobra.Command, args []string) {
		apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
		if err != nil {
			log.Fatal(err)
		}

		query := api.RestSearchUserRequest{}

		users, _, err := apiClient.UserServiceApi.SearchUsers(
			ctx,
			query,
		)
		if err != nil {
			fmt.Printf("could not list users: %s\n", err.Error())
			log.Fatal(err)
		}

		fmt.Printf("Found %d results\n", users.Total)
		if len(users.Groups) > 0 {
			fmt.Printf("* %d groups\n", len(users.Groups))
			for _, u := range users.Groups {
				fmt.Println("  - " + u.GroupLabel)
			}
		}

		if len(users.Users) > 0 {
			fmt.Printf("* %d users\n", len(users.Users))
			for _, u := range users.Users {
				fmt.Println("  - " + u.Login)
			}
		}

	},
}

var CreateUserCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a user",
	Run: func(cmd *cobra.Command, args []string) {
		apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
		if err != nil {
			log.Fatal(err)
		}

		newUser := api.IdmUser{
			GroupPath: "/",
			Login:     "randomUser",
			Password:  "password",
		}

		user, _, err := apiClient.UserServiceApi.PutUser(
			ctx,
			"randomUser",
			newUser,
		)
		if err != nil {
			log.Fatalf("could not create user %s: %s\n", "randomUser", err.Error())
		}

		fmt.Printf("Created user with login: %s and uuid: %s\n", user.Login, user.Uuid)

		ListUserCmd.Run(cmd, args)

	},
}

func init() {
	UserCmd.AddCommand(ListUserCmd)
	UserCmd.AddCommand(CreateUserCmd)

	RootCmd.AddCommand(UserCmd)
}
