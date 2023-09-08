# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminLogClearPost**](LogApi.md#AdminLogClearPost) | **Post** /admin/log/clear | 清空日志
[**AdminLogExportGet**](LogApi.md#AdminLogExportGet) | **Get** /admin/log/export | 导出日志
[**AdminLogListGet**](LogApi.md#AdminLogListGet) | **Get** /admin/log/list | 获取日志列表

# **AdminLogClearPost**
> InlineResponse2007 AdminLogClearPost(ctx, body)
清空日志

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormLogClearReq**](XbuyAppFormAdminFormLogClearReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminLogExportGet**
> InlineResponse2007 AdminLogExportGet(ctx, optional)
导出日志

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogApiAdminLogExportGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogApiAdminLogExportGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| 当前页码 | [default to 1]
 **limit** | **optional.Int32**| 每页数量 | [default to 10]
 **startTime** | **optional.String**| 开始日期 | 
 **endTime** | **optional.String**| 结束日期 | 
 **module** | **optional.String**| 应用端口 | 
 **memberId** | **optional.Int32**| 用户ID | 
 **takeUpTime** | **optional.Int32**| 请求耗时 | 
 **method** | **optional.String**| 请求方式 | 
 **url** | **optional.String**| 请求路径 | 
 **ip** | **optional.String**| 访问IP | 
 **errorCode** | **optional.String**| 状态码 | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminLogListGet**
> InlineResponse20056 AdminLogListGet(ctx, optional)
获取日志列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogApiAdminLogListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogApiAdminLogListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| 当前页码 | [default to 1]
 **limit** | **optional.Int32**| 每页数量 | [default to 10]
 **startTime** | **optional.String**| 开始日期 | 
 **endTime** | **optional.String**| 结束日期 | 
 **module** | **optional.String**| 应用端口 | 
 **memberId** | **optional.Int32**| 用户ID | 
 **takeUpTime** | **optional.Int32**| 请求耗时 | 
 **method** | **optional.String**| 请求方式 | 
 **url** | **optional.String**| 请求路径 | 
 **ip** | **optional.String**| 访问IP | 
 **errorCode** | **optional.String**| 状态码 | 

### Return type

[**InlineResponse20056**](inline_response_200_56.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

