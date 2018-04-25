# \LicenseServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LicenseStats**](LicenseServiceApi.md#LicenseStats) | **Get** /license/stats | [Enterprise Only] Display statistics about licenses usage


# **LicenseStats**
> CertLicenseStatsResponse LicenseStats(ctx, optional)
[Enterprise Only] Display statistics about licenses usage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forceRefresh** | **bool**|  | 

### Return type

[**CertLicenseStatsResponse**](certLicenseStatsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

