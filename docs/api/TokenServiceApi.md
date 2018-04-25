# \TokenServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResetPassword**](TokenServiceApi.md#ResetPassword) | **Post** /auth/reset-password | Finish up the reset password process by providing the unique token
[**ResetPasswordToken**](TokenServiceApi.md#ResetPasswordToken) | **Put** /auth/reset-password-token/{UserLogin} | Generate a unique token for the reset password process
[**Revoke**](TokenServiceApi.md#Revoke) | **Post** /auth/token/revoke | Revoke a JWT token


# **ResetPassword**
> RestResetPasswordResponse ResetPassword(ctx, body)
Finish up the reset password process by providing the unique token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**RestResetPasswordRequest**](RestResetPasswordRequest.md)|  | 

### Return type

[**RestResetPasswordResponse**](restResetPasswordResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **ResetPasswordToken**
> RestResetPasswordTokenResponse ResetPasswordToken(ctx, userLogin)
Generate a unique token for the reset password process

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **userLogin** | **string**|  | 

### Return type

[**RestResetPasswordTokenResponse**](restResetPasswordTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **Revoke**
> RestRevokeResponse Revoke(ctx, body)
Revoke a JWT token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**RestRevokeRequest**](RestRevokeRequest.md)|  | 

### Return type

[**RestRevokeResponse**](restRevokeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

