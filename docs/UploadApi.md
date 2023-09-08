# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminUploadFileDeletePost**](UploadApi.md#AdminUploadFileDeletePost) | **Post** /admin/upload/file/delete | 删除文件
[**AdminUploadFileListPost**](UploadApi.md#AdminUploadFileListPost) | **Post** /admin/upload/file/list | 文件列表
[**AdminUploadFilePost**](UploadApi.md#AdminUploadFilePost) | **Post** /admin/upload/file | 上传文件

# **AdminUploadFileDeletePost**
> InlineResponse2007 AdminUploadFileDeletePost(ctx, body)
删除文件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormFileDeleteReq**](XbuyAppFormAdminFormFileDeleteReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminUploadFileListPost**
> InlineResponse200180 AdminUploadFileListPost(ctx, body)
文件列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormFileListReq**](XbuyAppFormAdminFormFileListReq.md)|  | 

### Return type

[**InlineResponse200180**](inline_response_200_180.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminUploadFilePost**
> InlineResponse200179 AdminUploadFilePost(ctx, body)
上传文件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormUploadFileReq**](XbuyAppFormAdminFormUploadFileReq.md)|  | 

### Return type

[**InlineResponse200179**](inline_response_200_179.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

