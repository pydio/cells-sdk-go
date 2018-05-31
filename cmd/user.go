package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/pydio/cells-sdk-go/client/role_service"
	"github.com/pydio/cells-sdk-go/client/user_service"
	"github.com/pydio/cells-sdk-go/config"
	"github.com/pydio/cells-sdk-go/models"
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

		log.Fatal("not implemented")
		// apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// query := api.RestSearchUserRequest{}

		// users, _, err := apiClient.UserServiceApi.SearchUsers(
		// 	ctx,
		// 	query,
		// )
		// if err != nil {
		// 	fmt.Printf("could not list users: %s\n", err.Error())
		// 	log.Fatal(err)
		// }

		// fmt.Printf("Found %d results\n", users.Total)
		// if len(users.Groups) > 0 {
		// 	fmt.Printf("* %d groups\n", len(users.Groups))
		// 	for _, u := range users.Groups {
		// 		fmt.Println("  - " + u.GroupLabel)
		// 	}
		// }

		// if len(users.Users) > 0 {
		// 	fmt.Printf("* %d users\n", len(users.Users))
		// 	for _, u := range users.Users {
		// 		fmt.Println("  - " + u.Login)
		// 	}
		// }

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

		r := models.ServiceResourcePolicyActionREAD
		w := models.ServiceResourcePolicyActionWRITE
		allow := models.ServiceResourcePolicyPolicyEffectAllow
		policies := []*models.ServiceResourcePolicy{
			&models.ServiceResourcePolicy{Action: r, Effect: allow, Subject: "profile:" + common.PYDIO_PROFILE_STANDARD},
			&models.ServiceResourcePolicy{Action: w, Effect: allow, Subject: "user:" + createUserLogin},
			&models.ServiceResourcePolicy{Action: w, Effect: allow, Subject: "profile:" + common.PYDIO_PROFILE_ADMIN},
		}

		profile := common.PYDIO_PROFILE_STANDARD
		if createUserAdmin {
			profile = common.PYDIO_PROFILE_ADMIN
		}
		// Create User
		newUser := models.IdmUser{
			GroupPath:  "/",
			Login:      createUserLogin,
			Password:   createUserPwd,
			Policies:   policies,
			Attributes: map[string]string{"profile": profile},
		}

		result, err := apiClient.UserService.PutUser(&user_service.PutUserParams{
			Login:   createUserLogin,
			Body:    &newUser,
			Context: ctx,
		})
		if err != nil {
			fmt.Printf("could not create user %s: %s\n", createUserLogin, err.Error())
			log.Fatal(err)
		}

		// Create User Role with correct policies
		newRole := models.IdmRole{
			UUID:     result.Payload.UUID,
			Policies: policies,
			UserRole: true,
			Label:    "User " + createUserLogin + " role",
		}
		_, err = apiClient.RoleService.SetRole(&role_service.SetRoleParams{
			UUID:    result.Payload.UUID,
			Body:    &newRole,
			Context: ctx,
		})
		if err != nil {
			log.Fatalf("could not create role for user %s: %s\n", createUserLogin, err.Error())
		}

		fmt.Printf("Created user with login: %s and uuid: %s\n", result.Payload.Login, result.Payload.UUID)

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
