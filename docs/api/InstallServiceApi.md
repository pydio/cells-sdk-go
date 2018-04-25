# \InstallServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgreement**](InstallServiceApi.md#GetAgreement) | **Get** /install/agreement | Perform a check during install (like DB connection, php-fpm detection, etc)
[**GetInstall**](InstallServiceApi.md#GetInstall) | **Get** /install | Loads default values for install form
[**PerformInstallCheck**](InstallServiceApi.md#PerformInstallCheck) | **Post** /install/check | Perform a check during install (like DB connection, php-fpm detection, etc)
[**PostInstall**](InstallServiceApi.md#PostInstall) | **Post** /install | Post values to be saved for install


# **GetAgreement**
> InstallGetAgreementResponse GetAgreement(ctx, )
Perform a check during install (like DB connection, php-fpm detection, etc)

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InstallGetAgreementResponse**](installGetAgreementResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **GetInstall**
> InstallGetDefaultsResponse GetInstall(ctx, )
Loads default values for install form

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InstallGetDefaultsResponse**](installGetDefaultsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **PerformInstallCheck**
> InstallPerformCheckResponse PerformInstallCheck(ctx, body)
Perform a check during install (like DB connection, php-fpm detection, etc)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**InstallPerformCheckRequest**](InstallPerformCheckRequest.md)|  | 

### Return type

[**InstallPerformCheckResponse**](installPerformCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **PostInstall**
> InstallInstallResponse PostInstall(ctx, body)
Post values to be saved for install

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**InstallInstallRequest**](InstallInstallRequest.md)|  | 

### Return type

[**InstallInstallResponse**](installInstallResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

