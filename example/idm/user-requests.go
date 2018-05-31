package idm

import (
	"fmt"

	"github.com/pydio/cells-sdk-go/client/user_service"
	"github.com/pydio/cells-sdk-go/config"
	"github.com/pydio/cells-sdk-go/models"
	"github.com/pydio/cells/common"
)

func CreateUser2(groupPath, createUserLogin, createUserPwd string, createUserAdmin bool) (models.IdmUser, error) {

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
		GroupPath:  groupPath,
		Login:      createUserLogin,
		Password:   createUserPwd,
		Policies:   policies,
		Attributes: map[string]string{"profile": profile},
	}

	apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
	if err != nil {
		return models.IdmUser{}, err
	}

	result, err := apiClient.UserService.PutUser(&user_service.PutUserParams{
		Login:   createUserLogin,
		Body:    &newUser,
		Context: ctx,
	})
	if err != nil {
		fmt.Printf("could not create user %s: %s\n", createUserLogin, err.Error())
		return models.IdmUser{}, err
	}

	// fmt.Printf("created user, result: %s\n", result.Payload.Login)
	fmt.Printf("created user with login: %s and uuid: %s\n", result.Payload.Login, result.Payload.UUID)

	return *result.Payload, nil
}

// CreateUser simply creates a new user
func CreateUser(groupPath, login, password string) (models.IdmUser, error) {

	apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
	if err != nil {
		return models.IdmUser{}, err
	}

	newUser := models.IdmUser{
		GroupPath: groupPath,
		Login:     login,
		Password:  password,
	}

	result, err := apiClient.UserService.PutUser(&user_service.PutUserParams{
		Login:   login,
		Body:    &newUser,
		Context: ctx,
	})
	if err != nil {
		fmt.Printf("could not create user %s: %s\n", login, err.Error())
		return models.IdmUser{}, err
	}

	// fmt.Printf("created user, result: %s\n", result.Payload.Login)
	fmt.Printf("created user with login: %s and uuid: %s\n", result.Payload.Login, result.Payload.UUID)

	return *result.Payload, nil
	// return models.IdmUser{}, fmt.Errorf("not yet implemented")
}

// ListUsers returns an array containing all known users
func ListUsers() error {

	// apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
	// if err != nil {
	// 	return err
	// }

	// query := sdkapi.RestSearchUserRequest{}

	// users, _, err := apiClient.UserServiceApi.SearchUsers(
	// 	ctx,
	// 	query,
	// )
	// if err != nil {
	// 	fmt.Printf("could not list users: %s\n", err.Error())
	// 	return err
	// }

	// fmt.Printf("searched users, found %d\n", users.Total)

	return nil
}
