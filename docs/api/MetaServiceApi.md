# \MetaServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMeta**](MetaServiceApi.md#DeleteMeta) | **Post** /meta/delete/{NodePath} | Delete metadata of a given node
[**GetBulkMeta**](MetaServiceApi.md#GetBulkMeta) | **Post** /meta/bulk/get | List meta for a list of nodes, or a full directory using /path/_* syntax
[**GetMeta**](MetaServiceApi.md#GetMeta) | **Post** /meta/get/{NodePath} | Load metadata for a given node
[**SetMeta**](MetaServiceApi.md#SetMeta) | **Post** /meta/set/{NodePath} | Update metadata for a given node


# **DeleteMeta**
> TreeNode DeleteMeta(ctx, nodePath, body)
Delete metadata of a given node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nodePath** | **string**|  | 
  **body** | [**RestMetaNamespaceRequest**](RestMetaNamespaceRequest.md)|  | 

### Return type

[**TreeNode**](treeNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **GetBulkMeta**
> RestBulkMetaResponse GetBulkMeta(ctx, body)
List meta for a list of nodes, or a full directory using /path/_* syntax

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**RestGetBulkMetaRequest**](RestGetBulkMetaRequest.md)|  | 

### Return type

[**RestBulkMetaResponse**](restBulkMetaResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **GetMeta**
> TreeNode GetMeta(ctx, nodePath, body)
Load metadata for a given node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nodePath** | **string**|  | 
  **body** | [**RestMetaNamespaceRequest**](RestMetaNamespaceRequest.md)|  | 

### Return type

[**TreeNode**](treeNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **SetMeta**
> TreeNode SetMeta(ctx, nodePath, body)
Update metadata for a given node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nodePath** | **string**|  | 
  **body** | [**RestMetaCollection**](RestMetaCollection.md)|  | 

### Return type

[**TreeNode**](treeNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

