# ProtobufAny

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeUrl** | **string** | A URL/resource name whose content describes the type of the serialized protocol buffer message.  For URLs which use the scheme &#x60;http&#x60;, &#x60;https&#x60;, or no scheme, the following restrictions and interpretations apply:  * If no scheme is provided, &#x60;https&#x60; is assumed. * The last segment of the URL&#39;s path must represent the fully   qualified name of the type (as in &#x60;path/google.protobuf.Duration&#x60;).   The name should be in a canonical form (e.g., leading \&quot;.\&quot; is   not accepted). * An HTTP GET on the URL must yield a [google.protobuf.Type][]   value in binary format, or produce an error. * Applications are allowed to cache lookup results based on the   URL, or have them precompiled into a binary to avoid any   lookup. Therefore, binary compatibility needs to be preserved   on changes to types. (Use versioned type names to manage   breaking changes.)  Schemes other than &#x60;http&#x60;, &#x60;https&#x60; (or the empty scheme) might be used with implementation specific semantics. | [optional] [default to null]
**Value** | **string** | Must be a valid serialized protocol buffer of the above specified type. | [optional] [default to null]

[[Back to Model list]](../../README.md#documentation-for-models) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to README]](../../README.md)


