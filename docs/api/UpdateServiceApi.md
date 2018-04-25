# \UpdateServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyUpdate**](UpdateServiceApi.md#ApplyUpdate) | **Get** /update/{TargetVersion} | Apply an update to a given version
[**UpdateRequired**](UpdateServiceApi.md#UpdateRequired) | **Get** /update | Check the remote server to see if there are available binaries


# **ApplyUpdate**
> UpdateApplyUpdateResponse ApplyUpdate(ctx, targetVersion)
Apply an update to a given version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **targetVersion** | **string**|  | 

### Return type

[**UpdateApplyUpdateResponse**](updateApplyUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **UpdateRequired**
> UpdateUpdateResponse UpdateRequired(ctx, optional)
Check the remote server to see if there are available binaries

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **string**| Channel name. | 
 **packageName** | **string**| Name of the currently running application. | 
 **currentVersion** | **string**| Current version of the application. | 
 **gOOS** | **string**| Current GOOS. | 
 **gOARCH** | **string**| Current GOARCH. | 
 **serviceName** | **string**| Not Used : specific service to get updates for. | 

### Return type

[**UpdateUpdateResponse**](updateUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

