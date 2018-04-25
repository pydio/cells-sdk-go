# \ConfigServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigFormsDiscovery**](ConfigServiceApi.md#ConfigFormsDiscovery) | **Get** /config/discovery/forms/{ServiceName} | Publish Forms definition for building screens in frontend
[**ControlService**](ConfigServiceApi.md#ControlService) | **Post** /config/ctl | [Not Implemented]  Start/Stop a service
[**CreateEncryptionKey**](ConfigServiceApi.md#CreateEncryptionKey) | **Post** /config/encryption/create | Create a new master key
[**DeleteDataSource**](ConfigServiceApi.md#DeleteDataSource) | **Delete** /config/datasource/{Name} | Delete a datasource
[**DeleteEncryptionKey**](ConfigServiceApi.md#DeleteEncryptionKey) | **Post** /config/encryption/delete | Delete an existing master key
[**EndpointsDiscovery**](ConfigServiceApi.md#EndpointsDiscovery) | **Get** /config/discovery | Publish available endpoints
[**ExportEncryptionKey**](ConfigServiceApi.md#ExportEncryptionKey) | **Post** /config/encryption/export | Export a master key for backup purpose, protected with a password
[**GetConfig**](ConfigServiceApi.md#GetConfig) | **Get** /config/{FullPath} | Generic config Get using a full path in the config tree
[**GetDataSource**](ConfigServiceApi.md#GetDataSource) | **Get** /config/datasource/{Name} | Load datasource information
[**GetVersioningPolicy**](ConfigServiceApi.md#GetVersioningPolicy) | **Get** /config/versioning/{Uuid} | Load a given versioning policy
[**ImportEncryptionKey**](ConfigServiceApi.md#ImportEncryptionKey) | **Put** /config/encryption/import | Import a previously exported master key, requires the password created at export time
[**ListDataSources**](ConfigServiceApi.md#ListDataSources) | **Get** /config/datasource | List all defined datasources
[**ListEncryptionKeys**](ConfigServiceApi.md#ListEncryptionKeys) | **Post** /config/encryption/list | List registered master keys
[**ListPeerFolders**](ConfigServiceApi.md#ListPeerFolders) | **Post** /config/peers/{PeerAddress} | List folders on a peer, starting from root
[**ListPeersAddresses**](ConfigServiceApi.md#ListPeersAddresses) | **Get** /config/peers | List all detected peers (servers on which the app is running)
[**ListServices**](ConfigServiceApi.md#ListServices) | **Get** /config/ctl | List all services and their status
[**ListVersioningPolicies**](ConfigServiceApi.md#ListVersioningPolicies) | **Get** /config/versioning | List all defined versioning policies
[**OpenApiDiscovery**](ConfigServiceApi.md#OpenApiDiscovery) | **Get** /config/discovery/openapi | Publish available REST APIs
[**PutConfig**](ConfigServiceApi.md#PutConfig) | **Put** /config/{FullPath} | Generic config Put, using a full path in the config tree
[**PutDataSource**](ConfigServiceApi.md#PutDataSource) | **Post** /config/datasource/{Name} | Create or update a datasource


# **ConfigFormsDiscovery**
> RestDiscoveryResponse ConfigFormsDiscovery(ctx, serviceName)
Publish Forms definition for building screens in frontend

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serviceName** | **string**|  | 

### Return type

[**RestDiscoveryResponse**](restDiscoveryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **ControlService**
> CtlService ControlService(ctx, body)
[Not Implemented]  Start/Stop a service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**RestControlServiceRequest**](RestControlServiceRequest.md)|  | 

### Return type

[**CtlService**](ctlService.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **CreateEncryptionKey**
> EncryptionAdminCreateKeyResponse CreateEncryptionKey(ctx, body)
Create a new master key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**EncryptionAdminCreateKeyRequest**](EncryptionAdminCreateKeyRequest.md)|  | 

### Return type

[**EncryptionAdminCreateKeyResponse**](encryptionAdminCreateKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **DeleteDataSource**
> RestDeleteDataSourceResponse DeleteDataSource(ctx, name)
Delete a datasource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**|  | 

### Return type

[**RestDeleteDataSourceResponse**](restDeleteDataSourceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **DeleteEncryptionKey**
> EncryptionAdminDeleteKeyResponse DeleteEncryptionKey(ctx, body)
Delete an existing master key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**EncryptionAdminDeleteKeyRequest**](EncryptionAdminDeleteKeyRequest.md)|  | 

### Return type

[**EncryptionAdminDeleteKeyResponse**](encryptionAdminDeleteKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **EndpointsDiscovery**
> RestDiscoveryResponse EndpointsDiscovery(ctx, optional)
Publish available endpoints

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointType** | **string**|  | 

### Return type

[**RestDiscoveryResponse**](restDiscoveryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **ExportEncryptionKey**
> EncryptionAdminExportKeyResponse ExportEncryptionKey(ctx, body)
Export a master key for backup purpose, protected with a password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**EncryptionAdminExportKeyRequest**](EncryptionAdminExportKeyRequest.md)|  | 

### Return type

[**EncryptionAdminExportKeyResponse**](encryptionAdminExportKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **GetConfig**
> RestConfiguration GetConfig(ctx, fullPath, optional)
Generic config Get using a full path in the config tree

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fullPath** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fullPath** | **string**|  | 
 **data** | **string**|  | 

### Return type

[**RestConfiguration**](restConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **GetDataSource**
> ObjectDataSource GetDataSource(ctx, name, optional)
Load datasource information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
 **disabled** | **bool**|  | 
 **storageType** | **string**|  | [default to LOCAL]
 **objectsServiceName** | **string**|  | 
 **objectsHost** | **string**|  | 
 **objectsPort** | **int32**|  | 
 **objectsSecure** | **bool**|  | 
 **objectsBucket** | **string**|  | 
 **objectsBaseFolder** | **string**|  | 
 **apiKey** | **string**|  | 
 **apiSecret** | **string**|  | 
 **peerAddress** | **string**|  | 
 **watch** | **bool**|  | 
 **encryptionMode** | **string**|  | [default to CLEAR]
 **encryptionKey** | **string**|  | 
 **versioningPolicyName** | **string**|  | 
 **creationDate** | **int32**|  | 
 **lastSynchronizationDate** | **int32**|  | 

### Return type

[**ObjectDataSource**](objectDataSource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **GetVersioningPolicy**
> TreeVersioningPolicy GetVersioningPolicy(ctx, uuid, optional)
Load a given versioning policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uuid** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string**|  | 
 **name** | **string**|  | 
 **description** | **string**|  | 
 **versionsDataSourceName** | **string**|  | 
 **versionsDataSourceBucket** | **string**|  | 
 **maxTotalSize** | **string**|  | 
 **maxSizePerFile** | **string**|  | 
 **ignoreFilesGreaterThan** | **string**|  | 

### Return type

[**TreeVersioningPolicy**](treeVersioningPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **ImportEncryptionKey**
> EncryptionAdminImportKeyResponse ImportEncryptionKey(ctx, body)
Import a previously exported master key, requires the password created at export time

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**EncryptionAdminImportKeyRequest**](EncryptionAdminImportKeyRequest.md)|  | 

### Return type

[**EncryptionAdminImportKeyResponse**](encryptionAdminImportKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **ListDataSources**
> RestDataSourceCollection ListDataSources(ctx, )
List all defined datasources

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RestDataSourceCollection**](restDataSourceCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **ListEncryptionKeys**
> EncryptionAdminListKeysResponse ListEncryptionKeys(ctx, body)
List registered master keys

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**EncryptionAdminListKeysRequest**](EncryptionAdminListKeysRequest.md)|  | 

### Return type

[**EncryptionAdminListKeysResponse**](encryptionAdminListKeysResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **ListPeerFolders**
> RestNodesCollection ListPeerFolders(ctx, peerAddress, body)
List folders on a peer, starting from root

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **peerAddress** | **string**|  | 
  **body** | [**RestListPeerFoldersRequest**](RestListPeerFoldersRequest.md)|  | 

### Return type

[**RestNodesCollection**](restNodesCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **ListPeersAddresses**
> RestListPeersAddressesResponse ListPeersAddresses(ctx, )
List all detected peers (servers on which the app is running)

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RestListPeersAddressesResponse**](restListPeersAddressesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **ListServices**
> RestServiceCollection ListServices(ctx, optional)
List all services and their status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **statusFilter** | **string**|  | [default to ANY]

### Return type

[**RestServiceCollection**](restServiceCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **ListVersioningPolicies**
> RestVersioningPolicyCollection ListVersioningPolicies(ctx, )
List all defined versioning policies

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RestVersioningPolicyCollection**](restVersioningPolicyCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **OpenApiDiscovery**
> RestOpenApiResponse OpenApiDiscovery(ctx, optional)
Publish available REST APIs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointType** | **string**|  | 

### Return type

[**RestOpenApiResponse**](restOpenApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **PutConfig**
> RestConfiguration PutConfig(ctx, fullPath, body)
Generic config Put, using a full path in the config tree

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fullPath** | **string**|  | 
  **body** | [**RestConfiguration**](RestConfiguration.md)|  | 

### Return type

[**RestConfiguration**](restConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **PutDataSource**
> ObjectDataSource PutDataSource(ctx, name, body)
Create or update a datasource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**|  | 
  **body** | [**ObjectDataSource**](ObjectDataSource.md)|  | 

### Return type

[**ObjectDataSource**](objectDataSource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

