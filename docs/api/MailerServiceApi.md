# \MailerServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Send**](MailerServiceApi.md#Send) | **Post** /mailer/send | Send an email to a user or any email address


# **Send**
> MailerSendMailResponse Send(ctx, body)
Send an email to a user or any email address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**MailerMail**](MailerMail.md)|  | 

### Return type

[**MailerSendMailResponse**](mailerSendMailResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

