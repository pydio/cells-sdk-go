# \SearchServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Nodes**](SearchServiceApi.md#Nodes) | **Post** /search/nodes | Search indexed nodes (files/folders) on various aspects


# **Nodes**
> RestSearchResults Nodes(ctx, body)
Search indexed nodes (files/folders) on various aspects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**TreeSearchRequest**](TreeSearchRequest.md)|  | 

### Return type

[**RestSearchResults**](restSearchResults.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

