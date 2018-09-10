package tree

import (
	"github.com/pydio/cells-sdk-go/client/workspace_service"
	"github.com/pydio/cells-sdk-go/config"
	"github.com/pydio/cells-sdk-go/models"
)

// CreateWorkspace creates a new simple workspace using the current default connection.
func CreateWorkspace(datasource, slug, label, description string) (*models.IdmWorkspace, error) {

	apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
	if err != nil {
		return nil, err
	}

	iw := &models.IdmWorkspace{}
	iw.Label = label
	iw.Slug = slug
	iw.Description = description
	// TODO this should also be configurable
	iw.Scope = "ADMIN"
	iw.Attributes = "{\"DEFAULT_RIGHTS\":\"rw\"}"

	root := models.TreeNode{
		UUID:      "DATASOURCE:" + datasource,
		Path:      datasource,
		MetaStore: map[string]string{"name": "\"" + datasource + "\""},
	}

	iw.RootNodes = models.IdmWorkspaceRootNodes{
		"DATASOURCE:" + datasource: root,
	}

	params := workspace_service.NewPutWorkspaceParamsWithContext(ctx)
	params.Slug = slug
	params.Body = iw

	result, err := apiClient.WorkspaceService.PutWorkspace(params)
	if err != nil {
		return nil, err
	}

	return result.Payload, nil
}

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

	// for i, ws := range result.Payload.Workspaces {
	// 	fmt.Printf("#%d - %s\n", i, ws.Label)
	// }

	return result.Payload, nil
}

// DeleteWorkspace deregister a workspace.
func DeleteWorkspace(slug string) (bool, error) {

	apiClient, ctx, err := config.GetPreparedApiClient(config.DefaultConfig)
	if err != nil {
		return false, err
	}

	params := workspace_service.NewDeleteWorkspaceParamsWithContext(ctx)
	params.Slug = slug

	result, err := apiClient.WorkspaceService.DeleteWorkspace(params)
	if err != nil {
		return false, err
	}

	return result.Payload.Success, nil
}
