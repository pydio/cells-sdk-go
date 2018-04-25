# \UserMetaServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListUserMetaNamespace**](UserMetaServiceApi.md#ListUserMetaNamespace) | **Get** /user-meta/namespace | List defined meta namespaces
[**SearchUserMeta**](UserMetaServiceApi.md#SearchUserMeta) | **Post** /user-meta/search | Search a list of meta by node Id or by User id and by namespace
[**UpdateUserMeta**](UserMetaServiceApi.md#UpdateUserMeta) | **Put** /user-meta/update | Update/delete user meta
[**UpdateUserMetaNamespace**](UserMetaServiceApi.md#UpdateUserMetaNamespace) | **Put** /user-meta/namespace | Admin: update namespaces
[**UserBookmarks**](UserMetaServiceApi.md#UserBookmarks) | **Post** /user-meta/bookmarks | Special API for Bookmarks, will load userMeta and the associated nodes, and return as a node list


# **ListUserMetaNamespace**
> RestUserMetaNamespaceCollection ListUserMetaNamespace(ctx, )
List defined meta namespaces

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RestUserMetaNamespaceCollection**](restUserMetaNamespaceCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **SearchUserMeta**
> RestUserMetaCollection SearchUserMeta(ctx, body)
Search a list of meta by node Id or by User id and by namespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**IdmSearchUserMetaRequest**](IdmSearchUserMetaRequest.md)|  | 

### Return type

[**RestUserMetaCollection**](restUserMetaCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **UpdateUserMeta**
> IdmUpdateUserMetaResponse UpdateUserMeta(ctx, body)
Update/delete user meta

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**IdmUpdateUserMetaRequest**](IdmUpdateUserMetaRequest.md)|  | 

### Return type

[**IdmUpdateUserMetaResponse**](idmUpdateUserMetaResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **UpdateUserMetaNamespace**
> IdmUpdateUserMetaNamespaceResponse UpdateUserMetaNamespace(ctx, body)
Admin: update namespaces

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**IdmUpdateUserMetaNamespaceRequest**](IdmUpdateUserMetaNamespaceRequest.md)|  | 

### Return type

[**IdmUpdateUserMetaNamespaceResponse**](idmUpdateUserMetaNamespaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **UserBookmarks**
> RestBulkMetaResponse UserBookmarks(ctx, body)
Special API for Bookmarks, will load userMeta and the associated nodes, and return as a node list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**RestUserBookmarksRequest**](RestUserBookmarksRequest.md)|  | 

### Return type

[**RestBulkMetaResponse**](restBulkMetaResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

