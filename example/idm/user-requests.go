package idm

import (
	"fmt"

	sdkapi "github.com/pydio/cells-sdk-go/api"
	"github.com/pydio/cells-sdk-go/config"
)

// CreateUser simply creates a new user
func CreateUser(groupPath, login, password string) (sdkapi.IdmUser, error) {

	apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
	if err != nil {
		return sdkapi.IdmUser{}, err
	}

	newUser := sdkapi.IdmUser{
		GroupPath: groupPath,
		Login:     login,
		Password:  password,
	}

	user, _, err := apiClient.UserServiceApi.PutUser(
		ctx,
		login,
		newUser,
	)
	if err != nil {
		fmt.Printf("could not create user %s: %s\n", login, err.Error())
		return sdkapi.IdmUser{}, err
	}

	fmt.Printf("created user with login: %s and uuid: %s\n", user.Login, user.Uuid)

	return user, nil
}

// ListUsers returns an array containing all known users
func ListUsers() error {

	apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
	if err != nil {
		return err
	}

	query := sdkapi.RestSearchUserRequest{}

	users, _, err := apiClient.UserServiceApi.SearchUsers(
		ctx,
		query,
	)
	if err != nil {
		fmt.Printf("could not list users: %s\n", err.Error())
		return err
	}

	fmt.Printf("searched users, found %d\n", users.Total)

	return nil
}
