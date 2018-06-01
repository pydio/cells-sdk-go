package idm

import (
	"github.com/pydio/cells-sdk-go/client/user_service"
	"github.com/pydio/cells-sdk-go/config"
	"github.com/pydio/cells-sdk-go/models"
	"github.com/pydio/cells/common"
)

func CreateUser(groupPath, createUserLogin, createUserPwd string, createUserAdmin bool) (models.IdmUser, error) {

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
		return models.IdmUser{}, err
	}

	return *result.Payload, nil
}

// ListUsers returns an array containing all known users
func ListUsers() (*models.RestUsersCollection, error) {

	apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
	if err != nil {
		return nil, err
	}

	result, err := apiClient.UserService.SearchUsers(user_service.NewSearchUsersParamsWithContext(ctx))
	if err != nil {
		return nil, err
	}

	return result.Payload, nil
}
