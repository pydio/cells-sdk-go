# \LogServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Syslog**](LogServiceApi.md#Syslog) | **Post** /log/sys | Technical Logs, in Json or CSV format


# **Syslog**
> RestLogMessageCollection Syslog(ctx, body)
Technical Logs, in Json or CSV format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**LogListLogRequest**](LogListLogRequest.md)|  | 

### Return type

[**RestLogMessageCollection**](restLogMessageCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

