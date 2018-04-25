# \EnterpriseLogServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Audit**](EnterpriseLogServiceApi.md#Audit) | **Post** /log/audit | Auditable Logs, in Json or CSV format
[**AuditChartData**](EnterpriseLogServiceApi.md#AuditChartData) | **Post** /log/audit/chartdata | Retrieves aggregated audit logs to generate charts
[**AuditExport**](EnterpriseLogServiceApi.md#AuditExport) | **Post** /log/audit/export | Auditable Logs, in Json or CSV format
[**SyslogExport**](EnterpriseLogServiceApi.md#SyslogExport) | **Post** /log/sys/export | Technical Logs, in Json or CSV format


# **Audit**
> RestLogMessageCollection Audit(ctx, body)
Auditable Logs, in Json or CSV format

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

# **AuditChartData**
> RestTimeRangeResultCollection AuditChartData(ctx, body)
Retrieves aggregated audit logs to generate charts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**LogTimeRangeRequest**](LogTimeRangeRequest.md)|  | 

### Return type

[**RestTimeRangeResultCollection**](restTimeRangeResultCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **AuditExport**
> RestLogMessageCollection AuditExport(ctx, body)
Auditable Logs, in Json or CSV format

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

# **SyslogExport**
> RestLogMessageCollection SyslogExport(ctx, body)
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

