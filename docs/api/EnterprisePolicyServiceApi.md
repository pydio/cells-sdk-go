# \EnterprisePolicyServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePolicy**](EnterprisePolicyServiceApi.md#DeletePolicy) | **Delete** /policy/{Uuid} | Delete a security policy
[**PutPolicy**](EnterprisePolicyServiceApi.md#PutPolicy) | **Put** /policy | Update or create a security policy


# **DeletePolicy**
> RestDeleteResponse DeletePolicy(ctx, uuid)
Delete a security policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uuid** | **string**|  | 

### Return type

[**RestDeleteResponse**](restDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **PutPolicy**
> IdmPolicyGroup PutPolicy(ctx, body)
Update or create a security policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**IdmPolicyGroup**](IdmPolicyGroup.md)|  | 

### Return type

[**IdmPolicyGroup**](idmPolicyGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

