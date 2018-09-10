package cmd

import (
	"fmt"
	"time"
	"log"


	"github.com/spf13/cobra"

	"github.com/pydio/cells-sdk-go/client/meta_service"
	"github.com/pydio/cells-sdk-go/client/user_service"
	"github.com/pydio/cells-sdk-go/config"
	"github.com/pydio/cells-sdk-go/models"
)

var (
	knownPwd = map[string]string{
		"bob":   "P@ssw0rd",
		"alice": "P@ssw0rd",
		"bella": "hadid",
		"emma":  "watson",
		"patrick": "jane",
	}
)

var listUserFiles = &cobra.Command{
	Use:   "list-user-files",
	Short: "Ping demo server",
	Long:  `Send a listUsers request then tries to list workspaces for each user`,
	Run: func(cmd *cobra.Command, args []string) {

		//connect to the api
		httpClient := config.GetHttpClient(config.DefaultConfig)
		apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
		if err != nil {
			log.Fatal(err)
		}

		//retrieve users
		param := &user_service.SearchUsersParams{
			Context:    ctx,
			HTTPClient: httpClient,
		}

		result, err := apiClient.UserService.SearchUsers(param)
		if err != nil {
			fmt.Printf("could not list users: %s\n", err.Error())
			log.Fatal(err)
		}

		//list users
		var foundOne bool
		fmt.Printf("Found %d users on this Cells\n", len(result.Payload.Users))
		if len(result.Payload.Users) > 0 {
			for i, u := range result.Payload.Users {
				fmt.Println(i+1, " *********  ", u.Login)
			}
		}
		if len(result.Payload.Users) > 0 {
			for _, u := range result.Payload.Users {
				fmt.Println(" ----------------", u.Login, "----------------")
				if e := listingUserFiles(u.Login, knownPwd); e == nil {
					foundOne = true
				}
				time.Sleep(100 * time.Millisecond)
			}
		}
		if !foundOne {
			log.Fatal("Could not find any workspaces for any users, something strange happened!")
		}
	},
}

func listingUserFiles(login string, knownPasswords map[string]string) error {

	var userPass string

	if value, ok := knownPasswords[login]; ok {
		userPass = value
	} else {
		userPass = login
	}

	uSdkConfig := &config.SdkConfig{
		Protocol:     protocol,
		Url:          host,
		ClientKey:    id,
		ClientSecret: secret,
		User:         login,
		Password:     userPass,
		SkipVerify:   skipVerify,
	}

	config.DefaultConfig = uSdkConfig
	uHttpClient := config.GetHttpClient(uSdkConfig)
	uApiClient, ctx, err := config.GetPreparedApiClient(uSdkConfig)

	if err != nil {
		log.Println("could not log with this user :", login)
		return fmt.Errorf("could not log in, not able to fetch the password for %s %s", login, err.Error())
	} else {
		log.Println("Successfully logged ", login)
	}

	params := &meta_service.GetBulkMetaParams{
		Body: &models.RestGetBulkMetaRequest{NodePaths: []string{
			"/personal-files/*",
			//"/*",
		}},
		Context:    ctx,
		HTTPClient: uHttpClient,
	}

	result, err := uApiClient.MetaService.GetBulkMeta(params)
	if err != nil {
		return fmt.Errorf("could not list meta %s", err.Error())
	}

	if len(result.Payload.Nodes) > 0 {
		fmt.Printf("* %d meta\n", len(result.Payload.Nodes))
		fmt.Println("USER ", login)

		for _, u := range result.Payload.Nodes {
			fmt.Println("  -", u.Path)

		}

	}

	return nil
}

func init() {
	ExampleCmd.AddCommand(listUserFiles)
}
