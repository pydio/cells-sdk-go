# \DocStoreServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDoc**](DocStoreServiceApi.md#DeleteDoc) | **Post** /docstore/bulk_delete/{StoreID} | Delete one or more docs inside a given store
[**GetDoc**](DocStoreServiceApi.md#GetDoc) | **Get** /docstore/{StoreID}/{DocumentID} | Load one document by ID from a given store
[**ListDocs**](DocStoreServiceApi.md#ListDocs) | **Post** /docstore/{StoreID} | List all docs of a given store
[**PutDoc**](DocStoreServiceApi.md#PutDoc) | **Put** /docstore/{StoreID}/{DocumentID} | Put a document inside a given store


# **DeleteDoc**
> DocstoreDeleteDocumentsResponse DeleteDoc(ctx, storeID, body)
Delete one or more docs inside a given store

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **storeID** | **string**|  | 
  **body** | [**DocstoreDeleteDocumentsRequest**](DocstoreDeleteDocumentsRequest.md)|  | 

### Return type

[**DocstoreDeleteDocumentsResponse**](docstoreDeleteDocumentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **GetDoc**
> DocstoreGetDocumentResponse GetDoc(ctx, storeID, documentID)
Load one document by ID from a given store

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **storeID** | **string**|  | 
  **documentID** | **string**|  | 

### Return type

[**DocstoreGetDocumentResponse**](docstoreGetDocumentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **ListDocs**
> RestDocstoreCollection ListDocs(ctx, storeID, body)
List all docs of a given store

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **storeID** | **string**|  | 
  **body** | [**RestListDocstoreRequest**](RestListDocstoreRequest.md)|  | 

### Return type

[**RestDocstoreCollection**](restDocstoreCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **PutDoc**
> DocstorePutDocumentResponse PutDoc(ctx, storeID, documentID, body)
Put a document inside a given store

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **storeID** | **string**|  | 
  **documentID** | **string**|  | 
  **body** | [**DocstorePutDocumentRequest**](DocstorePutDocumentRequest.md)|  | 

### Return type

[**DocstorePutDocumentResponse**](docstorePutDocumentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

