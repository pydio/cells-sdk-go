# \ChangeServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChanges**](ChangeServiceApi.md#GetChanges) | **Post** /changes/{SeqID} | Get Changes


# **GetChanges**
> RestChangeCollection GetChanges(ctx, seqID, body)
Get Changes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **seqID** | **string**|  | 
  **body** | [**RestChangeRequest**](RestChangeRequest.md)|  | 

### Return type

[**RestChangeCollection**](restChangeCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

