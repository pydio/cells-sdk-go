# \EnterpriseConfigServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteExternalDirectory**](EnterpriseConfigServiceApi.md#DeleteExternalDirectory) | **Delete** /config/directories/{ConfigId} | [Enterprise Only] Delete external directory
[**DeleteVersioningPolicy**](EnterpriseConfigServiceApi.md#DeleteVersioningPolicy) | **Delete** /config/versioning/{Uuid} | [Enterprise Only] Delete a versioning policy
[**ListExternalDirectories**](EnterpriseConfigServiceApi.md#ListExternalDirectories) | **Get** /config/directories | [Enterprise Only] List additional user directories
[**PutExternalDirectory**](EnterpriseConfigServiceApi.md#PutExternalDirectory) | **Put** /config/directories/{ConfigId} | [Enterprise Only] Add/Create an external directory
[**PutVersioningPolicy**](EnterpriseConfigServiceApi.md#PutVersioningPolicy) | **Post** /config/versioning/{Uuid} | [Enterprise Only] Create or update a versioning policy


# **DeleteExternalDirectory**
> RestExternalDirectoryResponse DeleteExternalDirectory(ctx, configId)
[Enterprise Only] Delete external directory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **configId** | **string**|  | 

### Return type

[**RestExternalDirectoryResponse**](restExternalDirectoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **DeleteVersioningPolicy**
> RestDeleteVersioningPolicyResponse DeleteVersioningPolicy(ctx, uuid)
[Enterprise Only] Delete a versioning policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uuid** | **string**|  | 

### Return type

[**RestDeleteVersioningPolicyResponse**](restDeleteVersioningPolicyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **ListExternalDirectories**
> RestExternalDirectoryCollection ListExternalDirectories(ctx, )
[Enterprise Only] List additional user directories

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RestExternalDirectoryCollection**](restExternalDirectoryCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **PutExternalDirectory**
> RestExternalDirectoryResponse PutExternalDirectory(ctx, configId, body)
[Enterprise Only] Add/Create an external directory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **configId** | **string**|  | 
  **body** | [**RestExternalDirectoryConfig**](RestExternalDirectoryConfig.md)|  | 

### Return type

[**RestExternalDirectoryResponse**](restExternalDirectoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **PutVersioningPolicy**
> TreeVersioningPolicy PutVersioningPolicy(ctx, uuid, body)
[Enterprise Only] Create or update a versioning policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uuid** | **string**|  | 
  **body** | [**TreeVersioningPolicy**](TreeVersioningPolicy.md)|  | 

### Return type

[**TreeVersioningPolicy**](treeVersioningPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

