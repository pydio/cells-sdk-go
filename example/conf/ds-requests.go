package conf

import (
	"fmt"

	"github.com/pydio/cells-sdk-go/client/config_service"
	"github.com/pydio/cells-sdk-go/config"
	"github.com/pydio/cells-sdk-go/models"
)

// ListUsers returns an array containing all known users
func ListDatasources() (*models.RestDataSourceCollection, error) {

	fmt.Println("config url", config.DefaultConfig)
	fmt.Println("config url", config.DefaultConfig.Url)

	apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
	if err != nil {
		return nil, err
	}

	result, err := apiClient.ConfigService.ListDataSources(config_service.NewListDataSourcesParamsWithContext(ctx))
	if err != nil {
		return nil, err
	}

	return result.Payload, nil
}
