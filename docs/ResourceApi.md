# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminResourceGroupAddPost**](ResourceApi.md#AdminResourceGroupAddPost) | **Post** /admin/resource/group_add | 添加资源分组
[**AdminResourceGroupAddResourcePost**](ResourceApi.md#AdminResourceGroupAddResourcePost) | **Post** /admin/resource/group_add_resource | 分组添加资源
[**AdminResourceGroupBatchAddResourcePost**](ResourceApi.md#AdminResourceGroupBatchAddResourcePost) | **Post** /admin/resource/group_batch_add_resource | 分组批量添加资源
[**AdminResourceGroupListPost**](ResourceApi.md#AdminResourceGroupListPost) | **Post** /admin/resource/group_list | 获取所有API分组列表
[**AdminResourceGroupResourceListPost**](ResourceApi.md#AdminResourceGroupResourceListPost) | **Post** /admin/resource/group_resource_list | 分组资源列表
[**AdminResourceListPost**](ResourceApi.md#AdminResourceListPost) | **Post** /admin/resource/list | 获取API列表

# **AdminResourceGroupAddPost**
> InlineResponse200138 AdminResourceGroupAddPost(ctx, body)
添加资源分组

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormResourceGroupAddReq**](XbuyAppFormAdminFormResourceGroupAddReq.md)|  | 

### Return type

[**InlineResponse200138**](inline_response_200_138.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminResourceGroupAddResourcePost**
> InlineResponse2007 AdminResourceGroupAddResourcePost(ctx, body)
分组添加资源

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGroupAddResourceReq**](XbuyAppFormAdminFormGroupAddResourceReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminResourceGroupBatchAddResourcePost**
> InlineResponse2007 AdminResourceGroupBatchAddResourcePost(ctx, body)
分组批量添加资源

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGroupBatchAddResourceReq**](XbuyAppFormAdminFormGroupBatchAddResourceReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminResourceGroupListPost**
> InlineResponse200139 AdminResourceGroupListPost(ctx, body)
获取所有API分组列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormResourceGroupListReq**](XbuyAppFormAdminFormResourceGroupListReq.md)|  | 

### Return type

[**InlineResponse200139**](inline_response_200_139.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminResourceGroupResourceListPost**
> InlineResponse200140 AdminResourceGroupResourceListPost(ctx, body)
分组资源列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGroupResourceListReq**](XbuyAppFormAdminFormGroupResourceListReq.md)|  | 

### Return type

[**InlineResponse200140**](inline_response_200_140.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminResourceListPost**
> InlineResponse200141 AdminResourceListPost(ctx, body)
获取API列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormResourceListReq**](XbuyAppFormAdminFormResourceListReq.md)|  | 

### Return type

[**InlineResponse200141**](inline_response_200_141.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

