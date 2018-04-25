# \FrontendServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FrontBootConf**](FrontendServiceApi.md#FrontBootConf) | **Get** /frontend/bootconf | Add some data to the initial set of parameters loaded by the frontend
[**FrontLog**](FrontendServiceApi.md#FrontLog) | **Put** /frontend/frontlogs | Sends a log from front (php) to back
[**SettingsMenu**](FrontendServiceApi.md#SettingsMenu) | **Get** /frontend/settings-menu | Sends a tree of nodes to be used a menu in the Settings panel


# **FrontBootConf**
> RestFrontBootConfResponse FrontBootConf(ctx, )
Add some data to the initial set of parameters loaded by the frontend

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RestFrontBootConfResponse**](restFrontBootConfResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **FrontLog**
> RestFrontLogResponse FrontLog(ctx, body)
Sends a log from front (php) to back

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**RestFrontLogMessage**](RestFrontLogMessage.md)|  | 

### Return type

[**RestFrontLogResponse**](restFrontLogResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **SettingsMenu**
> RestSettingsMenuResponse SettingsMenu(ctx, )
Sends a tree of nodes to be used a menu in the Settings panel

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RestSettingsMenuResponse**](restSettingsMenuResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

