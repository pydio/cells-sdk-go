# \RoleServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRole**](RoleServiceApi.md#DeleteRole) | **Delete** /role/{Uuid} | Delete a Role by ID
[**GetRole**](RoleServiceApi.md#GetRole) | **Get** /role/{Uuid} | Get a Role by ID
[**SearchRoles**](RoleServiceApi.md#SearchRoles) | **Post** /role | Search Roles
[**SetRole**](RoleServiceApi.md#SetRole) | **Put** /role/{Uuid} | Create or update a Role


# **DeleteRole**
> IdmRole DeleteRole(ctx, uuid)
Delete a Role by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uuid** | **string**|  | 

### Return type

[**IdmRole**](IdmRole.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **GetRole**
> IdmRole GetRole(ctx, uuid, optional)
Get a Role by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uuid** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string**|  | 
 **label** | **string**|  | 
 **isTeam** | **bool**|  | 
 **groupRole** | **bool**|  | 
 **userRole** | **bool**|  | 
 **lastUpdated** | **int32**|  | 
 **autoApplies** | [**[]string**](string.md)|  | 
 **policiesContextEditable** | **bool**|  | 

### Return type

[**IdmRole**](IdmRole.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **SearchRoles**
> RestRolesCollection SearchRoles(ctx, body)
Search Roles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**RestSearchRoleRequest**](RestSearchRoleRequest.md)|  | 

### Return type

[**RestRolesCollection**](restRolesCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **SetRole**
> IdmRole SetRole(ctx, uuid, body)
Create or update a Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uuid** | **string**|  | 
  **body** | [**IdmRole**](IdmRole.md)|  | 

### Return type

[**IdmRole**](IdmRole.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

