# \ACLServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAcl**](ACLServiceApi.md#DeleteAcl) | **Post** /acl/bulk/delete | Delete one or more ACLs
[**PutAcl**](ACLServiceApi.md#PutAcl) | **Put** /acl | Store an ACL
[**SearchAcls**](ACLServiceApi.md#SearchAcls) | **Post** /acl | Search Acls


# **DeleteAcl**
> RestDeleteResponse DeleteAcl(ctx, body)
Delete one or more ACLs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**IdmAcl**](IdmAcl.md)|  | 

### Return type

[**RestDeleteResponse**](restDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **PutAcl**
> IdmAcl PutAcl(ctx, body)
Store an ACL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**IdmAcl**](IdmAcl.md)|  | 

### Return type

[**IdmAcl**](idmACL.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **SearchAcls**
> RestAclCollection SearchAcls(ctx, body)
Search Acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**RestSearchAclRequest**](RestSearchAclRequest.md)|  | 

### Return type

[**RestAclCollection**](restACLCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

