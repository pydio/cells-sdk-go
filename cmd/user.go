package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/pydio/cells-sdk-go/api"
	"github.com/pydio/cells-sdk-go/config"
	"github.com/pydio/cells/common"
)

var (
	createUserLogin string
	createUserPwd   string
	createUserAdmin bool
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

		r := api.SRPA_READ
		w := api.SRPA_WRITE
		allow := api.SRPPE_ALLOW
		policies := []api.ServiceResourcePolicy{
			{Action: &r, Effect: &allow, Subject: "profile:" + common.PYDIO_PROFILE_STANDARD},
			{Action: &w, Effect: &allow, Subject: "user:" + createUserLogin},
			{Action: &w, Effect: &allow, Subject: "profile:" + common.PYDIO_PROFILE_ADMIN},
		}

		profile := common.PYDIO_PROFILE_STANDARD
		if createUserAdmin {
			profile = common.PYDIO_PROFILE_ADMIN
		}
		// Create User
		newUser := api.IdmUser{
			GroupPath:  "/",
			Login:      createUserLogin,
			Password:   createUserPwd,
			Policies:   policies,
			Attributes: map[string]string{"profile": profile},
		}

		user, _, err := apiClient.UserServiceApi.PutUser(
			ctx,
			createUserLogin,
			newUser,
		)
		if err != nil {
			log.Fatalf("could not create user %s: %s\n", createUserLogin, err.Error())
		}

		// Create User Role with correct policies
		newRole := api.IdmRole{
			Uuid:     user.Uuid,
			Policies: policies,
			UserRole: true,
			Label:    "User " + createUserLogin + " role",
		}
		_, _, err = apiClient.RoleServiceApi.SetRole(ctx, user.Uuid, newRole)
		if err != nil {
			log.Fatalf("could not create role for user %s: %s\n", createUserLogin, err.Error())
		}

		fmt.Printf("Created user with login: %s and uuid: %s\n", user.Login, user.Uuid)

		ListUserCmd.Run(cmd, args)

	},
}

func init() {

	CreateUserCmd.Flags().StringVar(&createUserLogin, "create_login", "randomUser", "New user login")
	CreateUserCmd.Flags().StringVar(&createUserPwd, "create_password", "password", "New user password")
	CreateUserCmd.Flags().BoolVar(&createUserAdmin, "create_admin", false, "Set new user with admin capacities")

	UserCmd.AddCommand(ListUserCmd)
	UserCmd.AddCommand(CreateUserCmd)

	RootCmd.AddCommand(UserCmd)
}
