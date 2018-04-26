# \UserServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BindUsers**](UserServiceApi.md#BindUsers) | **Post** /user/{Login}/bind | Bind a user with her login and password
[**DeleteUser**](UserServiceApi.md#DeleteUser) | **Delete** /user/{Login} | Delete a user
[**GetUser**](UserServiceApi.md#GetUser) | **Get** /user/{Login} | Get a user by login
[**PutRoles**](UserServiceApi.md#PutRoles) | **Put** /user/roles/{Login} | Just save a user roles, without other datas
[**PutUser**](UserServiceApi.md#PutUser) | **Put** /user/{Login} | Create or update a user
[**SearchUsers**](UserServiceApi.md#SearchUsers) | **Post** /user | List/Search users


# **BindUsers**
> RestBindResponse BindUsers(ctx, login, body)
Bind a user with her login and password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **login** | **string**|  | 
  **body** | [**IdmUser**](IdmUser.md)|  | 

### Return type

[**RestBindResponse**](restBindResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **DeleteUser**
> RestDeleteResponse DeleteUser(ctx, login)
Delete a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **login** | **string**|  | 

### Return type

[**RestDeleteResponse**](restDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **GetUser**
> IdmUser GetUser(ctx, login, optional)
Get a user by login

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **login** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **login** | **string**|  | 
 **uuid** | **string**|  | 
 **groupPath** | **string**|  | 
 **password** | **string**|  | 
 **isGroup** | **bool**| Group specific data. | 
 **groupLabel** | **string**|  | 
 **policiesContextEditable** | **bool**|  | 

### Return type

[**IdmUser**](IdmUser.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **PutRoles**
> IdmUser PutRoles(ctx, login, body)
Just save a user roles, without other datas

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **login** | **string**|  | 
  **body** | [**IdmUser**](IdmUser.md)|  | 

### Return type

[**IdmUser**](IdmUser.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **PutUser**
> IdmUser PutUser(ctx, login, body)
Create or update a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **login** | **string**|  | 
  **body** | [**IdmUser**](IdmUser.md)|  | 

### Return type

[**IdmUser**](IdmUser.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **SearchUsers**
> RestUsersCollection SearchUsers(ctx, body)
List/Search users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**RestSearchUserRequest**](RestSearchUserRequest.md)|  | 

### Return type

[**RestUsersCollection**](restUsersCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

