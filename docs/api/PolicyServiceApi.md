# \PolicyServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPolicies**](PolicyServiceApi.md#ListPolicies) | **Post** /policy | List all defined security policies


# **ListPolicies**
> IdmListPolicyGroupsResponse ListPolicies(ctx, body)
List all defined security policies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**IdmListPolicyGroupsRequest**](IdmListPolicyGroupsRequest.md)|  | 

### Return type

[**IdmListPolicyGroupsResponse**](idmListPolicyGroupsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

