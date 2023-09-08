# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminNoticeDeletePost**](NoticeApi.md#AdminNoticeDeletePost) | **Post** /admin/notice/delete | 删除公告
[**AdminNoticeEditPost**](NoticeApi.md#AdminNoticeEditPost) | **Post** /admin/notice/edit | 修改/新增公告
[**AdminNoticeListGet**](NoticeApi.md#AdminNoticeListGet) | **Get** /admin/notice/list | 获取公告列表
[**AdminNoticeMaxSortGet**](NoticeApi.md#AdminNoticeMaxSortGet) | **Get** /admin/notice/max_sort | 公告最大排序
[**AdminNoticeNameUniqueGet**](NoticeApi.md#AdminNoticeNameUniqueGet) | **Get** /admin/notice/name_unique | 公告名称是否唯一
[**AdminNoticeViewGet**](NoticeApi.md#AdminNoticeViewGet) | **Get** /admin/notice/view | 获取指定信息

# **AdminNoticeDeletePost**
> InlineResponse2007 AdminNoticeDeletePost(ctx, body)
删除公告

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormNoticeDeleteReq**](XbuyAppFormAdminFormNoticeDeleteReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminNoticeEditPost**
> InlineResponse2007 AdminNoticeEditPost(ctx, body)
修改/新增公告

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormNoticeEditReq**](XbuyAppFormAdminFormNoticeEditReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminNoticeListGet**
> InlineResponse200118 AdminNoticeListGet(ctx, optional)
获取公告列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NoticeApiAdminNoticeListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NoticeApiAdminNoticeListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| 当前页码 | [default to 1]
 **limit** | **optional.Int32**| 每页数量 | [default to 10]
 **startTime** | **optional.String**| 开始日期 | 
 **endTime** | **optional.String**| 结束日期 | 
 **status** | **optional.Int32**| 状态 | 
 **title** | **optional.String**| 公告名称 | 
 **type_** | **optional.String**| 公告类型 | 

### Return type

[**InlineResponse200118**](inline_response_200_118.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminNoticeMaxSortGet**
> InlineResponse20021 AdminNoticeMaxSortGet(ctx, optional)
公告最大排序

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NoticeApiAdminNoticeMaxSortGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NoticeApiAdminNoticeMaxSortGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int64**| 公告ID | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminNoticeNameUniqueGet**
> InlineResponse20022 AdminNoticeNameUniqueGet(ctx, name, optional)
公告名称是否唯一

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| 公告名称 | 
 **optional** | ***NoticeApiAdminNoticeNameUniqueGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NoticeApiAdminNoticeNameUniqueGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.Int64**| 公告ID | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminNoticeViewGet**
> InlineResponse200119 AdminNoticeViewGet(ctx, id)
获取指定信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| 公告ID | 

### Return type

[**InlineResponse200119**](inline_response_200_119.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

