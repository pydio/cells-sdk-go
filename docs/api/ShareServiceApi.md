# \ShareServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCell**](ShareServiceApi.md#DeleteCell) | **Delete** /share/cell/{Uuid} | Delete a share room
[**DeleteShareLink**](ShareServiceApi.md#DeleteShareLink) | **Delete** /share/link/{Uuid} | Delete Share Link
[**GetCell**](ShareServiceApi.md#GetCell) | **Get** /share/cell/{Uuid} | Load a share room
[**GetShareLink**](ShareServiceApi.md#GetShareLink) | **Get** /share/link/{Uuid} | Load a share link with all infos
[**ListSharedResources**](ShareServiceApi.md#ListSharedResources) | **Post** /share/resources | List Shared Resources for current user or all users
[**PutCell**](ShareServiceApi.md#PutCell) | **Put** /share/cell | Put or Create a share room
[**PutShareLink**](ShareServiceApi.md#PutShareLink) | **Put** /share/link | Put or Create a share room


# **DeleteCell**
> RestDeleteCellResponse DeleteCell(ctx, uuid)
Delete a share room

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uuid** | **string**|  | 

### Return type

[**RestDeleteCellResponse**](restDeleteCellResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **DeleteShareLink**
> RestDeleteShareLinkResponse DeleteShareLink(ctx, uuid)
Delete Share Link

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uuid** | **string**|  | 

### Return type

[**RestDeleteShareLinkResponse**](restDeleteShareLinkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **GetCell**
> RestCell GetCell(ctx, uuid)
Load a share room

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uuid** | **string**|  | 

### Return type

[**RestCell**](restCell.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **GetShareLink**
> RestShareLink GetShareLink(ctx, uuid)
Load a share link with all infos

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uuid** | **string**|  | 

### Return type

[**RestShareLink**](restShareLink.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **ListSharedResources**
> RestListSharedResourcesResponse ListSharedResources(ctx, body)
List Shared Resources for current user or all users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**RestListSharedResourcesRequest**](RestListSharedResourcesRequest.md)|  | 

### Return type

[**RestListSharedResourcesResponse**](restListSharedResourcesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **PutCell**
> RestCell PutCell(ctx, body)
Put or Create a share room

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**RestPutCellRequest**](RestPutCellRequest.md)|  | 

### Return type

[**RestCell**](restCell.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **PutShareLink**
> RestShareLink PutShareLink(ctx, body)
Put or Create a share room

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**RestPutShareLinkRequest**](RestPutShareLinkRequest.md)|  | 

### Return type

[**RestShareLink**](restShareLink.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

