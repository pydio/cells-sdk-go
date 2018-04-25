# \AdminTreeServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAdminTree**](AdminTreeServiceApi.md#ListAdminTree) | **Post** /tree/admin/list | List files and folders starting at the root (first level lists the datasources)
[**StatAdminTree**](AdminTreeServiceApi.md#StatAdminTree) | **Post** /tree/admin/stat | Read a node information inside the admin tree


# **ListAdminTree**
> RestNodesCollection ListAdminTree(ctx, body)
List files and folders starting at the root (first level lists the datasources)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**TreeListNodesRequest**](TreeListNodesRequest.md)|  | 

### Return type

[**RestNodesCollection**](restNodesCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **StatAdminTree**
> TreeReadNodeResponse StatAdminTree(ctx, body)
Read a node information inside the admin tree

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**TreeReadNodeRequest**](TreeReadNodeRequest.md)|  | 

### Return type

[**TreeReadNodeResponse**](treeReadNodeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

