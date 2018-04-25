# \WorkspaceServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWorkspace**](WorkspaceServiceApi.md#DeleteWorkspace) | **Delete** /workspace/{Slug} | Delete an existing workspace
[**PutWorkspace**](WorkspaceServiceApi.md#PutWorkspace) | **Put** /workspace/{Slug} | Create or update a workspace
[**SearchWorkspaces**](WorkspaceServiceApi.md#SearchWorkspaces) | **Post** /workspace | Search workspaces on certain keys


# **DeleteWorkspace**
> RestDeleteResponse DeleteWorkspace(ctx, slug)
Delete an existing workspace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **slug** | **string**|  | 

### Return type

[**RestDeleteResponse**](restDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **PutWorkspace**
> IdmWorkspace PutWorkspace(ctx, slug, body)
Create or update a workspace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **slug** | **string**|  | 
  **body** | [**IdmWorkspace**](IdmWorkspace.md)|  | 

### Return type

[**IdmWorkspace**](idmWorkspace.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **SearchWorkspaces**
> RestWorkspaceCollection SearchWorkspaces(ctx, body)
Search workspaces on certain keys

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**RestSearchWorkspaceRequest**](RestSearchWorkspaceRequest.md)|  | 

### Return type

[**RestWorkspaceCollection**](restWorkspaceCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

