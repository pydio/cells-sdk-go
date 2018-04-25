# \GraphServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Relation**](GraphServiceApi.md#Relation) | **Get** /graph/relation/{UserId} | Compute relation of context user with another user
[**UserState**](GraphServiceApi.md#UserState) | **Get** /graph/state/{Segment} | Compute accessible workspaces for a given user


# **Relation**
> RestRelationResponse Relation(ctx, userId)
Compute relation of context user with another user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **userId** | **string**|  | 

### Return type

[**RestRelationResponse**](restRelationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **UserState**
> RestUserStateResponse UserState(ctx, segment)
Compute accessible workspaces for a given user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **segment** | **string**|  | 

### Return type

[**RestUserStateResponse**](restUserStateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

