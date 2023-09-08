# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminServiceAddPost**](ServiceApi.md#AdminServiceAddPost) | **Post** /admin/service/add | 增值服务添加/编辑
[**AdminServiceDetailPost**](ServiceApi.md#AdminServiceDetailPost) | **Post** /admin/service/detail | 增值服务详情
[**AdminServiceListPost**](ServiceApi.md#AdminServiceListPost) | **Post** /admin/service/list | 增值服务列表
[**AdminServicePausePost**](ServiceApi.md#AdminServicePausePost) | **Post** /admin/service/pause | 增值服务暂停/开启

# **AdminServiceAddPost**
> InlineResponse200146 AdminServiceAddPost(ctx, body)
增值服务添加/编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormServiceAddReq**](XbuyAppFormAdminFormServiceAddReq.md)|  | 

### Return type

[**InlineResponse200146**](inline_response_200_146.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminServiceDetailPost**
> InlineResponse200147 AdminServiceDetailPost(ctx, body)
增值服务详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormServiceDetailReq**](XbuyAppFormAdminFormServiceDetailReq.md)|  | 

### Return type

[**InlineResponse200147**](inline_response_200_147.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminServiceListPost**
> InlineResponse200148 AdminServiceListPost(ctx, body)
增值服务列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormServiceListReq**](XbuyAppFormAdminFormServiceListReq.md)|  | 

### Return type

[**InlineResponse200148**](inline_response_200_148.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminServicePausePost**
> InlineResponse200146 AdminServicePausePost(ctx, body)
增值服务暂停/开启

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormServicePauseReq**](XbuyAppFormAdminFormServicePauseReq.md)|  | 

### Return type

[**InlineResponse200146**](inline_response_200_146.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

