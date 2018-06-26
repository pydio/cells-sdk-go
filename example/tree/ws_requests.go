package conf

import (
	"fmt"

	"github.com/pydio/cells-sdk-go/client/workspace_service"
	"github.com/pydio/cells-sdk-go/config"
	"github.com/pydio/cells-sdk-go/models"
)

// ListWorkspaces returns an object with known workspaces and their number.
func ListWorkspaces() (*models.RestWorkspaceCollection, error) {

	apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
	if err != nil {
		return nil, err
	}

	params := workspace_service.NewSearchWorkspacesParamsWithContext(ctx)
	queries := make([]*models.IdmWorkspaceSingleQuery, 1)
	queries[0] = &models.IdmWorkspaceSingleQuery{}
	params.Body = &models.RestSearchWorkspaceRequest{Queries: queries, Operation: "AND"}
	result, err := apiClient.WorkspaceService.SearchWorkspaces(params)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Found %d workspaces\n", result.Payload.Total)

	for i, ws := range result.Payload.Workspaces {
		fmt.Printf("#%d - %s\n", i, ws.Label)
	}

	return result.Payload, nil
}

// // AddLocalDatasource creates a new local file based datasource.
// func AddLocalDatasource(name, peerAddress string, port int32, rootFolderAbsPath string) (*models.ObjectDataSource, error) {

// 	apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ods := &models.ObjectDataSource{}
// 	ods.StorageType = "LOCAL"
// 	ods.Name = name
// 	ods.PeerAddress = peerAddress
// 	ods.ObjectsPort = port
// 	ods.StorageConfiguration = map[string]string{
// 		"folder":    rootFolderAbsPath,
// 		"normalize": "false",
// 	}

// 	params := config_service.NewPutDataSourceParamsWithContext(ctx)
// 	params.Name = name
// 	params.Body = ods

// 	result, err := apiClient.ConfigService.PutDataSource(params)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result.Payload, nil
// }

// // DeleteLocalDatasource removes a local file based datasource, without removing underlying files.
// func DeleteLocalDatasource(name string) (bool, error) {

// 	apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
// 	if err != nil {
// 		return false, err
// 	}

// 	params := config_service.NewDeleteDataSourceParamsWithContext(ctx)
// 	params.Name = name

// 	result, err := apiClient.ConfigService.DeleteDataSource(params)
// 	if err != nil {
// 		return false, err
// 	}

// 	return result.Payload.Success, nil
// }
