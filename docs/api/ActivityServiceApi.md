# \ActivityServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchSubscriptions**](ActivityServiceApi.md#SearchSubscriptions) | **Post** /activity/subscriptions | Load subscriptions to other users/nodes feeds
[**Stream**](ActivityServiceApi.md#Stream) | **Post** /activity/stream | Load the the feeds of the currently logged user
[**Subscribe**](ActivityServiceApi.md#Subscribe) | **Post** /activity/subscribe | Manage subscriptions to other users/nodes feeds


# **SearchSubscriptions**
> RestSubscriptionsCollection SearchSubscriptions(ctx, body)
Load subscriptions to other users/nodes feeds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ActivitySearchSubscriptionsRequest**](ActivitySearchSubscriptionsRequest.md)|  | 

### Return type

[**RestSubscriptionsCollection**](restSubscriptionsCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **Stream**
> ActivityObject Stream(ctx, body)
Load the the feeds of the currently logged user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ActivityStreamActivitiesRequest**](ActivityStreamActivitiesRequest.md)|  | 

### Return type

[**ActivityObject**](activityObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **Subscribe**
> ActivitySubscription Subscribe(ctx, body)
Manage subscriptions to other users/nodes feeds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ActivitySubscription**](ActivitySubscription.md)|  | 

### Return type

[**ActivitySubscription**](activitySubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

