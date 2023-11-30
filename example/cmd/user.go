package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/pydio/cells-sdk-go/v4/client/user_service"
	"github.com/pydio/cells-sdk-go/v4/models"
	"github.com/pydio/cells-sdk-go/v4/transport/rest"
)

var (
	newUserName     string
	newUserPassword string
	newUserIsAdmin  bool

	readP  = models.ServiceResourcePolicyActionREAD
	writeP = models.ServiceResourcePolicyActionWRITE
	allowP = models.ServiceResourcePolicyPolicyEffectAllow

	DefaultUserPolicies = []*models.ServiceResourcePolicy{
		{
			Action:  models.NewServiceResourcePolicyAction(readP),
			Effect:  models.NewServiceResourcePolicyPolicyEffect(allowP),
			Subject: "profile:" + ProfileStandard,
		},
		{
			Action:  models.NewServiceResourcePolicyAction(writeP),
			Effect:  models.NewServiceResourcePolicyPolicyEffect(allowP),
			Subject: "profile:" + ProfileAdmin,
		},
	}

	ProfileAdmin    = "admin"
	ProfileStandard = "standard"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Set of sample commands to manage user via the API",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var addUserCmd = &cobra.Command{
	Use: "add",
	Long: `Test simple user creation.
	
	Example:
	$ go run main.go user -u https://localhost:8080 -l admin -p admin --skipVerify=true add --newUser test --newPwd='P@ssw0rd'
	`,
	Run: func(cmd *cobra.Command, args []string) {

		// Connect to the Pydio API via the sdkConfig
		ctx, apiClient, err := rest.GetClient(DefaultConfig, false)
		if err != nil {
			log.Fatal(err)
		}

		policies := append(DefaultUserPolicies, &models.ServiceResourcePolicy{
			Action:  models.NewServiceResourcePolicyAction(writeP),
			Effect:  models.NewServiceResourcePolicyPolicyEffect(allowP),
			Subject: "user:" + newUserName,
		})

		profile := ProfileStandard
		if newUserIsAdmin {
			profile = ProfileAdmin
		}

		// Create User
		newUser := user_service.PutUserBody{
			GroupPath:  "/",
			Password:   newUserPassword,
			Policies:   policies,
			Attributes: map[string]string{"profile": profile},
		}

		_, err = apiClient.UserService.PutUser(&user_service.PutUserParams{
			Login:   newUserName,
			Body:    newUser,
			Context: ctx,
		})
		if err != nil {
			log.Fatal("could not create user ", newUserName, ", aborting...\nCause:", err)
		}
		// fmt.Println(" ### user put...")
	},
}

func init() {
	addUserCmd.Flags().StringVar(&newUserName, "newUser", "", "Login of the new user")
	addUserCmd.Flags().StringVar(&newUserPassword, "newPwd", "", "Password for the new user")
	addUserCmd.Flags().BoolVar(&newUserIsAdmin, "newIsAdmin", false, "Flag to create an admin user")

	userCmd.AddCommand(addUserCmd)
	ExampleCmd.AddCommand(userCmd)
}
