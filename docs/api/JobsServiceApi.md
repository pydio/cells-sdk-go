# \JobsServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserControlJob**](JobsServiceApi.md#UserControlJob) | **Put** /jobs/user | Send Control Commands to one or many jobs / tasks
[**UserCreateJob**](JobsServiceApi.md#UserCreateJob) | **Put** /jobs/user/{JobName} | Create a predefined job to be run directly
[**UserDeleteTasks**](JobsServiceApi.md#UserDeleteTasks) | **Post** /jobs/tasks/delete | Send a control command to clean tasks on a given job
[**UserListJobs**](JobsServiceApi.md#UserListJobs) | **Post** /jobs/user | List jobs associated with current user


# **UserControlJob**
> JobsCtrlCommandResponse UserControlJob(ctx, body)
Send Control Commands to one or many jobs / tasks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**JobsCtrlCommand**](JobsCtrlCommand.md)|  | 

### Return type

[**JobsCtrlCommandResponse**](jobsCtrlCommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **UserCreateJob**
> RestUserJobResponse UserCreateJob(ctx, jobName, body)
Create a predefined job to be run directly

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobName** | **string**|  | 
  **body** | [**RestUserJobRequest**](RestUserJobRequest.md)|  | 

### Return type

[**RestUserJobResponse**](restUserJobResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **UserDeleteTasks**
> JobsDeleteTasksResponse UserDeleteTasks(ctx, body)
Send a control command to clean tasks on a given job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**JobsDeleteTasksRequest**](JobsDeleteTasksRequest.md)|  | 

### Return type

[**JobsDeleteTasksResponse**](jobsDeleteTasksResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **UserListJobs**
> RestUserJobsCollection UserListJobs(ctx, body)
List jobs associated with current user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**JobsListJobsRequest**](JobsListJobsRequest.md)|  | 

### Return type

[**RestUserJobsCollection**](restUserJobsCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

